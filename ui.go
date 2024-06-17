package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ui struct {
	labelKey    *widget.Label
	connButton  *widget.Button
	parseButton *widget.Button
	ipLabel     *widget.Label
	ipEntry     *widget.Entry
	portLabel   *widget.Label
	portEntry   *widget.Entry
	userLable   *widget.Label
	userEntry   *widget.Entry
	passLable   *widget.Label
	passEntry   *widget.Entry
	keyEntry    *widget.Entry
	textArea    *widget.Entry
	window      fyne.Window
}

func newMilvusmetar() *ui {
	return &ui{}
}

func (m *ui) loadUI(app fyne.App) {

	m.ipLabel = widget.NewLabel("ETCD IP")
	m.ipLabel.Resize(fyne.NewSize(30, 30))
	m.ipLabel.Alignment = fyne.TextAlignLeading
	m.ipLabel.Move(fyne.NewPos(10, 10))

	m.ipEntry = widget.NewEntry()
	m.ipEntry.SetPlaceHolder("input ip address")
	m.ipEntry.Resize(fyne.NewSize(200, 40))
	m.ipEntry.Move(fyne.NewPos(10, 40))

	m.portLabel = widget.NewLabel("PORT")
	m.portLabel.Resize(fyne.NewSize(30, 30))
	m.portLabel.Alignment = fyne.TextAlignLeading
	m.portLabel.Move(fyne.NewPos(240, 10))

	m.portEntry = widget.NewEntry()
	m.portEntry.SetPlaceHolder("input port")
	m.portEntry.Resize(fyne.NewSize(200, 40))
	m.portEntry.Move(fyne.NewPos(240, 40))

	m.userLable = widget.NewLabel("USER")
	m.userLable.Resize(fyne.NewSize(30, 30))
	m.userLable.Alignment = fyne.TextAlignLeading
	m.userLable.Move(fyne.NewPos(10, 80))

	m.userEntry = widget.NewEntry()
	m.userEntry.SetPlaceHolder("input username")
	m.userEntry.Resize(fyne.NewSize(200, 40))
	m.userEntry.Move(fyne.NewPos(10, 110))

	m.passLable = widget.NewLabel("PASSWORD")
	m.passLable.Resize(fyne.NewSize(30, 30))
	m.passLable.Alignment = fyne.TextAlignLeading
	m.passLable.Move(fyne.NewPos(240, 80))

	m.passEntry = widget.NewPasswordEntry()
	m.passEntry.SetPlaceHolder("input password")
	m.passEntry.Resize(fyne.NewSize(200, 40))
	m.passEntry.Move(fyne.NewPos(240, 110))

	m.connButton = widget.NewButton("connect", func() {})
	m.connButton.Importance = widget.HighImportance
	m.connButton.Resize(fyne.NewSize(100, 40))
	m.connButton.Move(fyne.NewPos(10, 160))

	m.labelKey = widget.NewLabel("input etcd key:")
	m.labelKey.Resize(fyne.NewSize(60, 30))
	m.labelKey.Alignment = fyne.TextAlignLeading
	m.labelKey.Move(fyne.NewPos(10, 210))

	m.keyEntry = widget.NewEntry()
	m.keyEntry.Resize(fyne.NewSize(450, 40))
	m.keyEntry.SetPlaceHolder("input etcd key")
	m.keyEntry.Move(fyne.NewPos(10, 250))

	m.parseButton = widget.NewButton("parse", func() {})
	m.parseButton.Importance = widget.HighImportance
	m.parseButton.Resize(fyne.NewSize(100, 40))
	m.parseButton.Move(fyne.NewPos(10, 300))

	m.textArea = widget.NewMultiLineEntry()
	m.textArea.Resize(fyne.NewSize(450, 240))
	m.textArea.Move(fyne.NewPos(10, 350))

	m.window = app.NewWindow("milvus meta parse")
	c := container.NewWithoutLayout(m.ipLabel,
		m.ipEntry,
		m.portLabel,
		m.portEntry,
		m.userLable,
		m.userEntry,
		m.passLable,
		m.passEntry,
		m.connButton,
		m.labelKey,
		m.keyEntry,
		m.parseButton,
		m.textArea)

	m.window.SetContent(c)
	m.window.Resize(fyne.NewSize(500, 600))
	m.window.CenterOnScreen()
	m.window.Show()
}
