package driver

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDriver 是 MySQL 的 starter-gorm 驱动
type MySQLDriver struct {
	markup.Component `class:"starter-gorm-driver-registry"`
}

func (inst *MySQLDriver) _Impl() (datasource.Driver, datasource.DriverRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *MySQLDriver) GetRegistration() *datasource.DriverRegistration {
	return &datasource.DriverRegistration{
		Driver: inst,
	}
}

// Accept 用于确定是否支持给定的配置
func (inst *MySQLDriver) Accept(cfg *datasource.Configuration) bool {
	name := cfg.Driver
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name == "mysql"
}

// Open 打开数据源
func (inst *MySQLDriver) Open(cfg *datasource.Configuration) (datasource.Source, error) {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//   dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if cfg == nil {
		return nil, errors.New("config==nil")
	}

	inst.prepareForDefaultPort(cfg)

	dsnbuilder := &strings.Builder{}
	dsnbuilder.WriteString(cfg.Username)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(cfg.Password)
	dsnbuilder.WriteString("@tcp(")
	dsnbuilder.WriteString(cfg.Host)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(strconv.Itoa(cfg.Port))
	dsnbuilder.WriteString(")/")
	dsnbuilder.WriteString(cfg.Database)
	dsnbuilder.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := dsnbuilder.String()

	// dialector := driver_pkg.Open(dsn)
	// if dialector == nil {
	// 	return nil, errors.New("driver_sqlserver.Open() return nil")
	// }

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &mysqlDataSource{db: db}, nil
}

func (inst *MySQLDriver) prepareForDefaultPort(cfg *datasource.Configuration) {
	const defport = 3306
	port := cfg.Port
	if port < 1 {
		port = defport
	}
	cfg.Port = port
}

////////////////////////////////////////////////////////////////////////////////

type mysqlDataSource struct {
	db *gorm.DB
}

func (inst *mysqlDataSource) _Impl() datasource.Source {
	return inst
}

func (inst *mysqlDataSource) DB() (*gorm.DB, error) {
	return inst.db, nil
}

func (inst *mysqlDataSource) Close() error {
	return nil
}
