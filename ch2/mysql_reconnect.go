package main

/**
在非Web API环境下，GO程序需要长时间运行，客户端与MySQL服务器之间的TCP连接是不稳定的。

MySQL-Server会在一定时间内自动切断连接
GO程序遇到空闲期时长时间没有MySQL查询，MySQL-Server也会切断连接回收资源
其他情况，在MySQL服务器中执行kill process杀掉某个连接，MySQL服务器重启

解决连接 go away 最好的方式就是断线重连

模拟方法: 修改mysql wait_timeout 为10s,来验证程序是否支持了断线重连
       1. mysql -uroot -p123456
       2. show global variables like 'wait_timeout';   => 28800
       3. set global wati_timeout=10;
       4. 创建测试数据库db_order
       create database db_order;
       5. 创建测试数据表tb_order 包含 order_id,status  我们每隔12秒插入一条数据, 看是否能自动连接上
       CREATE TABLE `tb_order` (
  `order_id` varchar(11) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1

代码写完后 发现在go的sql driver实现中有自动重连机制,默认重连一次
 */

import "database/sql"

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(mysqlhost:3306)/db_order")
	checkErr(err)
	for {
		stmt, err := db.Prepare("INSERT INTO tb_order SET order_id=?")
		checkErr(err)
		res, err := stmt.Exec("orderid_1234")
		checkErr(err)
		time.Sleep(time.Second * 20)
		fmt.Println(res)
	}

	fmt.Printf("%v", err)

}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}
