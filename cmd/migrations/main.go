package main

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
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

	driver, errD := postgres.WithInstance(sqlConnPool.DB, &postgres.Config{})
	if errD != nil {
		log.Fatal(errD)
	}

	// Financing
	mf, errf := migrate.NewWithDatabaseInstance(
		"file://sql/",
		"postgres",
		driver,
	)
	if errf != nil {
		log.Fatal(errf)
	}

	var err error
	for err == nil {
		err = mf.Steps(1)
	}
}
