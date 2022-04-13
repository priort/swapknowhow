package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"swapknowhow/courses/internal/courses"
)

type PostgresCoursesRepository struct {
	dbPool *pgxpool.Pool
	Close  func()
}

func NewPostgresCoursesRepository() *PostgresCoursesRepository {
	connection := "user=postgres password=password host=localhost port=5432 dbname=local-coursesdb"
	dbPool, err := pgxpool.Connect(context.Background(), connection)

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
	rows, err := repo.dbPool.Query(context.Background(), "select * from courses;")
	if err != nil {
		fmt.Printf("Error querying courses in db: %v\n", err)
	}
	defer rows.Close()
	var retrievedCourses []courses.Course

	for rows.Next() {
		var course courses.Course
		err := rows.Scan(&course.Uuid, &course.Created, &course.Name, &course.Rating, &course.Descripton, &course.DurationMillis)
		if err != nil {
			fmt.Printf("Error parsing course row %v\n", err)
		}
		retrievedCourses = append(retrievedCourses, course)
	}
	return retrievedCourses

}

func (repo *PostgresCoursesRepository) CreateCourse(course courses.Course) {
	insert := `INSERT INTO courses(uuid, created, name, rating, description, duration_millis) 
			   VALUES (gen_random_uuid(), now(), $1, $2, $3, $4 );`
	_, err := repo.dbPool.Exec(context.Background(), insert, course.Name, course.Rating, course.Descripton, course.DurationMillis)
	if err != nil {
		fmt.Printf("Error inserting course into db: %v\n", err)
	}
}
