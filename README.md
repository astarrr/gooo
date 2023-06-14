# Go and Kafka Personal Side Project

This project is a personal endeavor aimed at learning Go programming language and exploring the Kafka messaging system. It involves the development of a producer that reads Twitch messages and sends them to Kafka, and a Go consumer that retrieves and processes these messages from Kafka.

## Getting Started

To set up and run the project, follow these instructions:

1. Start Zookeeper:
   ```bash
   zkServer start
   ```

2. Start Kafka Server. I use:
   ```bash
   ./kafka-server-start /opt/homebrew/etc/kafka/server.properties
   ```

3. Start the Producer:
   ```bash
   go run producer/main.go
   ```

4. Start the Consumer:
   ```bash
   go run consumer/main.go
   ```

## Contributing

This project is open to contributions, suggestions, and improvements. If you have any ideas or feedback, please feel free to open an issue or submit a pull request on the project's repository.