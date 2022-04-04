package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"swapknowhow/courses/internal/courses"
)

var coursesRepo = courses.NewInMemoryCoursesRepository()

func CreateCourse(writer http.ResponseWriter, req *http.Request) {
	courseJson, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte("could not read from request body"))
		writer.WriteHeader(500)
	}
	var course courses.Course
	json.Unmarshal(courseJson, &course)
	coursesRepo.CreateCourse(course)
	writer.WriteHeader(201)
}

func GetCourses(writer http.ResponseWriter, req *http.Request) {
	coursesJson, err := json.Marshal(coursesRepo.GetCourses())
	if err != nil {
		_ = fmt.Errorf("error marshalling courses %v", err)
		writer.WriteHeader(500)
	} else {
		writer.Write(coursesJson)
	}
}
