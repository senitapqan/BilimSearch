package models

//models
type Course struct {
	Id   int    `json:"id" db:"name"`
	Name string `json:"name" db:"name"`
}

type Lesson struct {
	Id        int
	CourseId  int
	TeacherId int
	Period    string
}

type LessonItem struct {
	Id       int
	LessonId int
	Date     string
}

type CoursesStudents struct {
	StudentId int
	CourseId  int
}

type LessonsStudents struct {
	StudentId int
	LessonId  int
}
