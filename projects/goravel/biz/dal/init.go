package dal

import (
	"goravel/biz/dal/mysql"
	"goravel/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
