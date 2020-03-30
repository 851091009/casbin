package databases

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PGDB *gorm.DB

var Typename = "postgres"
var Username = "postgres"
var Password = "feita2020"
var Host = "127.0.0.1"
var Port = 5432
var Database = "go"

func init() {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", Host, Port, Database, Username, Password)
	db, err := gorm.Open(Typename, dsn)

	// sql 打印

	db.LogMode(true)
	// 配置连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if err != nil {
		fmt.Printf("postgres sql connect error %v", err)
	}
	PGDB = db
}
