// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgmysqldemo

import (
	code0xb35759 "github.com/bitwormhole/starter-gorm-mysql/src/demo/golang/code"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDemo struct {
	instance *code0xb35759.Demo
	 markup0x23084a.Component `class:"life"`
	Sources datasource0x68a737.SourceManager `inject:"#starter-gorm-source-manager"`
}

