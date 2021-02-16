package controller

import (
	"GinFrame/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
function=查询所有员工, methods=["GET"]
interface=/api/staffs
*/
func GetStaffs(c *gin.Context) {
	var where interface{}
	var staffs model.Staffs
	staffs=model.FindStaffs(where)
	c.JSON(http.StatusOK, gin.H{
		"data":staffs,
		"message": "得到所有员工",
	})
}
