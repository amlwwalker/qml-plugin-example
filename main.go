package main

import (
	"os"
	"sync"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

//go:generate cp main.go ./plugin-example-addon
//go:generate cp pluginLoader.go ./plugin-example-addon
//go:generate qtrcc desktop ./plugin-example-addon
//go:generate qtmoc desktop ./plugin-example-addon
//go:generate qtminimal desktop ./plugin-example-addon
//go:generate go build -tags=minimal -buildmode=plugin -o ./plugin-example-addon/plugin.so ./plugin-example-addon
//go:generate rm ./plugin-example-addon/main.go
//go:generate rm ./plugin-example-addon/pluginLoader.go

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
	plugin := loadAPlugin("plugin-example-addon/plugin.so")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	widgets.NewQApplication(len(os.Args), os.Args)
	app := qml.NewQQmlApplicationEngine(nil)

	app.AddImportPath("qrc:/qml/")
	app.Load(core.NewQUrl3("./qml/main.qml", 0))

	if err := addToGui(plugin, app); err != nil {
		fmt.Println("error adding to gui ", err)
	}

	widgets.QApplication_Exec()

}
