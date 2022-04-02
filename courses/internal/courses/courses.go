package courses

type Course struct {
	Name              string
	Rating            int
	Descripton        string
	DurationInSeconds int
}

type CoursesRepository interface {
	getCourses() []Course
}

type InMemoryCoursesRepository struct{}

func (InMemoryCoursesRepository) GetCourses() []Course {
	return []Course{{Name: "my course", Rating: 5, Descripton: "a nice course about programming", DurationInSeconds: 10}}
}

func NewInMemoryCoursesRepository() *InMemoryCoursesRepository {
	return &InMemoryCoursesRepository{}
}
