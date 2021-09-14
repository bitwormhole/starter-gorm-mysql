package mysql

import (
	"embed"

	startergorm "github.com/bitwormhole/starter-gorm"
	"github.com/bitwormhole/starter-gorm-mysql/src/main/etc"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/starter-gorm-mysql"
	myModuleVer  = "v0.0.3"
	myModuleRev  = 3
)

//go:embed src/main/resources
var theResFS embed.FS

// DriverModule 导出驱动模块：【github.com/bitwormhole/starter-gorm-mysql】
func DriverModule() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theResFS, "src/main/resources"))
	mb.OnMount(etc.ExportMySQLDriverConfig)

	mb.Dependency(startergorm.Module())

	return mb.Create()
}
