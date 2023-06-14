This is a personal project where I try to learn Go and Kafka.

Simply put, there will be a producer that reads Twitch messages and ships them towards Kafka.
Then, a Go consumer, should read messages from said Kafka. 

Instructions:
- Start Zookeeper
  - zkServer start
- Start Kafka Server. I use this command
  - ./kafka-server-start /opt/homebrew/etc/kafka/server.properties
- Start producer
  - go run producer/main.go
- Start consumer
  - go run consumer/main.go