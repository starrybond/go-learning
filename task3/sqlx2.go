package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int64   `db:"id"`     // 对应 books.id
	Title  string  `db:"title"`  // 对应 books.title
	Author string  `db:"author"` // 对应 books.author
	Price  float64 `db:"price"`  // 对应 books.price
}

// GetExpensiveBooks 查询价格 > minPrice 的书籍
func GetExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	// 复杂 SQL：多条件、排序、LIMIT 都可以往里加
	query := `
		SELECT id, title, author, price
		FROM books
		WHERE price > ?
		ORDER BY id
	`
	err := db.Select(&books, query, minPrice)
	return books, err
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	books, err := GetExpensiveBooks(db, 50) // 价格 > 50
	if err != nil {
		panic(err)
	}

	for _, b := range books {
		fmt.Printf("%3d | %-25s | %-20s | %.2f\n", b.ID, b.Title, b.Author, b.Price)
	}
}
