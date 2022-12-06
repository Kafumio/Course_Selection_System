package router

import (
	"CSS/controllers"
	"CSS/global"
	"CSS/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(middleware.CORS())
	engine.Use(middleware.LoggerToFile())

	infoGroup := engine.Group("/info")
	{
		infoGroup.GET("/getCurrentTerm", controllers.GetCurrentTerm)
		infoGroup.GET("/getForbidCourseSelection", controllers.GetForbidCourseSelection)
	}

	stuGroup := engine.Group("/student")
	{
		stuGroup.POST("/addStudent", controllers.AddStudent)
		stuGroup.POST("/loginStu", controllers.LoginStu)
		stuGroup.POST("/findStuBySearch", controllers.FindStuBySearch)
		stuGroup.GET("/findStuById/:sid", controllers.FindStuById)
		stuGroup.GET("/findStuByPage/:page/:size", controllers.FindStuByPage)
		stuGroup.GET("/getLength", controllers.GetStuLength)
		stuGroup.GET("/deleteStuById/:sid", controllers.DeleteStuById)
		stuGroup.POST("/updateStu", controllers.UpdateStu)
	}

	teaGroup := engine.Group("/teacher")
	{
		teaGroup.POST("/addTeacher", controllers.AddTeacher)
		teaGroup.POST("/loginTea", controllers.LoginTea)
		teaGroup.GET("/findTeaById/:tid", controllers.FindTeaById)
		teaGroup.POST("/findTeaBySearch", controllers.FindTeaBySearch)
		teaGroup.GET("/deleteTeaById/:tid", controllers.DeleteTeaById)
		teaGroup.POST("/updateTea", controllers.UpdateTea)
	}

	couGroup := engine.Group("/course")
	{
		couGroup.POST("/findCouBySearch", controllers.FindCouBySearch)
		couGroup.GET("/findCouById/:cid", controllers.FindCouById)
		couGroup.POST("/addCourse", controllers.AddCourse)
		couGroup.GET("/deleteCouById/:cid", controllers.DeleteCouById)
		couGroup.POST("/updateCou", controllers.UpdateCou)
	}

	ctGroup := engine.Group("/courseTeacher")
	{
		ctGroup.GET("/insertCT/:cid/:tid/:term", controllers.InsertCT)
		ctGroup.GET("/findMyCourse/:tid/:term", controllers.FindMyCourse)
		ctGroup.GET("/findCTById/:cid/:tid/:term", controllers.FindCTById)
		ctGroup.POST("/findCTInfo", controllers.FindCTInfo)
		ctGroup.POST("/deleteCTById", controllers.DeleteCTById)
	}

	sctGroup := engine.Group("/SCT")
	{
		sctGroup.POST("/insertSCT", controllers.InsertSCT)
		sctGroup.GET("/findCTInfoBySid/:sid/:term", controllers.FindCTInfoBySid)
		sctGroup.GET("/findAllTerm", controllers.FindAllTerm)
		sctGroup.POST("/deleteBySCT", controllers.DeleteBySCT)
		sctGroup.POST("/findSCTBySearch", controllers.FindSCTBySearch)
		sctGroup.GET("/findSCTById/:sid/:cid/:tid/:term", controllers.FindSCTById)
		sctGroup.GET("/updateSCTById/:sid/:cid/:tid/:term/:grade", controllers.UpdateSCTById)
		sctGroup.GET("/deleteSCTById/:sid/:cid/:tid/:term", controllers.DeleteSCTById)
	}
	err := engine.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		fmt.Printf("init error:%v\n", err)
		return
	}
}
