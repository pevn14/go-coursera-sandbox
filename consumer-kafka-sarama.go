package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()

	// Configuration de la version Kafka
	config.Version = sarama.V2_6_0_0

	// Configuration pour SASL/PLAIN
	config.Net.SASL.Enable = true
	config.Net.SASL.User = "2VA6VoYGxOwO7xUdDyAxkX"
	config.Net.SASL.Password = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29uZHVrdG9yLmlvIiwic291cmNlQXBwbGljYXRpb24iOiJhZG1pbiIsInVzZXJNYWlsIjpudWxsLCJwYXlsb2FkIjp7InZhbGlkRm9yVXNlcm5hbWUiOiIyVkE2Vm9ZR3hPd083eFVkRHlBeGtYIiwib3JnYW5pemF0aW9uSWQiOjc2MTQ4LCJ1c2VySWQiOjg4NjA4LCJmb3JFeHBpcmF0aW9uQ2hlY2siOiJjNjkwMDZlNC1mMDBiLTQ4MmQtOGFjNy1jODBiYTM3NTY2YjIifX0.4YH2g0vHMezL1s74vTOlxdzDjgLOUm-7jOqSsU8CaJI"
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext

	// Configuration pour utiliser SSL (dans votre cas SASL_SSL)
	config.Net.TLS.Enable = true

	// Nécessaire pour le SyncProducer
	config.Producer.Return.Successes = true


brokers := []string{"cluster.playground.cdkt.io:9092"}

// Create new consumer
master, err := sarama.NewConsumer(brokers, config)
if err != nil {
	panic(err)
}


defer func() {
	if err := master.Close(); err != nil {
		panic(err)
	}
}()

topics, _ := master.Topics()

consumer, errors := consume(topics, master)

signals := make(chan os.Signal, 1)
signal.Notify(signals, os.Interrupt)

// Count how many message processed
msgCount := 0

// Get signnal for finish
doneCh := make(chan struct{})
go func() {
	for {
		select {
		case msg := <-consumer:
			msgCount++
			fmt.Println("Received messages", string(msg.Key), string(msg.Value))
		case consumerError := <-errors:
			msgCount++
			fmt.Println("Received consumerError ", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
			doneCh <- struct{}{}
		case <-signals:
			fmt.Println("Interrupt is detected")
			doneCh <- struct{}{}
		}
	}
}()

<-doneCh
fmt.Println("Processed", msgCount, "messages")

}

func consume(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
consumers := make(chan *sarama.ConsumerMessage)
errors := make(chan *sarama.ConsumerError)
for _, topic := range topics {
	if strings.Contains(topic, "__consumer_offsets") {
		continue
	}
	partitions, _ := master.Partitions(topic)
// this only consumes partition no 1, you would probably want to consume all partitions
	consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
	if nil != err {
		fmt.Printf("Topic %v Partitions: %v", topic, partitions)
		panic(err)
	}
	fmt.Println(" Start consuming topic ", topic)
	go func(topic string, consumer sarama.PartitionConsumer) {
		for {
			select {
			case consumerError := <-consumer.Errors():
				errors <- consumerError
				fmt.Println("consumerError: ", consumerError.Err)

			case msg := <-consumer.Messages():
				consumers <- msg
				fmt.Println("Got message on topic ", topic, msg.Value)
			}
		}
	}(topic, consumer)
}

return consumers, errors
}
