package main

import (
	"fmt"
	"log"
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

	producer, err := sarama.NewSyncProducer([]string{"cluster.playground.cdkt.io:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to initialize producer: %v", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "test_sarama_go",
		Value: sarama.StringEncoder("Hello, Kafka using Sarama!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
