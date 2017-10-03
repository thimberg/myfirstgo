package main

import (
    "fmt"
    "database/sql"
  _ "github.com/go-sql-driver/mysql"
//    "strings"
	"runtime"
)

type MYDB struct {
	db *sql.DB
	err error
}

func (n *MYDB) connect() {
	// DB接続
	db, err := sql.Open("mysql", "wpuser:asdffdsa1234@/wp_myblog")
	if err != nil {
		panic(err.Error())
	}
	
	n.db = db
	n.err = err
	runtime.SetFinalizer(n, disconnect)
}

func disconnect(n *MYDB)  {
	n.db.Close() // 関数がリターンする直前に呼び出される
	fmt.Println("db connect is closed..")
}

func (n MYDB) getMaxPostID() (int64) {
fmt.Println(n)
	// 最大PostIDを取得
	var maxCount int64
	rows, err := n.db.Query(`SELECT MAX(id) FROM wp_posts`,)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&maxCount); err != nil {
				panic(err.Error())
			}
	}
	rows.Close()
	
	return maxCount
}

func (n MYDB) insertPost(postId int64, title string, contents string) {
	// WP_POSTSへ登録
    _, err := n.db.Exec(`INSERT INTO wp_posts VALUES (?, ?, current_timestamp, (current_timestamp - interval 9 hour), ?, ?, '', 'publish', 'closed', 'closed', '', ?, '', '', current_timestamp, (current_timestamp - interval 9 hour), '', 0, ?, 0, 'post', '' ,0)`,
		postId,
		2,
		contents, 
		title, 
		postId, 
		"http://35.202.22.58//?p="+fmt.Sprint(postId),
	)

    if err != nil {
		panic(err.Error())
	}
}