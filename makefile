all: build

build: 
	go build -o kafka_check nagios_kafka_topic_check.go