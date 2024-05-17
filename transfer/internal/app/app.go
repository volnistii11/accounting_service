package app

import (
	"context"
	"fmt"
	consumer_group2 "github.com/volnistii11/accounting_service/transfer/internal/app/api/kafka/consumer_group"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	var (
		wg   = &sync.WaitGroup{}
		conf = newConfig(cliFlags)
		ctx  = runSignalHandler(context.Background(), wg)
	)

	fmt.Printf("%+v\n", conf)
	cg, err := consumer_group2.NewConsumerGroup(
		conf.kafka,
		conf.consumerGroup,
		consumer_group2.NewConsumerGroupHandler(),
		consumer_group2.WithOffsetsInitial(sarama.OffsetOldest),
	)
	if err != nil {
		return err
	}
	defer cg.Close()

	runCGErrorHandler(ctx, cg, wg)

	cg.Run(ctx, wg)

	wg.Wait()

	return nil
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

func runCGErrorHandler(ctx context.Context, cg sarama.ConsumerGroup, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case chErr, ok := <-cg.Errors():
				if !ok {
					fmt.Println("[cg-error] error: chan closed")
					return
				}

				fmt.Printf("[cg-error] error: %s\n", chErr)
			case <-ctx.Done():
				fmt.Printf("[cg-error] ctx closed: %s\n", ctx.Err().Error())
				return
			}
		}
	}()
}
