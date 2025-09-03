package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsn    = "root:1234@tcp(127.0.0.1:3306)/"
	dbName = "school"
)

func main() {
	// 1. 先连 root 库，创建 school 数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接MySQL失败：%v", err)
	}
	defer db.Close()
	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName); err != nil {
		log.Fatalf("创建数据库失败：%v", err)
	}
	// 2. 再连 school 库
	dsnSchool := dsn + dbName + "?parseTime=true&charset=utf8mb4"
	db2, err := sql.Open("mysql", dsnSchool)
	if err != nil {
		log.Fatalf("连接school库失败：%v", err)
	}
	defer db2.Close()
	// 3. 创建 students 表
	ddl := `
	CREATE TABLE IF NOT EXISTS students (
		id    INT AUTO_INCREMENT PRIMARY KEY,
		name  VARCHAR(50) NOT NULL,
		age   INT,
		grade VARCHAR(20)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	if _, err := db2.Exec(ddl); err != nil {
		log.Fatalf("创建表失败：%v", err)
	}
	// 4.插入记录
	if _, err := db2.Exec(
		`INSERT INTO students(name,age,grade) VALUES(?,?,?)`, "张三", 20, "三年级"); err != nil {
		log.Fatalf("插入数据失败：%v", err)
	}
	// 5. 查询 age > 18
	rows, err := db2.Query("SELECT * FROM students WHERE age > 18")
	if err != nil {
		log.Fatalf("查询失败：%v", err)
	}
	defer rows.Close()
	fmt.Println("年龄>18的学生：")
	for rows.Next() {
		var id, age int
		var name, grade string
		if err := rows.Scan(&id, &name, &age, &grade); err != nil {
			log.Fatal("查询失败: %v", err)
		}
		fmt.Printf("id:%d name:%s age:%d grade:%s\n", id, name, age, grade)
	}
	// 6. 更新张三的年级
	if _, err := db2.Exec("UPDATE students SET grade = ? WHERE name = ?", "四年级", "张三"); err != nil {
		log.Fatalf("更新失败: %v", err)
	}
	// 7. 删除 age < 15
	if _, err := db2.Exec("DELETE FROM students WHERE age < 15"); err != nil {
		log.Fatalf("删除失败: %v", err)
	}
	// 8. 打印最终数据
	fmt.Println("\n最终students表：")
	finalRows, err := db2.Query("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer finalRows.Close()
	for finalRows.Next() {
		var id, age int
		var name, grade string
		if err := finalRows.Scan(&id, &name, &age, &grade); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id:%d name:%s age:%d grade:%s\n", id, name, age, grade)
	}

}
