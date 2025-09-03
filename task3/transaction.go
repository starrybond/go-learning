package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
建表语句
CREATE TABLE accounts (
id      INT PRIMARY KEY AUTO_INCREMENT,
balance DECIMAL(18,2) NOT NULL CHECK (balance >= 0)
);

CREATE TABLE transactions (
id             INT PRIMARY KEY AUTO_INCREMENT,
from_account_id INT NOT NULL,
to_account_id   INT NOT NULL,
amount          DECIMAL(18,2) NOT NULL CHECK (amount > 0),
created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (from_account_id) REFERENCES accounts(id),
FOREIGN KEY (to_account_id)   REFERENCES accounts(id)
);

-- 任选客户端执行
INSERT INTO accounts(balance) VALUES (500);
INSERT INTO accounts(balance) VALUES (0);
*/

func transfer(db *sql.DB, fromID, toID int64, amount float64) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx err:%w", err)
	}
	defer func() {
		if rbErr := tx.Rollback(); rbErr != nil {
			err = fmt.Errorf("rollback err:%w", rbErr)
		}
	}()

	var balance float64
	if err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&balance); err != nil {
		return fmt.Errorf("select err:%w", err)
	}
	if balance < amount {
		return fmt.Errorf("insufficient balance:%f", balance)
	}
	if _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID); err != nil {
		return fmt.Errorf("update err:%w", err)
	}
	if _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID); err != nil {
		return fmt.Errorf("update err:%w", err)
	}
	if _, err = tx.ExecContext(ctx, "INSERT INTO transactions(from_account_id, to_account_id, amount) VALUES (?, ?, ?);", fromID, toID, amount); err != nil {
		return fmt.Errorf("insert err:%w", err)
	}
	return tx.Commit()
}

func printAccounts(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var balance float64
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id:%d balance:%f\n", id, balance)
	}
}
func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/testdb?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	printAccounts(db)
	if err := transfer(db, 1, 2, 100); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("transfer success")
	}
	printAccounts(db)
}
