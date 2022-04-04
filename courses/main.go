package main

import (
	"log"
	"net/http"
	"swapknowhow/courses/api"
)

func main() {
	http.HandleFunc("/courses", func(writer http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			api.GetCourses(writer, req)
		case "POST":
			api.CreateCourse(writer, req)
		default:
			writer.Write([]byte("Invalid method"))
			writer.WriteHeader(400)
		}
	})
	log.Println("starting courses service on port 8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("could not start courses service")
	}
}
