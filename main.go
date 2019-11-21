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

	//now prep to load the plugin.
	// view := quick.NewQQuickView(nil)
	view := quick.NewQQuickViewFromPointer(app.RootObjects()[0].Pointer())
	// view.SetSource(core.NewQUrl3("./qml/main.qml", 0))
	// view.Show()
	//https://stackoverflow.com/questions/31890372/add-objects-to-qml-layout-from-c%5D
	qmlString := `
	import QtQuick 2.0
	Rectangle {
		id: settings
		anchors.fill: parent
		width: 200
		height: 200
		color: "blue";
		Component.onCompleted: {
				var subComponent = Qt.createQmlObject(' \
				import Settings 1.0; \
				Settings { \
						width: parent.width; \
						height: parent.height;}', settings);
				}
		}
		`
	stackLayout := view.RootObject().FindChild("stackLayout", core.Qt__FindChildrenRecursively)
	stackLayoutPointer := quick.NewQQuickItemFromPointer(stackLayout.Pointer())
	mainComponent := qml.NewQQmlComponent2(app, stackLayout)
	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl3("./qml/main.qml", 0))
	item := quick.NewQQuickItemFromPointer(mainComponent.Create(app.RootContext()).Pointer())
	app.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
	item.SetParentItem(stackLayoutPointer)
	// item := quick.NewQQuickItemFromPointer(mainComponent.Create(app.RootContext()).Pointer()) //create item and "cast" it to QQuickItem
	// app.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
	// item.SetParent(stackLayout)
	// item.SetParentItem(stackLayoutPointer)
	// mainComponent.ConnectStatusChanged(func(status qml.QQmlComponent__Status) {
	// 	if status == qml.QQmlComponent__Ready {

	// 		stackLayoutPointer := quick.NewQQuickItemFromPointer(stackLayout.Pointer())
	// 		item := quick.NewQQuickItemFromPointer(mainComponent.Create(app.RootContext()).Pointer()) //create item and "cast" it to QQuickItem
	// 		app.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
	// 		item.SetParent(stackLayout) //add invisible item to invisible parent (for auto-deletion ...)
	// 		item.SetParentItem(stackLayoutPointer)

	// 	} else {
	// 		fmt.Println("failed with status:", status)
	// 		for _, e := range mainComponent.Errors() {
	// 			fmt.Println("error:", e.ToString())
	// 		}
	// 	}
	// })
	// mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())
	widgets.QApplication_Exec()

}
