package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swapknowhow/courses/internal/courses"
)

func getCourses(writer http.ResponseWriter, req *http.Request) {
	coursesRepo := courses.NewInMemoryCoursesRepository()
	coursesJson, err := json.Marshal(coursesRepo.GetCourses())
	if err != nil {
		_ = fmt.Errorf("error marshalling courses %v", err)
		writer.WriteHeader(500)
	} else {
		writer.Write(coursesJson)
	}
}
