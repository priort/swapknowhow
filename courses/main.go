package main

import (
	"log"
	"net/http"
	"swapknowhow/courses/api"
	"swapknowhow/courses/internal/courses/db/postgres"
)

func main() {
	coursesApi := api.Api{CoursesRepo: postgres.NewPostgresCoursesRepository()}
	http.HandleFunc("/courses", coursesApi.Courses)
	log.Println("starting courses service on port 8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("could not start courses service")
	}
}
