package settings

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_quick "github.com/therecipe/qt/quick"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type settingsController_ITF interface {
	std_quick.QQuickItem_ITF
	settingsController_PTR() *settingsController
}

func (ptr *settingsController) settingsController_PTR() *settingsController {
	return ptr
}

func (ptr *settingsController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (ptr *settingsController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickItem_PTR().SetPointer(p)
	}
}

func PointerFromSettingsController(ptr settingsController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.settingsController_PTR().Pointer()
	}
	return nil
}

func NewSettingsControllerFromPointer(ptr unsafe.Pointer) (n *settingsController) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(settingsController)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *settingsController:
			n = deduced

		case *std_quick.QQuickItem:
			n = &settingsController{QQuickItem: *deduced}

		default:
			n = new(settingsController)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *settingsController) Init() { this.init() }

//export callbacksettingsController15bc46_Constructor
func callbacksettingsController15bc46_Constructor(ptr unsafe.Pointer) {
	this := NewSettingsControllerFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbacksettingsController15bc46_UpdateSettings
func callbacksettingsController15bc46_UpdateSettings(ptr unsafe.Pointer, id C.struct_Moc_PackedString, name C.struct_Moc_PackedString, description C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateSettings"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(id), cGoUnpackString(name), cGoUnpackString(description))
	}

}

func (ptr *settingsController) ConnectUpdateSettings(f func(id string, name string, description string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateSettings") {
			C.settingsController15bc46_ConnectUpdateSettings(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "updateSettings")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateSettings"); signal != nil {
			f := func(id string, name string, description string) {
				(*(*func(string, string, string))(signal))(id, name, description)
				f(id, name, description)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateSettings", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateSettings", unsafe.Pointer(&f))
		}
	}
}

func (ptr *settingsController) DisconnectUpdateSettings() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DisconnectUpdateSettings(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateSettings")
	}
}

func (ptr *settingsController) UpdateSettings(id string, name string, description string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.settingsController15bc46_UpdateSettings(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_Moc_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

func settingsController_QRegisterMetaType() int {
	return int(int32(C.settingsController15bc46_settingsController15bc46_QRegisterMetaType()))
}

func (ptr *settingsController) QRegisterMetaType() int {
	return int(int32(C.settingsController15bc46_settingsController15bc46_QRegisterMetaType()))
}

func settingsController_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.settingsController15bc46_settingsController15bc46_QRegisterMetaType2(typeNameC)))
}

func (ptr *settingsController) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.settingsController15bc46_settingsController15bc46_QRegisterMetaType2(typeNameC)))
}

func settingsController_QmlRegisterType() int {
	return int(int32(C.settingsController15bc46_settingsController15bc46_QmlRegisterType()))
}

func (ptr *settingsController) QmlRegisterType() int {
	return int(int32(C.settingsController15bc46_settingsController15bc46_QmlRegisterType()))
}

func settingsController_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.settingsController15bc46_settingsController15bc46_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *settingsController) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.settingsController15bc46_settingsController15bc46_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *settingsController) __childItems_atList(i int) *std_quick.QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := std_quick.NewQQuickItemFromPointer(C.settingsController15bc46___childItems_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *settingsController) __childItems_setList(i std_quick.QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46___childItems_setList(ptr.Pointer(), std_quick.PointerFromQQuickItem(i))
	}
}

func (ptr *settingsController) __childItems_newList() unsafe.Pointer {
	return C.settingsController15bc46___childItems_newList(ptr.Pointer())
}

func (ptr *settingsController) __grabTouchPoints_ids_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.settingsController15bc46___grabTouchPoints_ids_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *settingsController) __grabTouchPoints_ids_setList(i int) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46___grabTouchPoints_ids_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *settingsController) __grabTouchPoints_ids_newList() unsafe.Pointer {
	return C.settingsController15bc46___grabTouchPoints_ids_newList(ptr.Pointer())
}

func (ptr *settingsController) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.settingsController15bc46___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *settingsController) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *settingsController) __children_newList() unsafe.Pointer {
	return C.settingsController15bc46___children_newList(ptr.Pointer())
}

func (ptr *settingsController) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.settingsController15bc46___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *settingsController) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *settingsController) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.settingsController15bc46___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *settingsController) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.settingsController15bc46___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *settingsController) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *settingsController) __findChildren_newList() unsafe.Pointer {
	return C.settingsController15bc46___findChildren_newList(ptr.Pointer())
}

func (ptr *settingsController) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.settingsController15bc46___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *settingsController) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *settingsController) __findChildren_newList3() unsafe.Pointer {
	return C.settingsController15bc46___findChildren_newList3(ptr.Pointer())
}

