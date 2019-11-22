package main

// import (
// 	"os"
// 	"strings"
// 	"text/template"
// )

// func main() {
// 	qmlString := `
// 	import QtQuick 2.0
// 	Item {
// 		id: {{.Identity}}
// 		anchors.fill: parent
// 		Component.onCompleted: {
// 				var subComponent = Qt.createQmlObject(' \
// 				import {{.Name}} 1.0; \
// 				{{.Name}} { \
// 						width: parent.width; \
// 						height: parent.height;}', {{.Identity}});
// 				}
// 		}
// 		`
// 	var name string
// 	name = "Settings"
// 	var identity string
// 	identity = strings.ToLower(name)
// 	type templateObject struct {
// 		Name     string
// 		Identity string
// 	}
// 	var obj templateObject
// 	obj.Name = name
// 	obj.Identity = identity
// 	tmpl, err := template.New("qmlString").Parse(qmlString)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = tmpl.Execute(os.Stdout, obj)
// 	if err != nil {
// 		panic(err)
// 	}
// }
