package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	postgrestorage "github.com/vincentserpoul/gostorage/postgres"
)

func main() {
	log := logrus.New()

	conf := postgrestorage.Config{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "bench",
		Password: "bench",
		DbName:   "bench",
	}

	sqlConnPool, errQ := postgrestorage.NewConnPool(conf)
	if errQ != nil {
		log.Fatal(errQ.Error())
	}
	defer func() {
		if err := sqlConnPool.Close(); err != nil {
			log.Fatalf("connection pool close: %v", err)
		}
	}()
	log.Infof("Connected to %s DB", sqlConnPool.DriverName())

	http.HandleFunc("/postgres", sayHello)
	http.HandleFunc("/cockroach", sayHello)
	http.HandleFunc("/purego", sayHello)

	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Addr:         ":9013",
	}

	log.Infof("Listening on port 9013")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe: %s", err)
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}
