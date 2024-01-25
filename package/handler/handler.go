package handler

import (
	"BilimSearch/package/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func (h *Handler) getIds(role string, c *gin.Context) (int, int, error) {
	userId, err := getUserId(c)
	if err != nil {
		return -1, -1, err;
	}
	roleId, err := getRoleId(c, role)
	if err != nil {
		return -1, -1, nil
	}
	return userId, roleId, err;
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth") 
	{
		auth.POST("/sign-in", h.signIn) //done
		auth.POST("/sign-up", h.signUp)	//done
	}

	CRUD := router.Group("/admin") 
	{
		CRUD.Use(h.userIdentify())
		teacherCRUD := CRUD.Group("/teacher") 
		{
			teacherCRUD.POST("/add", h.addTeacher)
			teacherCRUD.DELETE("/delete/:id", h.deleteTeacher)
			teacherCRUD.GET("/", h.getTeachers)
			teacherCRUD.GET("/:id", h.getTeacher)
		}

		courseCRUD := CRUD.Group("/course") 
		{
			courseCRUD.POST("/add")
			courseCRUD.DELETE("/delete/:id")
			courseCRUD.GET("/get")
			courseCRUD.GET("/get/:id")
		}

		lessonCRUD := CRUD.Group("/lesson")
		{
			lessonCRUD.POST("/add")
			lessonCRUD.DELETE("/delete/:id")
			lessonCRUD.GET("/get")
			lessonCRUD.GET("/get/:id")
		}		
	}

	home := router.Group("/home") 
	{	
		home.Use(h.userIdentify())
		home.GET("/grades", h.homeGrades) 
		home.GET("/schedule", h.homeSchedule)  //done
		home.GET("/attendance", h.homeAttendance)
		home.GET("/attendance/:id", h.attendanceItem)
	
		myCourses := home.Group("/my-courses") 
		{
			myCourses.GET("/", h.myCourse) //done
			myCourses.GET("/:id", h.myCourseItem) //done
			myCourses.GET("/:id/grades", h.myCourseGrades) //process
			myCourses.GET("/:id/task/:task_id", h.myCourseTaskItem)
			myCourses.POST("/:id/task/:task_id/submit", h.myCourseTaskItemSend)
		}

		register := home.Group("/course-reg") 
		{
			register.GET("/", h.regCourse)
			register.DELETE("/:id", h.delCourseItem)

			register.GET("/all", h.regCourseAllItems)
			register.GET("/:id", h.regCourseItem)
			register.POST("/:id", h.regCoursePickLesson)
		}
	}
	

	return router
} 

func NewHandler(services service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

