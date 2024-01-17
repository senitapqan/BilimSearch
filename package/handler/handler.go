package handler

import (
	"BilimSearch/package/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth") 
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	home := router.Group("/home") 
	{	
		home.Use(h.userIdentify())
		home.GET("/grades", h.homeGrades)
		home.GET("/schedule", h.homeSchedule)
		home.GET("/attendance", h.homeAttendance)
		home.GET("/attendance/:id", h.attendanceItem)
	
		myCourses := home.Group("/my-courses") 
		{
			myCourses.GET("/", h.myCourse)
			myCourses.GET("/:id", h.myCourseItem)
			myCourses.GET("/:id/grades", h.myCourseGrades)
			myCourses.GET("/:id/task", h.myCourseTasks)
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

