package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"redis/broker"
	"redis/repository"
	"redis/router"
	"time"
)

func main() {
	for range 20 {
		err := broker.InitConnector()
		if err != nil {
			fmt.Printf("init redis broker failed, err:%v\n", err)

			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("init redis broker success")
		break
	}

	rdb := repository.NewClient()

	r := router.NewRouter(rdb)

	fmt.Println("start redis server...")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("LOCAL_SERVICES_PORT")), r))
}
