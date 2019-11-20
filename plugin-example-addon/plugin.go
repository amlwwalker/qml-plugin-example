package main

import (
	"fmt"
	"os"

	_ "github.com/amlwwalker/qml-plugin-example/plugin-example-addon/settings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

//this function is not needed, but it's still usefull for testing
func Init() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.Engine().AddImportPath("qrc:/qml")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	var mainComponent = qml.NewQQmlComponent2(view.Engine(), nil)
	mainComponent.ConnectStatusChanged(func(status qml.QQmlComponent__Status) {
		if status == qml.QQmlComponent__Ready {

			var item = quick.NewQQuickItemFromPointer(mainComponent.Create(view.RootContext()).Pointer()) //create item and "cast" it to QQuickItem
			item.SetParent(view)                                                                          //add invisible item to invisible parent (for auto-deletion ...)
			item.SetParentItem(view.ContentItem())                                                        //add visible item to visible parent

		} else {
			fmt.Println("failed with status:", status)
			for _, e := range mainComponent.Errors() {
				fmt.Println("error:", e.ToString())
			}
		}
	})

	qmlString := `
import QtQuick 2.0
import Settings 1.0

Settings {
	anchors.fill: parent
}
`

	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())

	view.Show()

	widgets.QApplication_Exec()
}
