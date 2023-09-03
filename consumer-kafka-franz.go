package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
)

func main() {
	topic := "test_franz_go"
	consumeGroup := "my-first-application"
	seeds := []string{"cluster.playground.cdkt.io:9092"}
	user := "2VA6VoYGxOwO7xUdDyAxkX"
	password := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29uZHVrdG9yLmlvIiwic291cmNlQXBwbGljYXRpb24iOiJhZG1pbiIsInVzZXJNYWlsIjpudWxsLCJwYXlsb2FkIjp7InZhbGlkRm9yVXNlcm5hbWUiOiIyVkE2Vm9ZR3hPd083eFVkRHlBeGtYIiwib3JnYW5pemF0aW9uSWQiOjc2MTQ4LCJ1c2VySWQiOjg4NjA4LCJmb3JFeHBpcmF0aW9uQ2hlY2siOiJjNjkwMDZlNC1mMDBiLTQ4MmQtOGFjNy1jODBiYTM3NTY2YjIifX0.4YH2g0vHMezL1s74vTOlxdzDjgLOUm-7jOqSsU8CaJI" // Remplacez ceci par votre mot de passe réel

	tlsDialer := func(ctx context.Context, network, address string) (net.Conn, error) {
		return tls.Dial(network, address, &tls.Config{})
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		kgo.Dialer(tlsDialer),
		kgo.SASL(plain.Auth{
			User: user,
			Pass: password,
		}.AsMechanism()),
		//kgo.ConsumerOffsets(kgo.NewOffset().AtStart()), // Consume from the beginning
		kgo.ConsumerGroup(consumeGroup), // A consumer group identifier
		kgo.ConsumeTopics(topic),
		kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelDebug, nil)), // ajout de l'option de journalisation pour le debug

	}

	cl, err := kgo.NewClient(opts...)
	if err != nil {
		fmt.Printf("Erreur lors de l'initialisation du client Kafka: %v\n", err)
		os.Exit(1)
	}
	defer cl.Close()

	// S'abonner au topic
	//cl.AssignGroup([]string{topic})    // ne marche pas pour le moment

	for {
		fetches := cl.PollFetches(context.Background())
		if errs := fetches.Errors(); len(errs) > 0 {
			fmt.Printf("Des erreurs ont été rencontrées: %v\n", errs)
			continue
		}

		// Filtrer par topic
		fetches.EachPartition(func(partition kgo.FetchTopicPartition) {
			if partition.Topic == topic {
				for _, record := range partition.Records {
					fmt.Printf("Message reçu avec clé: %s, valeur: %s\n", string(record.Key), string(record.Value))
				}
			}
		})
	}
}
