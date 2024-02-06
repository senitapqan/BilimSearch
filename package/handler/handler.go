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
	APP := router.Group("/app")
	{
		APP.Use(h.userIdentify())	
		CRUD := APP.Group("/admin") 
		{
			CRUD.Use(h.userIdentify())
			teacherCRUD := CRUD.Group("/teacher") 
			{
				teacherCRUD.POST("/add", h.addTeacher) //done
				teacherCRUD.DELETE("/delete/:id", h.deleteTeacher) //done
				teacherCRUD.GET("/", h.getTeachers)
				teacherCRUD.GET("/:id", h.getTeacher)
			}

			courseCRUD := CRUD.Group("/course") 
			{
				courseCRUD.POST("/add", h.addCourse) //done
				courseCRUD.DELETE("/delete/:id", h.deleteCourse)
				courseCRUD.GET("/get", h.getCourses)
				courseCRUD.GET("/get/:id", h.getCourses)
			}

			lessonCRUD := CRUD.Group("/lesson")
			{
				lessonCRUD.POST("/add", h.addLesson) //done
				lessonCRUD.DELETE("/delete/:id", h.deleteLesson)
				lessonCRUD.GET("/get", h.getLessons)
				lessonCRUD.GET("/get/:id", h.getLesson)
			}	
		
			lessonItemCRUD := CRUD.Group("/lessonItem")
			{
				lessonItemCRUD.POST("/add", h.addLessonItem) 
				lessonItemCRUD.DELETE("/delete/:id", h.deleteLessonItem)
				lessonItemCRUD.GET("/get", h.getLessonItems)
				lessonItemCRUD.GET("/get/:id", h.getLessonItem)
			}
		}

		teacher := APP.Group("/teacher") 
		{
			lessonTeacher := teacher.Group("/lesson") 
			{
				lessonTeacher.GET("/", h.getTeacherLessons)
				lessonTeacher.GET("/:id", h.getTeacherLesson)
				lessonTeacher.GET("/:id/")
			}
			
		}

		home := APP.Group("/home") 
		{	
			home.GET("/grades", h.homeGrades) 
			home.GET("/schedule", h.homeSchedule)  //done
			home.GET("/attendance", h.homeAttendance)
			home.GET("/attendance/:id", h.attendanceItem)
	
			myCourses := home.Group("/my-courses") 
			{
				myCourses.GET("/", h.myCourse) //done
				myCourses.GET("/:id", h.myCourseItem) //done
				myCourses.GET("/:id/grades", h.myCourseGrades) //process
				myCourses.GET("/:id/task/:task_id", h.myCourseTask)
				myCourses.POST("/:id/task/:task_id/submit", h.myCourseTaskSend)
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
	}

	return router
} 

func NewHandler(services service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