func (ptr *settingsController) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.settingsController15bc46___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *settingsController) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *settingsController) __qFindChildren_newList2() unsafe.Pointer {
	return C.settingsController15bc46___qFindChildren_newList2(ptr.Pointer())
}

func NewSettingsController(parent std_quick.QQuickItem_ITF) *settingsController {
	settingsController_QRegisterMetaType()
	tmpValue := NewSettingsControllerFromPointer(C.settingsController15bc46_NewSettingsController(std_quick.PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbacksettingsController15bc46_DestroySettingsController
func callbacksettingsController15bc46_DestroySettingsController(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~settingsController"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).DestroySettingsControllerDefault()
	}
}

func (ptr *settingsController) ConnectDestroySettingsController(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~settingsController"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~settingsController", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~settingsController", unsafe.Pointer(&f))
		}
	}
}

func (ptr *settingsController) DisconnectDestroySettingsController() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~settingsController")
	}
}

func (ptr *settingsController) DestroySettingsController() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DestroySettingsController(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *settingsController) DestroySettingsControllerDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DestroySettingsControllerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbacksettingsController15bc46_ActiveFocusChanged
func callbacksettingsController15bc46_ActiveFocusChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "activeFocusChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbacksettingsController15bc46_ActiveFocusOnTabChanged
func callbacksettingsController15bc46_ActiveFocusOnTabChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "activeFocusOnTabChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbacksettingsController15bc46_AntialiasingChanged
func callbacksettingsController15bc46_AntialiasingChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "antialiasingChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbacksettingsController15bc46_BaselineOffsetChanged
func callbacksettingsController15bc46_BaselineOffsetChanged(ptr unsafe.Pointer, vqr C.double) {
	if signal := qt.GetSignal(ptr, "baselineOffsetChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(vqr))
	}

}

//export callbacksettingsController15bc46_ChildMouseEventFilter
func callbacksettingsController15bc46_ChildMouseEventFilter(ptr unsafe.Pointer, item unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "childMouseEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_quick.QQuickItem, *std_core.QEvent) bool)(signal))(std_quick.NewQQuickItemFromPointer(item), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsControllerFromPointer(ptr).ChildMouseEventFilterDefault(std_quick.NewQQuickItemFromPointer(item), std_core.NewQEventFromPointer(event)))))
}

func (ptr *settingsController) ChildMouseEventFilterDefault(item std_quick.QQuickItem_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.settingsController15bc46_ChildMouseEventFilterDefault(ptr.Pointer(), std_quick.PointerFromQQuickItem(item), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbacksettingsController15bc46_ChildrenRectChanged
func callbacksettingsController15bc46_ChildrenRectChanged(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childrenRectChanged"); signal != nil {
		(*(*func(*std_core.QRectF))(signal))(std_core.NewQRectFFromPointer(vqr))
	}

}

//export callbacksettingsController15bc46_ClassBegin
func callbacksettingsController15bc46_ClassBegin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "classBegin"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *settingsController) ClassBeginDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_ClassBeginDefault(ptr.Pointer())
	}
}

//export callbacksettingsController15bc46_ClipChanged
func callbacksettingsController15bc46_ClipChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "clipChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbacksettingsController15bc46_ComponentComplete
func callbacksettingsController15bc46_ComponentComplete(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "componentComplete"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *settingsController) ComponentCompleteDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_ComponentCompleteDefault(ptr.Pointer())
	}
}

//export callbacksettingsController15bc46_ContainmentMaskChanged
func callbacksettingsController15bc46_ContainmentMaskChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "containmentMaskChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_Contains
func callbacksettingsController15bc46_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QPointF) bool)(signal))(std_core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsControllerFromPointer(ptr).ContainsDefault(std_core.NewQPointFFromPointer(point)))))
}

func (ptr *settingsController) ContainsDefault(point std_core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.settingsController15bc46_ContainsDefault(ptr.Pointer(), std_core.PointerFromQPointF(point))) != 0
	}
	return false
}

//export callbacksettingsController15bc46_DragEnterEvent
func callbacksettingsController15bc46_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *settingsController) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbacksettingsController15bc46_DragLeaveEvent
func callbacksettingsController15bc46_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *settingsController) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbacksettingsController15bc46_DragMoveEvent
func callbacksettingsController15bc46_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *settingsController) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbacksettingsController15bc46_DropEvent
func callbacksettingsController15bc46_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *settingsController) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbacksettingsController15bc46_EnabledChanged
func callbacksettingsController15bc46_EnabledChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enabledChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_Event
func callbacksettingsController15bc46_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsControllerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(ev)))))
}

