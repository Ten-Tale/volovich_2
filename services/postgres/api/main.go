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
	for range 10 {
		err := broker.InitConnector()
		if err != nil {
			fmt.Printf("init postgres broker failed, err:%v\n", err)

			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("init postgres broker success")
		break
	}

	db, err := repository.NewClient()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
	}

	r := router.NewRouter(db)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("LOCAL_SERVICES_PORT")), r))
}
