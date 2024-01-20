package dtos

type CourseGrades struct {
	CourseName    string `json:"course name"`
	PrefinalGrade int    `json:"prefinal grade"`
	FinalGrade    int    `json:"final grade"`
}
