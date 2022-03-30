package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/mkideal/cli"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

type Config struct {
	Broker        string `json:"broker"`
	InputFilePath string `json:"input_file_path"`
	FunctionRoot  string `json:"function_root"`
	InputTopic    string `json:"input_topic"`
	OutputTopic   string `json:"output_topic"`
	Script        string `json:"script"`
}

func ReadFile(filePath string) []byte {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed open file: ", err)
	}
	return b
}

func FileReader(filePath string, ch chan []byte) {
	b := ReadFile(filePath)
	ch <- b
}

func ExecCommand(config Config, payload string) *exec.Cmd {
	return exec.Command("sh", config.FunctionRoot+config.Script+".sh", payload)
}

func CreateConsumer(client pulsar.Client, config Config, ch chan []byte) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            config.InputTopic,
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
	cmd := ExecCommand(config, string(msg.Payload()))
	message, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to execute bash script: ", err)
	}
	fmt.Println("Message: ", message)
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
	Config string `cli:"c, config"  usage:"please appoint config path"`
}

func main() {
	config := Config{}
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		b := ReadFile(argv.Config)
		err := json.Unmarshal(b, &config)
		if err != nil {
			fmt.Println("Failed decode config: ", err)
		}
		client, err := pulsar.NewClient(pulsar.ClientOptions{
			URL:               config.Broker,
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
		if err != nil {
			log.Fatalf("Could not instantiate Pulsar client: %v", err)
		}
		defer client.Close()
		fileProCh := make(chan []byte)
		conProCh := make(chan []byte)
		go FileReader(config.InputFilePath, fileProCh)
		time.Sleep(1e9)
		go CreateConsumer(client, config, conProCh)
		time.Sleep(3e9)
		go CreateProducer(client, config.InputTopic, fileProCh)
		time.Sleep(3e9)
		go CreateProducer(client, config.OutputTopic, conProCh)
		time.Sleep(3e9)
		return nil
	}))
}
