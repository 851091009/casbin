package casbin

import (
	"database/sql"
	Orm "feita/databases"
	"feita/helpers"
	"fmt"
	"github.com/casbin/casbin/v2"
	casbinpgadapter "github.com/cychiuae/casbin-pg-adapter"
)

var Enforcer *casbin.Enforcer

func CasbinInit() {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", Orm.Host, Orm.Port, Orm.Database, Orm.Username, Orm.Password)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Printf("casbin postgres sql connect error %v", err)
	}

	tableName := helpers.CasbinTableName // 设置数据姓名
	adapter, err := casbinpgadapter.NewAdapter(db, tableName)

	if err != nil {
		fmt.Printf("casbin NewAdapter error %v", err)
	}

	enforcer, err := casbin.NewEnforcer("expansion/casbin/model.conf", adapter)

	if err != nil {
		fmt.Printf("casbin NewEnforcer error %v", err)
	}

	Enforcer = enforcer
}
