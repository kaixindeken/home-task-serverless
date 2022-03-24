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

func CreateProducer(client pulsar.Client, topic string, inputPath string) {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(inputPath) // just pass the file name
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

	for i := 0; i < 10; i++ {
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
	}

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}

type argT struct {
	cli.Helper
	Topic  string `cli:"t, topic"  usage:"topic unspecified"`
	Input  string `cli:"i, input"  usage:"input path unknown"`
	Output string `cli:"o, output" usage:"output path unknown"`
	Script string `cli:"s, script" usage:"script function unspecified"`
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
		go CreateConsumer(client, argv.Topic, argv.Script, argv.Output)
		time.Sleep(1e9)
		go CreateProducer(client, argv.Topic, argv.Input)
		time.Sleep(3e9)
		return nil
	}))

}
