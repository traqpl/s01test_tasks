package main

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func main() {

	//todo: prametrize from command line s
	// brokers := []string{"localhost:9092", "127.0.0.1:9092"}
	// topic := "test"
	// login := "somelogin"
	// password := "somepassword"
	topic := os.Args[1]
	login := os.Args[2]
	password := os.Args[3]
	brokers := os.Args[4:]

	var status string
	clusterStatus := true
	var problemServers []string

	for _, s := range brokers {
		server := []string{s}
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		config.Net.SASL.User = login
		config.Net.SASL.Password = password

		cluster, err := sarama.NewConsumer(server, config)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := cluster.Close(); err != nil {
				panic(err)
			}
		}()

		topics, _ := cluster.Topics()
		for index := range topics {
			if topics[index] == topic {
				status = "OK"
			}
		}
		if status != "OK" {
			clusterStatus = false
			problemServers = append(problemServers, s)
		}
	}
	if clusterStatus == false {
		status = "CRITICAL"
		fmt.Printf("%s - topic %s was NOT FOUND on servers %s\n", status, topic, problemServers)
		os.Exit(2)
	} else {
		fmt.Printf("%s - topic %s WAS FOUND on all cluster servers\n", status, topic)
		os.Exit(0)
	}
}