func (ptr *settingsController) EventDefault(ev std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.settingsController15bc46_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbacksettingsController15bc46_FocusChanged
func callbacksettingsController15bc46_FocusChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "focusChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbacksettingsController15bc46_FocusInEvent
func callbacksettingsController15bc46_FocusInEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewSettingsControllerFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *settingsController) FocusInEventDefault(vqf std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(vqf))
	}
}

//export callbacksettingsController15bc46_FocusOutEvent
func callbacksettingsController15bc46_FocusOutEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewSettingsControllerFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *settingsController) FocusOutEventDefault(vqf std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(vqf))
	}
}

//export callbacksettingsController15bc46_GeometryChanged
func callbacksettingsController15bc46_GeometryChanged(ptr unsafe.Pointer, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChanged"); signal != nil {
		(*(*func(*std_core.QRectF, *std_core.QRectF))(signal))(std_core.NewQRectFFromPointer(newGeometry), std_core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewSettingsControllerFromPointer(ptr).GeometryChangedDefault(std_core.NewQRectFFromPointer(newGeometry), std_core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *settingsController) GeometryChangedDefault(newGeometry std_core.QRectF_ITF, oldGeometry std_core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_GeometryChangedDefault(ptr.Pointer(), std_core.PointerFromQRectF(newGeometry), std_core.PointerFromQRectF(oldGeometry))
	}
}

//export callbacksettingsController15bc46_HeightChanged
func callbacksettingsController15bc46_HeightChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "heightChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_HoverEnterEvent
func callbacksettingsController15bc46_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverEnterEvent"); signal != nil {
		(*(*func(*std_gui.QHoverEvent))(signal))(std_gui.NewQHoverEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).HoverEnterEventDefault(std_gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *settingsController) HoverEnterEventDefault(event std_gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_HoverEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQHoverEvent(event))
	}
}

//export callbacksettingsController15bc46_HoverLeaveEvent
func callbacksettingsController15bc46_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QHoverEvent))(signal))(std_gui.NewQHoverEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).HoverLeaveEventDefault(std_gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *settingsController) HoverLeaveEventDefault(event std_gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_HoverLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQHoverEvent(event))
	}
}

//export callbacksettingsController15bc46_HoverMoveEvent
func callbacksettingsController15bc46_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverMoveEvent"); signal != nil {
		(*(*func(*std_gui.QHoverEvent))(signal))(std_gui.NewQHoverEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).HoverMoveEventDefault(std_gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *settingsController) HoverMoveEventDefault(event std_gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_HoverMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQHoverEvent(event))
	}
}

//export callbacksettingsController15bc46_ImplicitHeightChanged
func callbacksettingsController15bc46_ImplicitHeightChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "implicitHeightChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_ImplicitWidthChanged
func callbacksettingsController15bc46_ImplicitWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "implicitWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_InputMethodEvent
func callbacksettingsController15bc46_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *settingsController) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbacksettingsController15bc46_InputMethodQuery
func callbacksettingsController15bc46_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewSettingsControllerFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *settingsController) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.settingsController15bc46_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbacksettingsController15bc46_IsTextureProvider
func callbacksettingsController15bc46_IsTextureProvider(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isTextureProvider"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsControllerFromPointer(ptr).IsTextureProviderDefault())))
}

func (ptr *settingsController) IsTextureProviderDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.settingsController15bc46_IsTextureProviderDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbacksettingsController15bc46_KeyPressEvent
func callbacksettingsController15bc46_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *settingsController) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbacksettingsController15bc46_KeyReleaseEvent
func callbacksettingsController15bc46_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *settingsController) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbacksettingsController15bc46_MouseDoubleClickEvent
func callbacksettingsController15bc46_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *settingsController) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbacksettingsController15bc46_MouseMoveEvent
func callbacksettingsController15bc46_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *settingsController) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbacksettingsController15bc46_MousePressEvent
func callbacksettingsController15bc46_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *settingsController) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbacksettingsController15bc46_MouseReleaseEvent
func callbacksettingsController15bc46_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *settingsController) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbacksettingsController15bc46_MouseUngrabEvent
func callbacksettingsController15bc46_MouseUngrabEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseUngrabEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *settingsController) MouseUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_MouseUngrabEventDefault(ptr.Pointer())
	}
}

//export callbacksettingsController15bc46_OpacityChanged
func callbacksettingsController15bc46_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_ParentChanged
func callbacksettingsController15bc46_ParentChanged(ptr unsafe.Pointer, vqq unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "parentChanged"); signal != nil {
		(*(*func(*std_quick.QQuickItem))(signal))(std_quick.NewQQuickItemFromPointer(vqq))
	}

}

