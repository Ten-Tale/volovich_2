package broker

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func InitConnector() error {
	fileData, err := os.ReadFile("redis-connector.json")
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%s/connectors", os.Getenv("KAFKA_CONNECT_ADDR"), os.Getenv("KAFKA_CONNECT_PORT")),
		"application/json",
		bytes.NewReader(fileData),
	)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("Redis sink connector responsed with: %s", resp.Status))
	}

	return nil
}

// psql -U postgres -d university
// insert into public.students(id, name, surname) values('12345678', 'name', 'surname');
