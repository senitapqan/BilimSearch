package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	attendanceTable            = "t_attendance"
	attendanceItemTable        = "t_attendance_item"
	tasksTable                 = "t_task"
	taskItemTable              = "t_task_item"
	coursesTable               = "t_courses"
	coursesStudentsTable       = "t_courses_students"
	coursesStudentsGradesTable = "t_courses_grades_students"
	lessonsTable               = "t_lessons"
	lessonItemTable            = "t_lesson_items"
	lessonStudentsTable        = "t_lesson_students"
	usersTable                 = "t_users"
	rolesTable                 = "t_roles"
	usersRolesTable            = "t_users_roles"
	studentsTable              = "t_students"
	teachersTable              = "t_teachers"
	postsTable                 = "t_posts"

	teacherRoleId              = 2
	studentRoleId              = 1
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
