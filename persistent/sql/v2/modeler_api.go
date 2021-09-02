package sqlv2

import (
	sqlOculi "github.com/ravielze/oculi/persistent/sql"
)

func (i *Impl) RegisterObject(obj ...interface{}) sqlOculi.API {
	i.Object = append(i.Object, obj...)
	return i
}

func (i *Impl) ObjectFunction(onInstall func(), onReset func()) (install func(), reset func()) {
	install = func() {
		i.Database.AutoMigrate(i.Object...)
		if onInstall != nil {
			onInstall()
		}
	}
	reset = func() {
		if len(i.Object) > 0 {
			for x := (len(i.Object) - 1); x >= 0; x-- {
				i.Database.Migrator().DropTable(i.Object[x])
			}
		}
		if onReset != nil {
			onReset()
		}
	}
	return
}
