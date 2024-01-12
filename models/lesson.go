package models

type Attendance struct {
	Id           int
	LessonItemId int
}

type AttendanceItem struct {
	AttendanceItemId int
	StudentId        int
	Absence          bool
}

type Task struct {
	Id           int
	LessonItemId int
}

type TaskItem struct {
	TaskId    int
	StudentId int
	Grade     int
}
