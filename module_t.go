package mysql

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

//go:embed src/test/resources
var theTestResFS embed.FS

// DriverModule4test 导出驱动模块：【github.com/bitwormhole/starter-gorm-mysql#test】
func DriverModule4test() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myModuleName + "#test").Version(myModuleVer).Revision(myModuleRev)
	mb.Dependency(DriverModule())
	mb.Resources(collection.LoadEmbedResources(&theTestResFS, "src/test/resources"))

	return mb.Create()
}
