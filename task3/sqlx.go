package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
-- 1. 建表
CREATE TABLE employees (
    id         INT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(50) NOT NULL,
    department VARCHAR(50) NOT NULL,
    salary     DECIMAL(12,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 2. 插入演示数据
INSERT INTO employees (name, department, salary) VALUES
('张三', '技术部',  9000.00),
('李四', '技术部', 12000.00),
('王五', '技术部', 15000.00),
('赵六', '技术部',  8500.00),
('钱七', '技术部', 11000.00),
('孙八', '财务部',  8000.00),
('周九', '人事部',  7500.00),
('吴十', '销售部', 25000.00);  -- 工资最高
*/

type Employee struct {
	ID         int64   `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func GetEmployeesInTech(db *sqlx.DB) ([]Employee, error) {
	var emps []Employee
	err := db.Select(&emps,
		`SELECT * FROM employees WHERE department = ?`, "技术部")
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesInTech failed: %w", err)
	}
	return emps, nil
}

func GetHighestPaidEmployee(db *sqlx.DB) (*Employee, error) {
	var emp Employee
	err := db.Get(&emp, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("GetHighestPaidEmployee failed: %w", err)
	}
	return &emp, nil
}

func main() {
	db, err := sqlx.Connect("mysql", "root:1234@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	techEmps, err := GetEmployeesInTech(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("技术部员工：%+v\n", techEmps)

	topEarner, err := GetHighestPaidEmployee(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("最高薪员工：%+v\n", topEarner)
}
