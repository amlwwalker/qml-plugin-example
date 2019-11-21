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

	//now prep to load the plugin.
	view := quick.NewQQuickViewFromPointer(app.RootObjects()[0].Pointer())

	stackLayout := view.RootObject().FindChild("stackLayout", core.Qt__FindChildrenRecursively)

	mainComponent := qml.NewQQmlComponent2(app, nil)
	mainComponent.ConnectStatusChanged(func(status qml.QQmlComponent__Status) {
		if status == qml.QQmlComponent__Ready {

			item := quick.NewQQuickItemFromPointer(mainComponent.Create(app.RootContext()).Pointer()) //create item and "cast" it to QQuickItem
			app.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
			item.SetParent(stackLayout) //add invisible item to invisible parent (for auto-deletion ...)
			item.SetParentItem(view.ContentItem())

		} else {
			fmt.Println("failed with status:", status)
			for _, e := range mainComponent.Errors() {
				fmt.Println("error:", e.ToString())
			}
		}
	})

	qmlString := `
		import QtQuick 2.0
		Item {
			id: rootItem
			anchors.fill: parent
			Component.onCompleted: {
				var subComponent = Qt.createQmlObject(' \
				import Settings 1.0; \
				Settings { \
					id: settings; \
					width: parent.width; \
					height: parent.height;}', rootItem);
			}
		}
	`
	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())
	// view.Show()
	widgets.QApplication_Exec()

}
