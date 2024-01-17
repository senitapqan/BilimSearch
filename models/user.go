package models

type Student struct {
	Id       int
	UserId   int
}

type Teacher struct {
	Id       int
	UserId   int
}

type Roles struct {
	Id       int
	RoleName string
}

type User struct {
	Id       int
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Name     string `json:"name"     binding:"required"`
	Surname  string `json:"surname"  binding:"required"`
}

type RolesHeaders struct {
	Role string
	Id   int
}
