package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/mkideal/cli"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	broker       = "pulsar://172.17.0.2:6650"
	filePath     = "/dev2/stdin"
	functionRoot = "/root/functions/"
)

func FileReader(filePath string, ch chan []byte) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed open file: ", err)
	}
	ch <- b
}

func ExecCommand(script string, payload string) *exec.Cmd {
	return exec.Command("sh", functionRoot+script+".sh", payload)
}

func CreateConsumer(client pulsar.Client, topic string, script string, ch chan []byte) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "test",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	cmd := ExecCommand(script, string(msg.Payload()))
	message, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to execute bash script: ", err)
	}
	ch <- message

	consumer.Ack(msg)

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}

func CreateProducer(client pulsar.Client, topic string, ch chan []byte) {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		log.Fatal(err)
	}

	payload := <-ch
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: payload,
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message: ", string(payload))
}

type argT struct {
	cli.Helper
	InputTopic  string `cli:"i, input-topic"  usage:"please define input topic"`
	OutputTopic string `cli:"o, output-topic" usage:"please define output topic"`
	Script      string `cli:"s, script"       usage:"please define script function"`
}

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               broker,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}
	defer client.Close()

	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		fileProCh := make(chan []byte)
		conProCh := make(chan []byte)
		go FileReader(filePath, fileProCh)
		time.Sleep(3e9)
		go CreateConsumer(client, argv.InputTopic, argv.Script, conProCh)
		time.Sleep(3e9)
		go CreateProducer(client, argv.InputTopic, fileProCh)
		time.Sleep(3e9)
		go CreateProducer(client, argv.InputTopic, conProCh)
		time.Sleep(3e9)
		return nil
	}))

}
