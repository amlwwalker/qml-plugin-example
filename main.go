package main

import (
	"fmt"
	"os"
	"plugin"
	"sync"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	// "github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/qml"
)

//go:generate cp main.go ./plugin-example-addon
//go:generate qtrcc desktop ./plugin-example-addon
//go:generate qtmoc desktop ./plugin-example-addon
//go:generate qtminimal desktop ./plugin-example-addon
//go:generate go build -tags=minimal -buildmode=plugin -o ./plugin-example-addon/plugin.so ./plugin-example-addon
//go:generate rm ./plugin-example-addon/main.go

//go:generate qtmoc
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
	fmt.Printf("p %+v\r\n", p)
	if false { //this could be used to test the plugin, but it's not needed to call any initialization for the plugin to work with your app
		f, err := p.Lookup("Init")
		if err != nil {
			panic(err)
		}
		f.(func())()
	}
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	widgets.NewQApplication(len(os.Args), os.Args)
	app := qml.NewQQmlApplicationEngine(nil)
	app.AddImportPath("qrc:/qml/")
	app.Load(core.NewQUrl3("./qml/main.qml", 0))
	widgets.QApplication_Exec()

}
