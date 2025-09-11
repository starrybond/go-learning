package controller

import (
	"blog/model"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 「博客文章」模块的 CRUD 控制器

type PostReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 发表文章
func CreatePost(c *gin.Context) {
	var req PostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := c.GetUint("userID")
	post := model.Post{Title: req.Title, Content: req.Content, UserID: uid}
	if err := model.DB.Create(&post).Error; err != nil {
		utils.L.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": post.ID})
}

// 公开文章列表
func ListPost(c *gin.Context) {
	var posts []model.Post
	model.DB.Preload("User").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// 查询文章详情
func GetPost(c *gin.Context) {
	// 从 URL /api/posts/:id 里取出字符串 ID
	id := c.Param("id")
	var post model.Post
	//SELECT * FROM posts WHERE id = ? LIMIT 1;
	//SELECT * FROM users WHERE id IN (1);  -- 1 就是 post.user_id
	if err := model.DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	uid := c.GetUint("userID")
	var post model.Post
	// SELECT * FROM posts WHERE id = ? LIMIT 1;
	if err := model.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if post.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
		return
	}
	var req PostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.DB.Model(&post).Updates(map[string]interface{}{"title": req.Title, "content": req.Content})
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	uid := c.GetUint("userID")
	var post model.Post
	if err := model.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if post.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
		return
	}
	model.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
