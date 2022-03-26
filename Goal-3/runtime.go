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
	input  = "/Users/zhangruian/dev/stdin"
	output = "/Users/zhangruian/dev/stdout"
)

func CreateConsumer(client pulsar.Client, topic string, script string, outputPath string) {
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
	cmd := exec.Command("./functions/"+script+".sh", string(msg.Payload()))
	message, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to execute bash script: ", err)
	}
	err = ioutil.WriteFile(outputPath, message, 0777)
	if err != nil {
		fmt.Println("Failed to write file: ", err)
	}
	fmt.Printf("%s\n", message)
	consumer.Ack(msg)

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}

func CreateProducer(client pulsar.Client, topic string, inputPath string) {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Failed to publish InputPath: ", err)
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: b,
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")
}

type argT struct {
	cli.Helper
	Consume bool   `cli:"c, consume"  usage:"sign of create consumer"`
	Produce bool   `cli:"p, produce"  usage:"sign of create producer"`
	Topic   string `cli:"t, topic"  usage:"topic unspecified"`
	Script  string `cli:"s, script" usage:"script function unspecified"`
}

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}
	defer client.Close()

	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		if argv.Consume && !argv.Produce {
			CreateConsumer(client, argv.Topic, argv.Script, output)
		} else if !argv.Consume && argv.Produce {
			CreateProducer(client, argv.Topic, input)
		}
		return nil
	}))
}
