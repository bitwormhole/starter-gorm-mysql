package etc

import (
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"gorm.io/gorm"
)

// Demo ...
type Demo struct {
	markup.Component `class:"life"`

	Sources datasource.SourceManager `inject:"#starter-gorm-source-manager"`
}

func (inst *Demo) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *Demo) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.onStart,
	}
}

func (inst *Demo) onStart() error {

	s, err := inst.Sources.GetSource("default")
	if err != nil {
		return err
	}

	db, err := s.DB()
	if err != nil {
		return err
	}

	f := &Foo{}
	db.Find(&f)

	return db.Error
}

////////////////////////////////////////////////////////////////////////////////

type Foo struct {
	gorm.Model
}
