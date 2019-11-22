package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"plugin"
	"strings"
	"text/template"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
)

type CustomPlugin interface {
	Init()
	Name() string
	Description() string
}

func loadAPlugin(path string) CustomPlugin {

	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("p %+v\r\n", p)
	tmpPlugin, err := p.Lookup("Plugin")
	if err != nil {
		panic(err)
	}
	var plugin CustomPlugin
	plugin, ok := tmpPlugin.(CustomPlugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
	plugin.Init()
	// description := plugin.Description()
	name := plugin.Name()
	fmt.Println(name + " loaded")
	return plugin
}

func addToGui(c CustomPlugin, engine *qml.QQmlApplicationEngine) error {
	//https://stackoverflow.com/questions/31890372/add-objects-to-qml-layout-from-c%5D
	qmlTemplate := `
	import QtQuick 2.0
	Item {
		id: {{.Identity}}
		anchors.fill: parent
		Component.onCompleted: {
				var subComponent = Qt.createQmlObject(' \
				import {{.Name}} 1.0; \
				{{.Name}} { \
						width: parent.width; \
						height: parent.height;}', {{.Identity}});
				}
		}
		`
	type templateObject struct {
		Name     string
		Identity string
	}
	var obj templateObject
	if len(strings.Split(c.Name(), " ")) > 1 {
		return errors.New("name must be a single word")
	}
	obj.Name = strings.Title(strings.ToLower(c.Name()))
	obj.Identity = strings.ToLower(c.Name())
	tmpl, err := template.New("qmlString").Parse(qmlTemplate)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, obj)
	if err != nil {
		return err
	}
	var qmlString = buf.String()
	fmt.Println("qmlString ", qmlString)
	//get a reference to the object you want to put this as a child of
	stackLayout := engine.RootObjects()[0].FindChild("stackLayout", core.Qt__FindChildrenRecursively)
	stackLayoutPointer := quick.NewQQuickItemFromPointer(stackLayout.Pointer())
	//create a component to hold the child element
	mainComponent := qml.NewQQmlComponent2(engine, nil)
	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())
	//create the actual component as a qml item
	item := quick.NewQQuickItemFromPointer(mainComponent.Create(engine.RootContext()).Pointer())
	engine.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
	// //specify the parent
	item.SetParent(stackLayout)
	item.SetParentItem(stackLayoutPointer)
	return nil
}
