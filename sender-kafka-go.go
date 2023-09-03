package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
	// Créer un logger
	logger := log.New(os.Stderr, "[Kafka] ", log.LstdFlags)

	// Configuration SASL
	mechanism := plain.Mechanism{
		Username: "2VA6VoYGxOwO7xUdDyAxkX",
		Password: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29uZHVrdG9yLmlvIiwic291cmNlQXBwbGljYXRpb24iOiJhZG1pbiIsInVzZXJNYWlsIjpudWxsLCJwYXlsb2FkIjp7InZhbGlkRm9yVXNlcm5hbWUiOiIyVkE2Vm9ZR3hPd083eFVkRHlBeGtYIiwib3JnYW5pemF0aW9uSWQiOjc2MTQ4LCJ1c2VySWQiOjg4NjA4LCJmb3JFeHBpcmF0aW9uQ2hlY2siOiJjNjkwMDZlNC1mMDBiLTQ4MmQtOGFjNy1jODBiYTM3NTY2YjIifX0.4YH2g0vHMezL1s74vTOlxdzDjgLOUm-7jOqSsU8CaJI",
	}

	// Configuration du writer en utilisant le dialer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{"cluster.playground.cdkt.io:9092"}, // Assurez-vous de remplacer par l'adresse de votre broker
		Topic:            "test_go",                      // Assurez-vous de remplacer par le nom de votre topic
		Logger:           logger,
		Dialer:           &kafka.Dialer{
			TLS:           &tls.Config{InsecureSkipVerify: true},
			SASLMechanism: mechanism,
		},
	})

	// Envoi d'un message

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// key := "my key"
	message := "Hello, Kafka using go! " + currentTime

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			// Key: []byte(key),
			Value: []byte(message),
		},
	)
	if err != nil {
		logger.Fatalf("failed to write message: %v", err)
	}

	w.Close()
}
