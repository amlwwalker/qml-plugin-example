package main

import (
	"fmt"
	"os"
	"plugin"
	"sync"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
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
	if true { //this could be used to test the plugin, but it's not needed to call any initialization for the plugin to work with your app
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

	//https://stackoverflow.com/questions/31890372/add-objects-to-qml-layout-from-c%5D
	qmlString := `
	import QtQuick 2.0
	Item {
		id: settings
		anchors.fill: parent
		Component.onCompleted: {
				var subComponent = Qt.createQmlObject(' \
				import Settings 1.0; \
				Settings { \
						width: parent.width; \
						height: parent.height;}', settings);
				}
		}
		`

	//get a reference to the object you want to put this as a child of
	stackLayout := app.RootObjects()[0].FindChild("stackLayout", core.Qt__FindChildrenRecursively)
	stackLayoutPointer := quick.NewQQuickItemFromPointer(stackLayout.Pointer())
	//create a component to hold the child element
	mainComponent := qml.NewQQmlComponent2(app, nil)
	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())
	//create the actual component as a qml item
	item := quick.NewQQuickItemFromPointer(mainComponent.Create(app.RootContext()).Pointer())
	app.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
	// //specify the parent
	item.SetParent(stackLayout)
	item.SetParentItem(stackLayoutPointer)

	widgets.QApplication_Exec()

}
