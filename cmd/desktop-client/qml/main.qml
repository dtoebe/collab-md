import QtQuick 2.11
import QtQuick.Controls 1.4
import QtQuick.Layouts 1.3

ApplicationWindow {
    width: 640
    height: 480
    color: "#f1f1f1"
    visible: true

	toolBar: ToolBar {
		width: parent.width

		RowLayout {
			width: parent.width
			height: parent.height

			Button {
				Layout.alignment: Qt.AlignLeft
				text: "Copy Markdown"

				onClicked: {
					bridge.mdCopyText = mdarea.text
				}
			}

			Button {
				Layout.alignment: Qt.AlignCenter
				text: "Connect to server"

				onClicked: {
					bridge.connClick = ""
					bridge.connClick = "clicked"
				}
			}

			Button {
				Layout.alignment: Qt.AlignRight
				text: "Copy HTML"

				onClicked: {
					bridge.rtCopyText = mdarea.text
				}
			}
		}
	}

    RowLayout {
        width: parent.width
        height: parent.height

        TextArea {
            id: mdarea
            Layout.alignment: Qt.AlignCenter
            Layout.preferredWidth: (parent.width / 2) - 5;
            Layout.preferredHeight: parent.height - 10
            text: bridge.mdareaText
			Keys.onReleased: {
				bridge.mdareaText = mdarea.text
			}
        }

        TextArea {
            id: rtarea
            Layout.alignment: Qt.AlignCenter
            Layout.preferredWidth: (parent.width / 2) - 5
            Layout.preferredHeight: parent.height - 10
            textFormat: TextEdit.RichText
            text: bridge.rtareaText
        }
	}
}
