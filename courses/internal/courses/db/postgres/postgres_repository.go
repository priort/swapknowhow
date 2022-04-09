package postgres

import (
	"context"
	"log"
	"os"
	"swapknowhow/courses/internal/courses"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresCoursesRepository struct {
	dbPool *pgxpool.Pool
	Close  func()
}

func NewPostgresCoursesRepository() *PostgresCoursesRepository {
	dbPool, err := pgxpool.Connect(context.Background(), "postgresql://localhost:5432/coursesdb")

	if err != nil {
		log.Fatal("db connection could not be established")
		os.Exit(1)
	}

	return &PostgresCoursesRepository{
		dbPool: dbPool,
		Close:  func() { dbPool.Close() },
	}

}

func (repo *PostgresCoursesRepository) GetCourses() []courses.Course {
	return make([]courses.Course, 0, 10)
}

func (repo *PostgresCoursesRepository) CreateCourse(course courses.Course) {
	// repo.dbPool.Exec(context.Background(), "INSERT INTO ")
}
