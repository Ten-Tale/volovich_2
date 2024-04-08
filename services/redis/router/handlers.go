package router

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"redis/domain"
	"strings"
)

func getStudents(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Get(context.Background(), "university.public.students:Struct{id=12345678}").Bytes()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Write(val)
	}
}

func getStudents2(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		students := r.URL.Query().Get("students")

		if len(students) == 0 {
			w.WriteHeader(400)

			w.Write([]byte("Bad request"))
			return
		}

		ctx := context.Background()

		var studentTemplate domain.Student
		studentList := []domain.Student{}

		for _, studentId := range strings.Split(students, ",") {
			student := rdb.Get(ctx, studentId)

			err := student.Err()

			if err != nil {
				log.Fatal(err)
			}

			studentResult, err := student.Bytes()

			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(studentResult, &studentTemplate)

			if err != nil {
				log.Fatal(err)
			}

			studentList = append(studentList, studentTemplate)
		}

		marshalledStudentList, err := json.Marshal(studentList)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(200)
		w.Write(marshalledStudentList)
	}
}
