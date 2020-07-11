package db

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilibs/gosql/v2"

	"shop_server_go/config"
)

var (
	username = config.GetDBValue("user")
	password = config.GetDBValue("password")
	host     = config.GetDBValue("host")
	port     = config.GetDBValue("port")
	database = config.GetDBValue("database")
)

func init() {
	configs := make(map[string]*gosql.Config)
	env := os.Getenv("ENV")
  configs["default"] = &gosql.Config{
        Enable:  true,
        Driver:  "mysql",
        Dsn: strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8&parseTime=True"}, ""),
				ShowSql: env == "development",
			}
	gosql.Connect(configs)
}
