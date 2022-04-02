package api

import (
	"io/ioutil"
	"net/http/httptest"

	"testing"
)

var stubbedCourses = `[{"Name":"my course","Rating":5,"Descripton":"a nice course about programming","DurationInSeconds":10}]`

func TestGetCourses(t *testing.T) {

	req := httptest.NewRequest("GET", "/courses", nil)
	recorder := httptest.NewRecorder()

	getCourses(recorder, req)

	res := recorder.Result()
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading from http response writer, %v", err)
	}
	actual := string(bytes)
	if actual != stubbedCourses {
		t.Errorf("Expected %v, got %v", stubbedCourses, actual)
	}

}
