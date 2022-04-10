package courses

import (
	"github.com/gofrs/uuid"
	"time"
)

type Course struct {
	Uuid           uuid.UUID
	Created        time.Time
	Name           string
	Rating         int
	Descripton     string
	DurationMillis int
}

type CoursesRepository interface {
	GetCourses() []Course
	CreateCourse(course Course)
}
