package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/danieeelfr/my-large-json-file-reader/internal/repository"
	"github.com/danieeelfr/my-large-json-file-reader/internal/service"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("starting the application...")
	err := godotenv.Load("./../local.env")
	if err != nil {
		_ = godotenv.Load("./local.env")
	}
}

func main() {
	// TODO: this gracefulShutdown can be moved to a pkg folder
	// wait for termination signal and register database & http server clean-up operations
	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"redis": func(ctx context.Context) error {
			fmt.Print("redis shutdown")
			return nil
		},
		"file-dispose": func(ctx context.Context) error {
			fmt.Print("redis shutdown")
			return nil
		},
		// Add other cleanup operations here
	})
	cfg := config.NewConfig()

	repo := repository.NewRepository(cfg)
	srv := service.NewPortService(cfg)

	ports, err := srv.ReadAndDecode()
	if err != nil {
		log.Fatal(err)
	}

	if err := repo.Set(ports.Records); err != nil {
		log.Fatal(err)
	}

	// TODO
	// result, err := repo.Get("", 100)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	<-wait

	fmt.Println("finished without errors!")

}

// TODO: this operation can be moved to a pkg folder
// operation is a clean up function on shutting down
type operation func(ctx context.Context) error

// TODO: this gracefulShutdown can be moved to a pkg folder
// gracefulShutdown waits for termination syscalls and doing clean up operations after received it
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
