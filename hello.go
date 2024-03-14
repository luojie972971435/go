package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // 连接 MySQL 数据库
    db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/pgone_act")
    if err != nil {
        fmt.Println("数据库连接失败:", err)
        return
    }
    defer db.Close()

    // 查询数据
    rows, err := db.Query("SELECT * FROM user_agent_level")
    if err != nil {
        fmt.Println("查询数据失败:", err)
        return
    }
    defer rows.Close()

    // 遍历查询结果
    for rows.Next() {
        var id int
        var uid int
        var pid string
        if err := rows.Scan(&id, &uid, &pid); err != nil {
            fmt.Println("读取数据失败:", err)
            return
        }
        fmt.Printf("ID: %d, Uid: %d, Pid: %s\n", id, uid, pid)
    }

    // 检查遍历过程中是否发生错误
    if err := rows.Err(); err != nil {
        fmt.Println("遍历数据失败:", err)
        return
    }
}