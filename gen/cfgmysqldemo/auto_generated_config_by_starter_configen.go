// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgmysqldemo

import (
	code0xb35759 "github.com/bitwormhole/starter-gorm-mysql/src/demo/golang/code"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
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

	// component: com0-code0xb35759.Demo
	cominfobuilder.Next()
	cominfobuilder.ID("com0-code0xb35759.Demo").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo : the factory of component: com0-code0xb35759.Demo
type comFactory4pComDemo struct {

    mPrototype * code0xb35759.Demo

	
	mSourcesSelector config.InjectionSelector

}

func (inst * comFactory4pComDemo) init() application.ComponentFactory {

	
	inst.mSourcesSelector = config.NewInjectionSelector("#starter-gorm-source-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo) newObject() * code0xb35759.Demo {
	return & code0xb35759.Demo {}
}

func (inst * comFactory4pComDemo) castObject(instance application.ComponentInstance) * code0xb35759.Demo {
	return instance.Get().(*code0xb35759.Demo)
}

func (inst * comFactory4pComDemo) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Sources = inst.getterForFieldSourcesSelector(context)
	return context.LastError()
}

//getterForFieldSourcesSelector
func (inst * comFactory4pComDemo) getterForFieldSourcesSelector (context application.InstanceContext) datasource0x68a737.SourceManager {

	o1 := inst.mSourcesSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.SourceManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-code0xb35759.Demo")
		eb.Set("field", "Sources")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.SourceManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




