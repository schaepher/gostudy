package gostudy

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCert(t *testing.T) {
	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// CA 证书
	caCert, err := ioutil.ReadFile("ca-chain.cert.pem")
	check(err)
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)

	// 客户端证书
	clientCrt, err := tls.LoadX509KeyPair("vendor.cert.pem", "vender.key.pem")
	check(err)

	// 客户端
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{clientCrt},
	}
	client := &http.Client{Transport: tr}

	// 请求
	api := "https://example.com/connectors/v1/aggregator/vendor"
	body := []byte{}
	req, err := http.NewRequest("POST", api, bytes.NewReader(body))
	check(err)
	req.Header.Add("Content-Encoding", "gzip")

	// 发送
	resp, err := client.Do(req)
	check(err)

	if resp.StatusCode == 201 {
		fmt.Println("Created")
	}
}
