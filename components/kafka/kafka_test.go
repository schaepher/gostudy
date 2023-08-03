package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/cast"
)

// NewStatisticConsumer 新建统计数据消费者
func NewStatisticConsumer(addrs []string, consumerGroupName string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return sarama.NewConsumerGroup(addrs, consumerGroupName, config)
}

func TestKafkaOffsetByTime(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Session.Timeout = time.Minute * 10

	client, err := sarama.NewClient(kafkaAddr, config)
	if err != nil {
		return
	}
	defer client.Close()

	cg, err := sarama.NewConsumerGroupFromClient("group-L2", client)
	if err != nil {
		println(err)
		return
	}
	ts := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	ctx := context.Background()
	for {
		println("start")
		if err = cg.Consume(ctx, []string{"the-topic"}, statisticsHandler{
			topic:  "the-topic",
			client: client,
			start:  ts,
		}); err != nil {
			println(err.Error())
			return
		}
		cg.Close()
		println("end")
		println()
	}
}

// statisticsHandler 处理器
type statisticsHandler struct {
	topic  string
	client sarama.Client
	start  time.Time
}

func (h statisticsHandler) Setup(sess sarama.ConsumerGroupSession) error {
	ids, err := h.client.Partitions(h.topic)
	if err != nil {
		return nil
	}

	ts := h.start.Unix() * 1000
	for _, id := range ids {
		offset, err := h.client.GetOffset(h.topic, id, ts)
		if err != nil {
			return nil
		}
		println(offset)
		sess.MarkOffset(h.topic, id, sarama.OffsetOldest, "")
		sess.ResetOffset(h.topic, id, sarama.OffsetOldest, "")
	}

	return nil
}

func (statisticsHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (h statisticsHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg := <-claim.Messages():
			if msg == nil {
				continue
			}
			fmt.Println(msg.Offset, msg.Timestamp.Unix(), string(msg.Value))
			if msg.Offset%100 == 0 {
				println("mark ", msg.Offset)
				sess.MarkMessage(msg, "")
				sess.Commit()
			}
		case <-sess.Context().Done():
			return nil
		}
	}
}

var kafkaAddr = []string{"192.168.1.1:9092"}

func TestKafkaConsume(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	c, err := sarama.NewConsumer(kafkaAddr, config)
	if err != nil {
		println(err)
		return
	}

	p, err := c.Partitions("the-topic")
	if err != nil {
		println(err)
		return
	}
	for _, partition := range p {
		con , er := c.ConsumePartition("the-topic", partition, 1)
		if er != nil {
			println(er)
			return
		}
		for element := range con.Messages() {
			if element != nil {
				println(string(element.Value))
			}
		}
	}
}

func TestKafkaConsumeG(t *testing.T) {
	cg, err := NewStatisticConsumer(kafkaAddr, "the-topic")
	if err != nil {
		return
	}

	ctx := context.Background()
	for {
		println("start")
		if err = cg.Consume(ctx, []string{"the-topic"}, statisticsHandler{}); err != nil {
			return
		}
		cg.Close()
		println("end")
		println()
	}
}

func TestKafkaInit(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0
	admin, err := sarama.NewClusterAdmin(kafkaAddr, config)
	if err != nil {
		panic(err)
	}
	err = admin.CreateTopic("test", &sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 1}, true)
	if err != nil {
		panic(err)
	}
}

func TestKafkaAdd(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0
	// 牺牲性能避免数据丢失
	config.Producer.RequiredAcks = sarama.WaitForLocal
	// 牺牲性能避免数据丢失
	config.Producer.Return.Successes = false
	// Random 的速度比较快
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewAsyncProducer(kafkaAddr, config)
	if err != nil {
		return
	}

	for i := 0; i < 1000; i++ {
		producer.Input() <- &sarama.ProducerMessage{
			Topic: "test",
			Key:   sarama.StringEncoder("a"),
			Value: sarama.StringEncoder(cast.ToString(i)),
		}
	}

	producer.Close()
}
