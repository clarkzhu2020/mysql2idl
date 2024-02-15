package dal

import (
	"grom/biz/dal/mysql"
	"grom/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
