package main

import (
	"os"
	"plugin"
	"sync"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

//go:generate go generate ./plugin-example-addon
//go:generate qtmoc
//go:generate qtminimal
//go:generate go build -tags=minimal

type Controller struct {
	core.QObject
	qApp *widgets.QApplication
}

var once sync.Once
var instance *Controller

func Instance() *Controller {
	once.Do(func() {
		instance = NewController(nil)
	})
	return instance
}

func main() {
	p, err := plugin.Open("plugin-example-addon/plugin.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("Init")
	if err != nil {
		panic(err)
	}
	f.(func())()
	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	Instance()
	Instance().qApp = qApp

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("./qml/main.qml", 0))
	widgets.QApplication_Exec()
}
