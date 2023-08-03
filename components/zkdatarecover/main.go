package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-zookeeper/zk"
)

func shouldUpload(conn *zk.Conn, localIPPort string) (should bool) {
	if exist, _, _ := conn.Exists("/root/app/sub_app/master"); !exist {
		conn.Create("/root", nil, 0, zk.WorldACL(zk.PermAll))
		conn.Create("/root/app", nil, 0, zk.WorldACL(zk.PermAll))
		conn.Create("/root/app/sub_app", []byte(""), 0, zk.WorldACL(zk.PermAll))
		conn.Create("/root/app/sub_app/master", []byte(localIPPort), 0, zk.WorldACL(zk.PermAll))
	}

	ipPort, _, err := conn.Get("/root/app/sub_app/master")
	if err != nil {
		return
	}

	// 非本机时检查是否服务器挂了
	if string(ipPort) != localIPPort {
		if exist, _, _ := conn.Exists("/root/server/" + string(ipPort)); exist {
			return
		}

		// 重置 IP
		if _, err = conn.Set("/root/app/sub_app/master", []byte(localIPPort), -1); err != nil {
			return
		}
	}

	return true
}

func main() {

	printCurrent := false

	zkhosts := []string{"192.168.1.1"}

	var conn *zk.Conn
	var err error
	byts, _ := ioutil.ReadFile("zookeeper.session")
	if len(byts) != 0 {
		var id int64
		var pass string
		fmt.Sscanf(string(byts), "%d,%s", &id, &pass)
		passwd, _ := hex.DecodeString(pass)
		conn, _, err = zk.Connect(zkhosts, 70*time.Second, zk.WithSessionID(id), zk.WithPassword(passwd))
	} else {
		conn, _, err = zk.Connect(zkhosts, 70*time.Second)
	}

	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		time.Sleep(time.Second * 2)
		ioutil.WriteFile("zookeeper.session", []byte(fmt.Sprintf("%d,%s", conn.SessionID(), hex.EncodeToString(conn.Password()))), os.ModePerm)
		wg.Done()
	}()

	agentIPsbyt, err := ioutil.ReadFile("zkips")
	reg, _ := regexp.Compile(`(.+?) (.+)\s*`)
	match := reg.FindAllStringSubmatch(string(agentIPsbyt), -1)

	serverAgentIP := make(map[string]string)
	for _, elements := range match {
		serverAgentIP[elements[2]] = elements[1] + ":8002"
	}

	fmt.Println()
	fmt.Println("zk: ")

	ch, _, err := conn.Children("/root/server")
	fmt.Println("servers", ch)

	zkServerAgentIP := make(map[string]string)
	for _, server := range ch {
		s := fmt.Sprintf("/root/server/%s", server)

		nodeData, _, err := conn.Get(s)
		if err != nil {
			continue
		}

		if printCurrent {
			fmt.Println()
			fmt.Println()
			fmt.Println(server, string(nodeData))
		}

		agents := strings.Split(string(nodeData), ";")

		for _, ip := range agents {
			zkServerAgentIP[ip] = server
		}
	}

	fmt.Println()
	for agent, server := range zkServerAgentIP {
		if _, ok := serverAgentIP[agent]; !ok {
			// fmt.Println(agent, "not found in log")
			continue
		}

		if server != serverAgentIP[agent] {
			fmt.Printf("agent %s should in %s (log), but is in %s from zk. \n", agent, serverAgentIP[agent], server)
		}
	}

	toAdd := make(map[string][]string)
	for _, server := range ch {
		toAdd[server] = make([]string, 0)
	}

	for agent, server := range serverAgentIP {
		// 过滤掉已存在于 zk 的 IP
		if _, ok := zkServerAgentIP[agent]; ok {
			continue
		}

		toAdd[server] = append(toAdd[server], agent)
	}

	// fmt.Println()
	// fmt.Println("to delete: ")

	// toDelete := map[string]bool{
	// }

	// for _, server := range ch {
	// 	s := fmt.Sprintf("/root/server/%s", server)

	// 	nodeData, _, err := conn.Get(s)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	agents := strings.Split(string(nodeData), ";")
	// 	newAgents := make([]string, 0)
	// 	for _, ip := range agents {
	// 		if _,ok := toDelete[ip]; !ok {
	// 			newAgents = append(newAgents, ip)
	// 		}
	// 	}


	// 	update := strings.Join(newAgents, ";") + ";"
	// 	_, err = conn.Set(s, []byte(update), -1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	// return

	fmt.Println()
	fmt.Println("to add: ")

	for server, agents := range toAdd {
		if len(agents) == 0 {
			continue
		}
		fmt.Println(server, strings.Join(agents, ";")+";")
	}

	ipReg, _ := regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
	for server, agents := range toAdd {
		if len(agents) == 0 {
			continue
		}
		s := fmt.Sprintf("/root/server/%s", server)
		fmt.Println("saving:", s)
		nodeData, _, err := conn.Get(s)
		if err != nil {
			path, err := conn.Create(s, nil, 0, zk.WorldACL(zk.PermAll))
			if err != nil {
				panic(err)
			}
			fmt.Println(path)
		}

		dataToAdd := strings.Join(agents, ";") + ";"

		newData := string(nodeData) + dataToAdd

		splitedIPs := strings.Split(string(newData), ";")
		for _, ip := range splitedIPs {
			if ip != "" && !ipReg.MatchString(ip) {
				panic("not match")
			}
		}

		_, err = conn.Set(s, []byte(newData), -1)
		if err != nil {
			fmt.Println(err)
		}
	}

	wg.Wait()
}
