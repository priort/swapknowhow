package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"swapknowhow/courses/api"
	"swapknowhow/courses/internal/courses/db/postgres"
)

func main() {
	pgUser, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		log.Fatal("No POSTGRES_USER env variable")
	}
	pgPasswrod, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		log.Fatal("No POSTGRES_PASSWORD env variable")
	}
	pgHost, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		log.Fatal("No POSTGRES_HOST env variable")
	}
	pgPort, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		log.Fatal("No POSTGRES_PORT env variable")
	}
	pgPortInt, err := strconv.Atoi(pgPort)
	if err != nil {
		log.Fatal("No POSTGRES_PORT env variable must be a number")
	}
	pgDbName, exists := os.LookupEnv("POSTGRES_DB_NAME")
	if !exists {
		log.Fatal("No POSTGRES_DB_NAME env variable")
	}

	coursesApi := api.Api{CoursesRepo: postgres.NewPostgresCoursesRepository(
		postgres.PostgresConfig{
			User:         pgUser,
			Password:     pgPasswrod,
			Host:         pgHost,
			Port:         pgPortInt,
			DatabaseName: pgDbName})}

	//coursesApi := api.Api{CoursesRepo: postgres.NewPostgresCoursesRepository(
	//	postgres.PostgresConfig{
	//		User:         "postgres",
	//		Password:     "password",
	//		Host:         "localhost",
	//		Port:         5432,
	//		DatabaseName: "local-coursesdb"})}

	http.HandleFunc("/courses", coursesApi.Courses)
	log.Println("starting courses service on port 8082")
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("could not start courses service")
	}
}
