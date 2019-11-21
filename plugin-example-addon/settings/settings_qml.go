package settings

import (
	"fmt"

	"github.com/therecipe/qt/quick"
)

func init() {
	settingsController_QmlRegisterType2("Settings", 1, 0, "SettingsController")
	fmt.Println("settings has initialised")
}

type settingsController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//<-controller
	// _ func()                                                       `signal:"showImageLabel"`
	// _ func()                                                       `signal:"clearDetails"`
	// _ func(profileLabelText string)                                `signal:"showFileProfile"`
	_ func(id, name, description string) `signal:"updateSettings"`
	// _ func(date *core.QDate) `signal:"displayQDateDetails"`
}

func (s *settingsController) init() {

	//<-controller
	// controller.Instance().ConnectUpdateSettings(s.UpdateSettings)

}
