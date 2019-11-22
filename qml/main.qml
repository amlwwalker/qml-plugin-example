import QtQuick.Controls 2.4 //Menu
import QtQuick.Dialogs 1.2  //MessageDialog
import QtQuick.Layouts 1.3  //GridLayout
import QtQuick 2.2

import "." as Custom               //needed for name clash with std Controls

// import Settings 1.0

ApplicationWindow {
  id: app
  visible: true
  title: "plugins"
  minimumWidth: 800; minimumHeight: 768
  Rectangle {
    anchors.fill: parent
    id: stackLayout
    // objectName: "stackLayout"
  }
}
