package controller

// 「评论模块」的控制器
// CreateComment – 对指定文章发评论（需登录）
// ListComment – 获取指定文章的所有评论（公开）

import (
	"net/http"
	"strconv"

	"blog/model"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type CommentReq struct {
	Content string `json:"content" binding:"required"`
}

func CreateComment(c *gin.Context) {
	var req CommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pid, _ := strconv.Atoi(c.Param("id"))
	uid := c.GetUint("userID")
	comm := model.Comment{Content: req.Content, UserID: uid, PostID: uint(pid)}
	if err := model.DB.Create(&comm).Error; err != nil {
		utils.L.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": comm.ID})
}

func ListComment(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("id"))
	var comms []model.Comment
	//SELECT * FROM comments WHERE post_id = 1;
	//SELECT * FROM users WHERE id IN (2,3,4);
	model.DB.Where("post_id = ?", pid).Preload("User").Find(&comms)
	c.JSON(http.StatusOK, gin.H{"data": comms})
}
