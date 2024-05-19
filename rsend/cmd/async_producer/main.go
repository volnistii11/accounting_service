package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/volnistii11/accounting_service/rsend/internal/domain/money_requests"
	"github.com/volnistii11/accounting_service/rsend/internal/infra/kafka/async_producer"
)

func main() {
	var (
		wg    = &sync.WaitGroup{}
		conf  = newConfig(cliFlags)
		isRun = true
	)

	ctx, cancel := context.WithCancel(context.Background())
	ctx = runSignalHandler(ctx, wg)

	prod, err := async_producer.NewAsyncProducer(conf.kafka,
		async_producer.WithRequiredAcks(sarama.WaitForAll),
		async_producer.WithMaxRetries(5),
		async_producer.WithRetryBackoff(10*time.Millisecond),
		async_producer.WithProducerFlushMessages(3),
		async_producer.WithProducerFlushFrequency(5*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	runTicker(ctx, wg)
	runKafkaSuccess(ctx, prod, wg)
	runKafkaErrors(ctx, prod, wg)

	var event money_requests.Event
	for i := 0; i < conf.app.repeatCnt && isRun; i++ {
		factory := money_requests.NewDefaultFactory(conf.app.startID)
		for i := 0; i < conf.app.count && isRun; i++ {
			event = factory.Create(money_requests.EventMoneyRequestSent)

			bytes, err := json.Marshal(event)
			if err != nil {
				log.Fatal(err)
			}

			msg := &sarama.ProducerMessage{
				Topic: conf.producer.topic,
				Key:   sarama.StringEncoder(strconv.FormatInt(event.ID, 10)),
				Value: sarama.ByteEncoder(bytes),
				Headers: []sarama.RecordHeader{
					{
						Key:   []byte("accounting_service"),
						Value: []byte("rsend"),
					},
				},
				Timestamp: time.Now(),
			}

			select {
			case <-ctx.Done():
				log.Printf("[produce] ctx closed: %d\n", event.ID)
				isRun = false
				break
			case prod.Input() <- msg:
				log.Printf("[produce] msg sent: %d\n", event.ID)
			}

			time.Sleep(200 * time.Millisecond)
		}
	}

	err = prod.Close()
	if err != nil {
		log.Fatal(err)
	}

	cancel()
	wg.Wait()
}

func runTicker(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer fmt.Println("[ticker] terminate")
		fmt.Println("[ticker] start")
		defer wg.Done()
		tickN := 0

		ticker := time.NewTicker(time.Millisecond * 100)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("[ticker] ctx closed")
				return
			case <-ticker.C:
				tickN++
				log.Printf("[ticker] %d\n", tickN*100)
			}
		}
	}()
}

func runKafkaSuccess(ctx context.Context, prod sarama.AsyncProducer, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer fmt.Println("[sent success] terminate")
		fmt.Println("[sent success] start")
		defer wg.Done()

		successCh := prod.Successes()

		for {
			select {
			case <-ctx.Done():
				log.Println("[sent success] ctx closed")
				return
			case msg := <-successCh:
				if msg == nil {
					log.Println("[sent success] chan closed")
					return
				}
				log.Printf("[sent success] key: %q, partition: %d, offset: %d\n", msg.Key, msg.Partition, msg.Offset)
			}
		}
	}()
}

func runKafkaErrors(ctx context.Context, prod sarama.AsyncProducer, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer fmt.Println("[sent error] terminate")
		fmt.Println("[sent error] start")
		defer wg.Done()

		errCh := prod.Errors()

		for {
			select {
			case <-ctx.Done():
				log.Println("[sent error] ctx closed")
				return
			case msgErr := <-errCh:
				if msgErr == nil {
					log.Println("[sent error] chan closed")
					return
				}
				log.Printf("[sent error] err %s, topic: %q, offset: %d\n", msgErr.Err, msgErr.Msg.Topic, msgErr.Msg.Offset)
			}
		}
	}()
}

func runSignalHandler(ctx context.Context, wg *sync.WaitGroup) context.Context {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	sigCtx, cancel := context.WithCancel(ctx)

	wg.Add(1)
	go func() {
		defer fmt.Println("[signal] terminate")
		defer signal.Stop(sigterm)
		defer wg.Done()
		defer cancel()

		for {
			select {
			case sig, ok := <-sigterm:
				if !ok {
					fmt.Printf("[signal] signal chan closed: %s\n", sig.String())
					return
				}

				fmt.Printf("[signal] signal recv: %s\n", sig.String())
				return
			case _, ok := <-sigCtx.Done():
				if !ok {
					fmt.Println("[signal] context closed")
					return
				}

				fmt.Printf("[signal] ctx done: %s\n", ctx.Err().Error())
				return
			}
		}
	}()

	return sigCtx
}
