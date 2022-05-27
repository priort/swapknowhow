package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"swapknowhow/courses/internal/courses"
	"testing"
)

func TestCoursesService(t *testing.T) {
	os.Setenv("COURSES_SERVICE_HOST", "localhost")
	os.Setenv("COURSES_SERVICE_PORT", "8082")
	coursesServiceHost, exists := os.LookupEnv("COURSES_SERVICE_HOST")
	if !exists {
		t.Fatal("No COURSES_SERVICE_HOST environment variable")
	}
	coursesServicePort, exists := os.LookupEnv("COURSES_SERVICE_PORT")
	if !exists {
		t.Fatal("No COURSES_SERVICE_PORT environment variable")
	}
	courseToCreate := courses.Course{
		Name:           "My test course",
		Rating:         5,
		Descripton:     "This is a test course",
		DurationMillis: 199,
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%s/courses", coursesServiceHost, coursesServicePort))
	if err != nil {
		t.Fatalf("error from getting courses %v", err)
	}
	if err != nil {
		t.Fatalf("error reading get courses response body %v", err)
	}

	coursesResponse := callCoursesEndpointWithGet(t, coursesServiceHost, coursesServicePort)
	fmt.Println(coursesResponse)
}

func callCoursesEndpointWithGet(t *testing.T, coursesServiceHost string, coursesServicePort string) []courses.Course {
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/courses", coursesServiceHost, coursesServicePort))
	if err != nil {
		t.Fatalf("error from getting courses %v", err)
	}
	if err != nil {
		t.Fatalf("error reading get courses response body %v", err)
	}
	var coursesResponse []courses.Course
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &coursesResponse)
	if err != nil {
		t.Fatalf("error unmarshalling get courses response body %v", err)
	}
	return coursesResponse
}
