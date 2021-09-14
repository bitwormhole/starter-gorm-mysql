package main

import (
	"testing"

	mysql "github.com/bitwormhole/starter-gorm-mysql"
	"github.com/bitwormhole/starter/tests"
)

func TestMySQLDriver(t *testing.T) {

	i := tests.Starter(t)
	i.Use(mysql.DriverModule())
	i.Use(mysql.DriverModule4test())

	rt := i.RunTest()

	rt.Exit()
}
