package mysql

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	driver_pkg "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Driver struct {
}

func (inst *Driver) Open(cfg *datasource.Configuration) (datasource.Source, error) {

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

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
	builder.Config1 = cfg
	builder.Config2 = gc
	builder.Dialector = dialector
	return builder.Create()
}
