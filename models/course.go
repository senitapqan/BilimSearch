package models

type Course struct {
	Id   int    `json:"course_id" db:"name"`
	Name string `json:"course_name" db:"name"`
}

type Lesson struct {
	Id        int    `json:"lesson_id"`
	CourseId  int    `json:"course_id"`
	TeacherId int    `json:"teacher_id"`
	Period    string `json:"period"`
}

type LessonItem struct {
	Id       int    `json:"lessonItem_id"`
	LessonId int    `json:"lesson_id"`
	Date     string `json:"date"`
}

type CoursesStudents struct {
	StudentId int
	CourseId  int
}

type LessonsStudents struct {
	StudentId int
	LessonId  int
}


