package models

type Student struct {
	Id       int
	Username string
	Password string
	Email    string
	Name     string
	Surname  string
	RoleId   int
}

type Teacher struct {
	Id       int
	Username string
	Password string
	Email    string
	Name     string
	Surname  string
	RoleId   int
}

type Roles struct {
	Id       int
	RoleName string
}
