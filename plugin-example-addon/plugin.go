package main

import (
	"fmt"

	_ "github.com/amlwwalker/example-qml/qml-plugin-example/plugin-example-addon/settings"
)

var Plugin GuiPlugin

type GuiPlugin struct {
}

//this function is not needed, but it's still usefull for testing
func (c GuiPlugin) Init() {
	fmt.Println("loaded plugin example")
}

func (c GuiPlugin) Description() string {
	return "this plugin displays the settings for the application"
}

func (c GuiPlugin) Name() string {
	return "Settings"
}
