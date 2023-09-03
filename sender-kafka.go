package main

import (
	"context"
	// "log"
	"sync"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
)

func main() {

	topic := "test_franz_go"
	message := "Hello, Kafka using Franz-go!"
	seeds := []string{"cluster.playground.cdkt.io:9092"}

	// SASL Plain credentials
	user := "2VA6VoYGxOwO7xUdDyAxkX"
	password := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F1dGguY29uZHVrdG9yLmlvIiwic291cmNlQXBwbGljYXRpb24iOiJhZG1pbiIsInVzZXJNYWlsIjpudWxsLCJwYXlsb2FkIjp7InZhbGlkRm9yVXNlcm5hbWUiOiIyVkE2Vm9ZR3hPd083eFVkRHlBeGtYIiwib3JnYW5pemF0aW9uSWQiOjc2MTQ4LCJ1c2VySWQiOjg4NjA4LCJmb3JFeHBpcmF0aW9uQ2hlY2siOiJjNjkwMDZlNC1mMDBiLTQ4MmQtOGFjNy1jODBiYTM3NTY2YjIifX0.4YH2g0vHMezL1s74vTOlxdzDjgLOUm-7jOqSsU8CaJI"
	
	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		// SASL Options
		kgo.SASL(plain.Auth{
			User: user,
			Pass: password,
		}.AsMechanism()),
	}

	
	cl, err := kgo.NewClient(opts...)
	if err != nil {
		fmt.Println("Failed to initialize Kafka client: %v", err)
	} else {
		fmt.Println("Init ok")
	}
	
	var wg sync.WaitGroup
	wg.Add(1)
	
	record := &kgo.Record{
		Topic: topic,
		Value: []byte(message),
	}

	ctx := context.Background()
	cl.Produce(ctx, record, func(r *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		} else {
			fmt.Println("Send ok")
		}
	})

	wg.Wait()
// Alternatively, ProduceSync exists to synchronously produce a batch of records.
if err := cl.ProduceSync(ctx, record).FirstErr(); err != nil {
	fmt.Printf("record had a produce error while synchronously producing: %v\n", err)
}

	// 2.) Consuming messages from a topic
	for {
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			panic(fmt.Sprint(errs))
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}

		// or a callback function.
		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			for _, record := range p.Records {
				fmt.Println(string(record.Value), "from range inside a callback!")
			}

			// We can even use a second callback!
			p.EachRecord(func(record *kgo.Record) {
				fmt.Println(string(record.Value), "from a second callback!")
			})
		})
	}
}
