package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"swapknowhow/courses/internal/courses"
	"testing"
)

var stubbedCourses = `[{"Name":"my course","Rating":5,"Descripton":"a nice course about programming","DurationInSeconds":10}]`

func TestCanCreateAndRetrieveCourses(t *testing.T) {

	courseToCreate := courses.Course{
		Name:              "test course",
		Rating:            5,
		Descripton:        "course to test",
		DurationInSeconds: 50,
	}
	jsonBody, _ := json.Marshal(courseToCreate)
	body := strings.NewReader(string(jsonBody))
	courseCreationRequest := httptest.NewRequest("POST", "/courses", body)
	courseCreationResponseRecorder := httptest.NewRecorder()

	CreateCourse(courseCreationResponseRecorder, courseCreationRequest)
	courseCreationResponse := courseCreationResponseRecorder.Result()
	defer courseCreationResponse.Body.Close()

	if courseCreationResponse.StatusCode != 201 {
		t.Errorf("expected course courseCreationResponse status: %v actual: %v", 201, courseCreationResponse.StatusCode)
	}

	req := httptest.NewRequest("GET", "/courses", nil)
	recorder := httptest.NewRecorder()

	GetCourses(recorder, req)

	res := recorder.Result()
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading from http response writer, %v", err)
	}
	var coursesResponse []courses.Course
	json.Unmarshal(bytes, &coursesResponse)
	if len(coursesResponse) != 1 {
		t.Errorf("Expected coursesResponse length :%v, got %v", 1, len(coursesResponse))
	}

	course := (coursesResponse)[0]
	if course != courseToCreate {
		t.Errorf("Expected :%v, got %v", courseToCreate, course)
	}
}
