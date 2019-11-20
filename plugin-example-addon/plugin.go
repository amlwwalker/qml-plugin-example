package main

import (
	"fmt"
	"os"

	_ "github.com/amlwwalker/example-qml/plugin-example-addon/settings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

//go:generate qtminimal
//go:generate go build -tags=minimal -buildmode=plugin -o plugin.so

func Init() {

	var view = quick.NewQQuickView(nil)
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

	var qmlString = `
import QtQuick 2.0
import QtQuick.Controls 2.0
Import Settings 1.0

Item {
  anchors.fill: parent //fill the view.ContentItem
  id: rootItem

  Button {
    text: "Button"

    anchors {
      top: parent.top
      left: parent.left
      right: parent.right
    }

    onClicked: Qt.createQmlObject('import QtQuick 2.0; Rectangle {color: "blue"; anchors {left: rootItem.left; right: rootItem.right; bottom: rootItem.bottom} height: 50}', rootItem, "dynamicSnippet1");
  }
}
`

	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())

	widgets.NewQApplication(len(os.Args), os.Args)

	widgets.NewQPushButton2("HELLO", nil).Show()

	widgets.QApplication_Exec()
}
