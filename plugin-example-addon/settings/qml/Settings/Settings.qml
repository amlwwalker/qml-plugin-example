import QtQuick 2.7          //Rectangle
import QtQuick.Controls 2.1 //StackView
import QtQuick.Layouts 1.3  //GridLayout

import Settings 1.0
import Theme 1.0
import "." as T
SettingsController {
    //   property int margin: 5
  // Rectangle {
      // width: 600
      // height: 400
      // width: mainLayout.implicitWidth + 2 * margin
      // height: mainLayout.implicitHeight + 2 * margin
      // minimumWidth: mainLayout.Layout.minimumWidth + 2 * margin
      // minimumHeight: mainLayout.Layout.minimumHeight + 2 * margin
      // width: parent.width
      // height: parent.height
      Layout.fillWidth: true
      ColumnLayout {
          id: mainLayout
          // anchors.fill: parent
          anchors.right: parent.right
          anchors.left: parent.left
        //   anchors.margins: margin
          GroupBox {
              id: rowBox
              Layout.fillWidth: true

              GridLayout {
                  id: versioningLayout
                  rows: 4
                  columns: 1
                  flow: GridLayout.TopToBottom
                  anchors.fill: parent
                  Label { 
                    text: "Set the file naming convention"; 
                    font.pixelSize: 18
                    color: Theme.mainColor
                    font.family: "Helvetica"
                    Layout.fillWidth: true 
                  }
                Row { //filename should be a given
                    spacing: 5
                    Label { text: "%{filename}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    // Label { text: "%{hash}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    // Label { text: "%{timeUnix}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    // Label { text: "%{client}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    // Label { text: "%{job}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    // Label { text: "%{owner}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                  }
                  Row {
                    spacing: 5
                    Label { text: "%{bigVersion}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    Label { text: "%{littleVersion}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    Label { text: "%{microVersion}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    Label { text: "%{time}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    Label { text: "%{creator}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                    Label { text: "%{description}"; color: Theme.mainColor; MouseArea {anchors.fill: parent; cursorShape: Qt.PointingHandCursor; onClicked: versionConvention.text += "."+parent.text}}
                  }
              Row {
                      Layout.fillWidth: true

                  TextField {
                      id: versionConvention
                      placeholderText: "%{filename}.%{bigVersion}.%{littleVersion}.%{microVersion} -> [%{time}] [%{client}] [%{job}] [%{creator}] [%{owner}] [%{hash}] %{message}"
                      text: "%{filename}.%{bigVersion}.%{littleVersion}.%{microVersion}"
                      width: 500
                      Layout.fillWidth: true
                  }
                  // T.Button {
                  //     text: "Update Convention"
                  //     Layout.fillWidth: true
                  //     anchors.right: versioningLayout.right
                  // }
              }
              }
          }

          GroupBox {
              id: gridBox
              title: "Application Settings"
              Layout.fillWidth: true
              GridLayout {
                  id: gridLayout
                  rows: 4
                  // columns: 3
                  flow: GridLayout.TopToBottom
                  anchors.fill: parent

                  Label { text: "Happy to send back statistics about how the app is performing"; Layout.fillWidth: true  }
                  Label { text: "Save a screen shot of the file when creating versions"; Layout.fillWidth: true  }
                  Label { text: "Notify when a new file is ready"; Layout.fillWidth: true  }
                  Label { text: "Notify when a patch is created"; Layout.fillWidth: true  }

                  CheckBox { id: statisticsCheck; checked: true}
                  CheckBox { id: screenshotCheck; checked: true}
                  CheckBox { id: newFileReadySystemNotifyCheck; checked: true}
                  CheckBox { id: patchSystemNotifyCheck; checked: true}
                  // CheckBox { id: patchSystemNotifyCheck; checked: true}
                  // Label { text: ""; Layout.fillWidth: true }
                  // Label { text: ""; Layout.fillWidth: true}
                  // Label { text: ""; Layout.fillWidth: true}
                  
              }
          }
          // TextArea {
          //     id: t3
          //     text: "This fills the whole cell"
          //     Layout.minimumHeight: 30
          //     Layout.fillHeight: true
          //     Layout.fillWidth: true
          // }
          // GroupBox {
          //     id: stackBox
          //     title: "Stack layout"
          //     implicitWidth: 200
          //     implicitHeight: 60
          //     Layout.fillWidth: true
          //     Layout.fillHeight: true
          //     StackLayout {
          //         id: stackLayout
          //         anchors.fill: parent

          //         function advance() { currentIndex = (currentIndex + 1) % count }

          //         Repeater {
          //             id: stackRepeater
          //             model: 5
          //             Rectangle {
          //                 color: Qt.hsla((0.5 + index)/stackRepeater.count, 0.3, 0.7, 1)
          //                 Button { anchors.centerIn: parent; text: "Page " + (index + 1); onClicked: { stackLayout.advance() } }
          //             }
          //         }
          //     }
          // }
      }
  // }
}