//export callbacksettingsController15bc46_ReleaseResources
func callbacksettingsController15bc46_ReleaseResources(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "releaseResources"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *settingsController) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_ReleaseResourcesDefault(ptr.Pointer())
	}
}

//export callbacksettingsController15bc46_RotationChanged
func callbacksettingsController15bc46_RotationChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rotationChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_ScaleChanged
func callbacksettingsController15bc46_ScaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scaleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_SmoothChanged
func callbacksettingsController15bc46_SmoothChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "smoothChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbacksettingsController15bc46_StateChanged
func callbacksettingsController15bc46_StateChanged(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	}

}

//export callbacksettingsController15bc46_TextureProvider
func callbacksettingsController15bc46_TextureProvider(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "textureProvider"); signal != nil {
		return std_quick.PointerFromQSGTextureProvider((*(*func() *std_quick.QSGTextureProvider)(signal))())
	}

	return std_quick.PointerFromQSGTextureProvider(NewSettingsControllerFromPointer(ptr).TextureProviderDefault())
}

func (ptr *settingsController) TextureProviderDefault() *std_quick.QSGTextureProvider {
	if ptr.Pointer() != nil {
		tmpValue := std_quick.NewQSGTextureProviderFromPointer(C.settingsController15bc46_TextureProviderDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbacksettingsController15bc46_TouchEvent
func callbacksettingsController15bc46_TouchEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchEvent"); signal != nil {
		(*(*func(*std_gui.QTouchEvent))(signal))(std_gui.NewQTouchEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).TouchEventDefault(std_gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *settingsController) TouchEventDefault(event std_gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_TouchEventDefault(ptr.Pointer(), std_gui.PointerFromQTouchEvent(event))
	}
}

//export callbacksettingsController15bc46_TouchUngrabEvent
func callbacksettingsController15bc46_TouchUngrabEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchUngrabEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *settingsController) TouchUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_TouchUngrabEventDefault(ptr.Pointer())
	}
}

//export callbacksettingsController15bc46_TransformOriginChanged
func callbacksettingsController15bc46_TransformOriginChanged(ptr unsafe.Pointer, vqq C.longlong) {
	if signal := qt.GetSignal(ptr, "transformOriginChanged"); signal != nil {
		(*(*func(std_quick.QQuickItem__TransformOrigin))(signal))(std_quick.QQuickItem__TransformOrigin(vqq))
	}

}

//export callbacksettingsController15bc46_Update
func callbacksettingsController15bc46_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *settingsController) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_UpdateDefault(ptr.Pointer())
	}
}

//export callbacksettingsController15bc46_UpdatePolish
func callbacksettingsController15bc46_UpdatePolish(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updatePolish"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *settingsController) UpdatePolishDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_UpdatePolishDefault(ptr.Pointer())
	}
}

//export callbacksettingsController15bc46_VisibleChanged
func callbacksettingsController15bc46_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_WheelEvent
func callbacksettingsController15bc46_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *settingsController) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbacksettingsController15bc46_WidthChanged
func callbacksettingsController15bc46_WidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "widthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_WindowChanged
func callbacksettingsController15bc46_WindowChanged(ptr unsafe.Pointer, window unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowChanged"); signal != nil {
		(*(*func(*std_quick.QQuickWindow))(signal))(std_quick.NewQQuickWindowFromPointer(window))
	}

}

//export callbacksettingsController15bc46_XChanged
func callbacksettingsController15bc46_XChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_YChanged
func callbacksettingsController15bc46_YChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_ZChanged
func callbacksettingsController15bc46_ZChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "zChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacksettingsController15bc46_ChildEvent
func callbacksettingsController15bc46_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *settingsController) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbacksettingsController15bc46_ConnectNotify
func callbacksettingsController15bc46_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSettingsControllerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *settingsController) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbacksettingsController15bc46_CustomEvent
func callbacksettingsController15bc46_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *settingsController) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbacksettingsController15bc46_DeleteLater
func callbacksettingsController15bc46_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSettingsControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *settingsController) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbacksettingsController15bc46_Destroyed
func callbacksettingsController15bc46_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbacksettingsController15bc46_DisconnectNotify
func callbacksettingsController15bc46_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSettingsControllerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *settingsController) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbacksettingsController15bc46_EventFilter
func callbacksettingsController15bc46_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSettingsControllerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *settingsController) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.settingsController15bc46_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbacksettingsController15bc46_ObjectNameChanged
func callbacksettingsController15bc46_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbacksettingsController15bc46_TimerEvent
func callbacksettingsController15bc46_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewSettingsControllerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *settingsController) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.settingsController15bc46_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["settings.settingsController_ITF"] = settingsController{}
	qt.FuncMap["settings.NewSettingsController"] = NewSettingsController
}
