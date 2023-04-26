package main

import (
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/store"
)

func main() {
	srv := service.New(service.Name("example"))
	srv.Init()

	records, err := store.Read("mykey", func(r *store.ReadOptions) {
		r.Table = "example"
	})
	if err != nil {
		fmt.Println("Error reading from store: ", err)
	}

	if len(records) == 0 {
		fmt.Println("No records")
	}
	for _, record := range records {
		fmt.Printf("key: %v, value: %v\n", record.Key, string(record.Value))
	}

	time.Sleep(1 * time.Hour)
}
