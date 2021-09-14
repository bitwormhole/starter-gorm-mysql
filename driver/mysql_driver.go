package driver

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	driver_pkg "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDriver 是 mysql 的 starter-gorm 驱动
type MySQLDriver struct {
	markup.Component
}

func (inst *MySQLDriver) _Impl() datasource.Driver {
	return inst
}

func (inst *MySQLDriver) Accept(cfg *datasource.Configuration) bool {
	name := cfg.Driver
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return (name == "mysql")
}

func (inst *MySQLDriver) prepareForDefaultPort(cfg *datasource.Configuration) {
	const defport = 3306
	port := cfg.Port
	if port < 1 {
		port = defport
	}
	cfg.Port = port
}

func (inst *MySQLDriver) Open(cfg *datasource.Configuration) (datasource.Source, error) {

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if cfg == nil {
		return nil, errors.New("config==nil")
	}

	inst.prepareForDefaultPort(cfg)

	dsnb := &strings.Builder{}
	dsnb.WriteString(cfg.Username)
	dsnb.WriteString(":")
	dsnb.WriteString(cfg.Password)
	dsnb.WriteString("@tcp(")
	dsnb.WriteString(cfg.Host)
	dsnb.WriteString(":")
	dsnb.WriteString(strconv.Itoa(cfg.Port))
	dsnb.WriteString(")/")
	dsnb.WriteString(cfg.Database)
	dsnb.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := dsnb.String()

	dialector := driver_pkg.Open(dsn)
	if dialector == nil {
		return nil, errors.New("driver_sqlserver.Open() return nil")
	}

	gc := &gorm.Config{}

	builder := &datasource.GormDataSourceBuilder{}
	builder.Config1 = *cfg
	builder.Config2 = *gc
	builder.Dialector = dialector
	return builder.Open()
}
