package mysql

import (
	"embed"

	startergorm "github.com/bitwormhole/starter-gorm"
	"github.com/bitwormhole/starter-gorm-mysql/gen/cfgmysql"
	"github.com/bitwormhole/starter-gorm-mysql/gen/cfgmysqldemo"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "github.com/bitwormhole/starter-gorm-mysql"
	theModuleVersion  = "v0.0.5"
	theModuleRevision = 5
	theModuleResPath  = "src/main/resources"
)

//go:embed src/main/resources
var theModuleResFS embed.FS

// DriverModule 导出[starter-gorm-mysql]模块
func DriverModule() application.Module {

	mb := &application.ModuleBuilder{}

	mb.Name(theModuleName + "#driver").Version(theModuleVersion).Revision(theModuleRevision)
	mb.OnMount(cfgmysql.ExportConfigForMySQL)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))
	mb.Dependency(startergorm.Module())

	return mb.Create()
}

// DemoModule 导出[#demo]模块
func DemoModule() application.Module {

	mb := &application.ModuleBuilder{}

	mb.Name(theModuleName + "#demo").Version(theModuleVersion).Revision(theModuleRevision)
	mb.OnMount(cfgmysqldemo.ExportConfigForMySQLDemo)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))
	mb.Dependency(DriverModule())

	return mb.Create()
}
