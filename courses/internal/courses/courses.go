package courses

type Course struct {
	Name              string
	Rating            int
	Descripton        string
	DurationInSeconds int
}

type CoursesRepository interface {
	GetCourses() []Course
	CreateCourse(course Course)
}

type InMemoryCoursesRepository struct {
	courses []Course
}

func (r *InMemoryCoursesRepository) GetCourses() []Course {
	return r.courses
}

func (r *InMemoryCoursesRepository) CreateCourse(course Course) {
	r.courses = append(r.courses, course)
	println(r.courses)
}

func NewInMemoryCoursesRepository() *InMemoryCoursesRepository {
	return &InMemoryCoursesRepository{courses: make([]Course, 0, 10)}
}
