package v1

import (
	"fmt"
	"ginbblog/dao"
	"ginbblog/model"
	"ginbblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddComment 新增评论
func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	fmt.Println(data)
	code := dao.AddComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetComment 获取单个评论信息
func GetComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := dao.GetComment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := dao.DeleteComment(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCommentCount 获取评论数量
func GetCommentCount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	total := dao.GetCommentCount(id)
	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

// GetCommentList 后台查询评论列表
func GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := dao.GetCommentList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetCommentListFront 展示页面显示评论列表
func GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := dao.GetCommentListFront(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

// CheckComment 通过审核
func CheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code := dao.CheckComment(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// UncheckComment 撤下评论审核
func UncheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code := dao.UncheckComment(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
