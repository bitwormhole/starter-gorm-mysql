// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgmysql

import (
	driver0x639c88 "github.com/bitwormhole/starter-gorm-mysql/driver"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-driver0x639c88.MySQLDriver
	cominfobuilder.Next()
	cominfobuilder.ID("com0-driver0x639c88.MySQLDriver").Class("starter-gorm-driver-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMySQLDriver{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMySQLDriver : the factory of component: com0-driver0x639c88.MySQLDriver
type comFactory4pComMySQLDriver struct {

    mPrototype * driver0x639c88.MySQLDriver

	

}

func (inst * comFactory4pComMySQLDriver) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMySQLDriver) newObject() * driver0x639c88.MySQLDriver {
	return & driver0x639c88.MySQLDriver {}
}

func (inst * comFactory4pComMySQLDriver) castObject(instance application.ComponentInstance) * driver0x639c88.MySQLDriver {
	return instance.Get().(*driver0x639c88.MySQLDriver)
}

func (inst * comFactory4pComMySQLDriver) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMySQLDriver) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMySQLDriver) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMySQLDriver) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMySQLDriver) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMySQLDriver) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




