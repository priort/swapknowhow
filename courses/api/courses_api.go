package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"swapknowhow/courses/internal/courses"
)

type Api struct {
	CoursesRepo courses.CoursesRepository
}

func (api *Api) Courses(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		api.getCourses(writer, req)
	case "POST":
		api.createCourse(writer, req)
	default:
		writer.Write([]byte("Invalid method"))
		writer.WriteHeader(400)
	}
}

func (api *Api) createCourse(writer http.ResponseWriter, req *http.Request) {
	courseJson, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("error reading request body %v\n", err)
		writer.WriteHeader(500)
		return
	}
	var course courses.Course
	err = json.Unmarshal(courseJson, &course)
	if err != nil {
		fmt.Printf("error deserializing course %v \n", err)
		writer.WriteHeader(500)
		return
	}
	api.CoursesRepo.CreateCourse(course)
	writer.WriteHeader(201)
}

func (api *Api) getCourses(writer http.ResponseWriter, _ *http.Request) {
	coursesJson, err := json.Marshal(api.CoursesRepo.GetCourses())
	if err != nil {
		fmt.Printf("error marshalling courses %v", err)
		writer.WriteHeader(500)
	} else {
		_, err := writer.Write(coursesJson)
		if err != nil {
			fmt.Printf("error writing response %v", err)
		}
	}
}
