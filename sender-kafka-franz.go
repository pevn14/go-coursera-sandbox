package main

import (
	"context"
	"fmt"
	// "os"
	"time"
	"crypto/tls"
	"net"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
)

func main() {
	topic := "test_franz_go"
	// key := "my key"
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	message := "Hello, Kafka using Franz-go! " + currentTime
	seeds := []string{"cluster.playground.cdkt.io:9092"}

	user := "2VA6VoYGxOwO7xUdDyAxkX"
	password := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29uZHVrdG9yLmlvIiwic291cmNlQXBwbGljYXRpb24iOiJhZG1pbiIsInVzZXJNYWlsIjpudWxsLCJwYXlsb2FkIjp7InZhbGlkRm9yVXNlcm5hbWUiOiIyVkE2Vm9ZR3hPd083eFVkRHlBeGtYIiwib3JnYW5pemF0aW9uSWQiOjc2MTQ4LCJ1c2VySWQiOjg4NjA4LCJmb3JFeHBpcmF0aW9uQ2hlY2siOiJjNjkwMDZlNC1mMDBiLTQ4MmQtOGFjNy1jODBiYTM3NTY2YjIifX0.4YH2g0vHMezL1s74vTOlxdzDjgLOUm-7jOqSsU8CaJI" // Remplacez ceci par votre mot de passe réel

	tlsDialer := func(ctx context.Context, network, address string) (net.Conn, error) {
		return tls.Dial(network, address, &tls.Config{})
	}
	
	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		kgo.Dialer(tlsDialer), // Use the TLS dialer function
		kgo.SASL(plain.Auth{
			User: user,
			Pass: password,
		}.AsMechanism()),
		//kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil)), // ajout de l'option de journalisation pour le debug
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		fmt.Printf("Failed to initialize Kafka client: %v\n", err)
		return
	}
	defer client.Close()

	record := &kgo.Record{
		Topic: topic,
		// Key:   []byte(key),
		Value: []byte(message),
	}

	ctx := context.Background()

	// Utilisez ProduceSync pour une production synchrone et traitez les erreurs
	errs := client.ProduceSync(ctx, record)
	if err := errs.FirstErr(); err != nil {
		fmt.Printf("record had a produce error while synchronously producing: %v\n", err)
	} else {
		fmt.Println("Message sent successfully!")
	}
}
