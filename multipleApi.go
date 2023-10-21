package main

import (
	"context"
	"os"
	"time"

	"github.com/go-kit/log"
	metahttp "github.com/onmetahq/meta-http/pkg/meta_http"
)

func createHttpClient(url string, logger log.Logger) metahttp.Requests {
	return metahttp.NewClient(url, logger, 15*time.Second)
}
func initLogger() log.Logger {
	logger := log.NewJSONLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	return logger
}

func fetchEmployees(client metahttp.Requests) (EmployeeResponse, error) {
	ctx := context.Background()
	headeres := map[string]string{}
	res := EmployeeResponse{}
	_, err := client.Get(ctx, "/todos/1", headeres, &res)

	return res, err
}

type EmployeeResponse struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func syncCall(client metahttp.Requests) time.Duration {
	startTime := time.Now()
	for i := 0; i < 5; i++ {
		fetchEmployees(client)
	}
	return time.Since(startTime)
}
func AsyncfetchEmployees(client metahttp.Requests, out chan time.Duration) {}

func AsyncCall(Client metahttp.Requests) time.Duration {
	startTime := time.Now()
	channel := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(ch chan bool) {
			fetchEmployees(Client)
			ch <- true
		}(channel)

	}
	for i := 0; i < 5; i++ {
		<-channel
	}
	return time.Since(startTime)

}

func main() {

	logger := initLogger()
	client := createHttpClient("https://jsonplaceholder.typicode.com", logger)
	syncDuration := syncCall(client)
	AsyncDuration := AsyncCall(client)
	logger.Log("message", "finished", "durationSync", syncDuration, "durationAsync", AsyncDuration)
}
