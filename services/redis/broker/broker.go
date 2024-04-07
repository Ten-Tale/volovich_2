package broker

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const connectorConfig = "{" +
	"\"name\": \"redis-sink\",\n" +
	"\"config\": {\n" +
	"\"connector.class\": \"com.github.jcustenborder.kafka.connect.redis.RedisSinkConnector\",\n" +
	"\"tasks.max\": \"1\",\n" +
	"\"topics\": \"university.public.students\",\n" +
	"\"\": \"redis:6379\",\n" +
	"\"redis.uri\": \"redis://redis:6379\",\n" +
	"\"key.converter\": \"org.apache.kafka.connect.storage.StringConverter\",\n" +
	"\"value.converter\": \"org.apache.kafka.connect.json.JsonConverter\",\n" +
	"\"value.converter.schemas.enable\": \"false\"\n" +
	"}\n" +
	"}"

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
