package main

import (
	"github.com/bitwormhole/starter"

	starter_gorm_mysql "github.com/bitwormhole/starter-gorm-mysql"
)

func main() {
	i := starter.InitApp()
	i.UseMain(starter_gorm_mysql.DemoModule())
	i.Run()
}
