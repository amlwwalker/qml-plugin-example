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

type PluginManager struct {
	engine *qml.QQmlApplicationEngine
}

func (p *PluginManager) Init(engine *qml.QQmlApplicationEngine) {
	p.engine = engine
}

// const viewTemplate = `
// import QtQuick 2.0
// Item {
// 	id: {{.Identity}}
// 	anchors.fill: parent
// 	Component.onCompleted: {
// 			var subComponent = Qt.createQmlObject(' \
// 			import {{.Name}} 1.0; \
// 			{{.Name}} { \
// 					width: parent.width; \
// 					height: parent.height;}', {{.Identity}});
// 			}
// 	}
// 	`

const viewTemplate = `
import {{.Name}} 1.0;
{{.Name}} {
	// anchors.fill: parent
}
`
const tabTemplate = `
import QtQuick.Controls 2.4;
TabButton {
	text: qsTr("{{.Title}}")
}
`

func InitialisePlugin(path string) CustomPlugin {

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

/*
Next is to have a more generic function

and add this;

      TabButton {
          text: qsTr("Settings")
			}

*/
func (p *PluginManager) InitialisePlugin(c CustomPlugin) error {
	if err := p.configureTitle(c); err != nil {
		fmt.Println("error configureing title bar")
		return err
	}
	if err := p.configureView(c); err != nil {
		fmt.Println("error configureing statusview")
		return err
	}
	return nil
}
func (p *PluginManager) configureTitle(c CustomPlugin) error {
	type templateObject struct {
		Title string
	}
	var obj templateObject
	obj.Title = strings.Title(strings.ToLower(c.Name()))
	tmpl, err := template.New("qmlTitle").Parse(tabTemplate)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, obj)
	if err != nil {
		return err
	}
	var qmlString = buf.String()
	fmt.Println("adding tabbar ", qmlString)
	p.addToGui("tabBar", qmlString)
	return nil
}
func (p *PluginManager) configureView(c CustomPlugin) error {
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
	tmpl, err := template.New("qmlStackView").Parse(viewTemplate)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, obj)
	if err != nil {
		return err
	}
	var qmlString = buf.String()
	fmt.Println("adding view ", qmlString)
	p.addToGui("stackLayout", qmlString)
	return nil
}
func (p *PluginManager) addToGui(childElement, qmlString string) error {
	//https://stackoverflow.com/questions/31890372/add-objects-to-qml-layout-from-c%5D
	if p.engine == nil {
		return errors.New("Engine not intialised")
	}
	//get a reference to the object you want to put this as a child of
	stackLayout := p.engine.RootObjects()[0].FindChild(childElement, core.Qt__FindChildrenRecursively)
	stackLayoutPointer := quick.NewQQuickItemFromPointer(stackLayout.Pointer())
	//create a component to hold the child element
	mainComponent := qml.NewQQmlComponent2(p.engine, nil)
	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())
	//create the actual component as a qml item
	item := quick.NewQQuickItemFromPointer(mainComponent.Create(p.engine.RootContext()).Pointer())
	p.engine.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
	// //specify the parent
	// item.SetParent(stackLayout)
	// item.SetParentItem(stackLayoutPointer)
	if childElement == "stackLayout" {
		stackLayout.InvokeMethod("addToStackLayout", core.NewQVariant1(item))
	} else {
		item.SetParent(stackLayout)
		item.SetParentItem(stackLayoutPointer)
	}
	return nil
}
