package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"redis/repository"
	"redis/router"
	"time"
)

func init() {
	time.Sleep(100 * time.Millisecond)
}

func main() {
	rdb := repository.NewClient()

	r := router.NewRouter(rdb)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("LOCAL_SERVICES_PORT")), r))
}
