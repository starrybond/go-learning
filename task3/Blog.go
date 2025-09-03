package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
-- 1. 用户
INSERT INTO users (id, name, email, created_at, updated_at) VALUES
(1, 'Alice',  'alice@example.com',  NOW(), NOW()),
(2, 'Bob',    'bob@example.com',    NOW(), NOW());

-- 2. 文章
INSERT INTO posts (id, title, content, user_id, created_at, updated_at) VALUES
(1, 'Hello GORM',  'GORM quick start tutorial ...',  1, NOW(), NOW()),
(2, 'Go Context',  'Deep dive into context pkg ...', 1, NOW(), NOW()),
(3, 'MySQL Tips',  'Some MySQL performance tricks',  2, NOW(), NOW());

-- 3. 评论
INSERT INTO comments (id, content, user_id, post_id, created_at, updated_at) VALUES
(1, 'Great post!',                       2, 1, NOW(), NOW()),
(2, 'Looking forward to more articles.', 2, 1, NOW(), NOW()),
(3, 'Very helpful, thanks!',             1, 2, NOW(), NOW()),
(4, 'Could you add benchmarks?',         2, 2, NOW(), NOW()),
(5, 'Nice tips, saved my day!',          1, 3, NOW(), NOW());

*/

// ===================== 1. 定义模型 =====================

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:32;not null"`
	Email     string `gorm:"size:64;uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PostCount int    `gorm:"default:0"`
	Posts     []Post `gorm:"foreignKey:UserID"` // 一对多：User 拥有多篇 Post
}

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:128;not null"`
	Content       string `gorm:"type:text"`
	UserID        uint   // 外键，对应 User.ID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CommentStatus string    `gorm:"type:varchar(20);default:'有评论'"`
	Comments      []Comment `gorm:"foreignKey:PostID"` // 一对多：Post 拥有多条 Comment
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   // 评论作者
	PostID    uint   // 所属文章
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ===================== 2. 初始化数据库 =====================

func initDB() *gorm.DB {
	// 换成自己的账号密码
	dsn := "root:1234@tcp(127.0.0.1:3306)/blogdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // 打印 SQL
	})
	if err != nil {
		log.Fatalf("connect db failed: %v", err)
	}
	return db
}

// ===================== 3. 自动迁移建表 =====================
func migrate(db *gorm.DB) {
	// 会按照依赖顺序（先 User，再 Post，最后 Comment）自动建表
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Fatalf("migrate failed: %v", err)
	}
	fmt.Println("migrate success")
}

func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments"). // 把每篇文章的评论一并查出来
					Where("user_id = ?", userID).
					Order("id asc").
					Find(&posts).Error
	return posts, err
}

type PostWithCommentCount struct {
	Post
	CommentCount int
}

func GetPostWithMostComments(db *gorm.DB) ([]PostWithCommentCount, error) {
	var res []PostWithCommentCount
	sub := db.Model(&Comment{}).
		Select("post_id, COUNT(*) as comment_count").
		Group("post_id")

	err := db.Model(&Post{}).
		Select("posts.*, c.comment_count").
		Joins("JOIN (?) c ON posts.id = c.post_id", sub).
		Order("c.comment_count DESC").
		Limit(1). // 只要一条
		Scan(&res).Error
	return res, err
}

// 创建文章 → 用户文章数 +1
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// 删除评论 → 若文章已无评论则置状态为“无评论”
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var cnt int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			UpdateColumn("comment_status", "无评论").Error
	}
	return nil
}

func main() {
	// 1.模型定义
	db := initDB()
	//migrate(db)

	// 2.关联查询
	// 查询用户（id = 1）的所有文章及评论
	posts, err := GetUserPostsWithComments(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range posts {
		fmt.Printf("Post #%d: %s\n", p.ID, p.Title)
		for _, c := range p.Comments {
			fmt.Printf("  Comment #%d: %s\n", c.ID, c.Content)
		}
	}

	// 查询评论数量最多的文章（支持多条并列第一）
	pcs, err := GetPostWithMostComments(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, pc := range pcs {
		fmt.Printf("Most commented post: %s (comments: %d)\n",
			pc.Title, pc.CommentCount)
	}

	// 3. 现场验证钩子
	fmt.Println("-------- 测试 Post AfterCreate 钩子 --------")
	u := User{Name: "HookTester", Email: "hook@example.com"}
	db.Create(&u)
	fmt.Printf("创建用户后 PostCount = %d\n", u.PostCount) // 0

	p := Post{Title: "Hook Post", Content: "...", UserID: u.ID}
	db.Create(&p)

	var freshUser User
	db.First(&freshUser, u.ID)
	fmt.Printf("创建文章后 PostCount = %d\n", freshUser.PostCount) // 期望 1

	fmt.Println("-------- 测试 Comment AfterDelete 钩子 --------")
	c := Comment{Content: "first comment", UserID: u.ID, PostID: p.ID}
	db.Create(&c)

	// 查看文章状态
	db.First(&p, p.ID)
	fmt.Printf("有评论时 status = %q\n", p.CommentStatus) // 有评论

	db.Delete(&c) // 触发 AfterDelete

	db.First(&p, p.ID)
	fmt.Printf("删除评论后 status = %q\n", p.CommentStatus) // 无评论

}
