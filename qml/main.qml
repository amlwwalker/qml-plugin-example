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
    TabBar {
      id: bar
      objectName: "tabBar"
      Layout.fillWidth: true
      Layout.fillHeight: true      
      width: parent.width
      background: Rectangle {
          color: "transparent"
      }
      TabButton {
          text: qsTr("Calendar")
      }
      TabButton {
          text: qsTr("Listing")
      }
    }
    SwipeView {
      id: stackLayout
      objectName: "stackLayout"
      width: parent.width
      implicitHeight: parent.height - bar.height
      anchors.top: bar.bottom
      
      currentIndex: bar.currentIndex
      Item {
          id: calendarTab
          Rectangle { 
            id: calendar 
            width: parent.width
            height: parent.height
            color: "red"
            Text {
              text: "calendar"
            }
          }
      }
      Item {
        id: listingTab
          Rectangle {
            id: listing
            width: parent.width
            height: parent.height
            color: "green"
            Text {
              text: "listing"
            }
          }
      }
      function addToStackLayout(item){stackLayout.addItem(item);}
    }
  }
}
