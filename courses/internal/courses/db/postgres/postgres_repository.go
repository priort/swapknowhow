package postgres

import (
	"context"
	"fmt"
	"os"
	"swapknowhow/courses/internal/courses"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresCoursesRepository struct {
	dbPool *pgxpool.Pool
	Close  func()
}

type PostgresConfig struct {
	User         string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

func NewPostgresCoursesRepository(config PostgresConfig) *PostgresCoursesRepository {
	connection := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		config.User, config.Password, config.Host, config.Port, config.DatabaseName)

	dbPool, err := pgxpool.Connect(context.Background(), connection)

	if err != nil {
		fmt.Println("db connection could not be established")
		fmt.Printf("%v", err)
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
