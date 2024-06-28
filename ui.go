package main

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"milvusmetagui/show"
	"strings"
	"time"
)

type ui struct {
	labelKey    *widget.Label
	connButton  *widget.Button
	parseButton *widget.Button
	ipEntry     *widget.Entry
	portEntry   *widget.Entry
	userEntry   *widget.Entry
	passEntry   *widget.Entry
	keyEntry    *widget.Entry
	textArea    *widget.Entry
	etcdClient  *clientv3.Client
	window      fyne.Window
}

func newMilvusmetar() *ui {
	return &ui{}
}

func (m *ui) connect() {
	endpoints := m.ipEntry.Text + ":" + m.portEntry.Text
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoints},
		Username:    m.userEntry.Text,
		Password:    m.passEntry.Text,
		DialTimeout: 2 * time.Second,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(),
		},
	})
	if err != nil {
		dialog.NewError(err, m.window).Show()
		return
	} else {
		title := "info"
		msg := "Connected successfully"
		dialog.NewInformation(title, msg, m.window).Show()
		m.parseButton.Enable()
	}
	m.etcdClient = c
}

func (m *ui) parse() {
	m.textArea.SetText("")
	keyPrefix := m.keyEntry.Text
	if keyPrefix == "" {
		m.textArea.SetText("please input etcd key")
		return
	}
	opts := []clientv3.OpOption{
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
		clientv3.WithLimit(5000),
		clientv3.WithRange(clientv3.GetPrefixRangeEnd(keyPrefix)),
	}
	ctx := context.Background()
	resp, _ := m.etcdClient.Get(ctx, keyPrefix, opts...)
	var result string
	//keyPrefix := "by-dev/meta/root-coord/database/collection-info/1"
	if strings.Contains(keyPrefix, "database/collection-info") {
		result = show.ShowCollsInfo(resp)
	}
	//keyPrefix := "by-dev/meta/root-coord/database/db-info"
	if strings.Contains(keyPrefix, "database/db-info") {
		result = show.ShowDbsInfo(resp)
	}
	//keyPrefix := "by-dev/meta/root-coord/fields/449657611215201437"
	if strings.Contains(keyPrefix, "root-coord/fields") {
		result = show.ShowFieldsInfo(resp)
	}
	//keyPrefix := "by-dev/meta/field-index/449657611215201437/449663566298480829"
	if strings.Contains(keyPrefix, "field-index") {
		result = show.ShowIndexesInfo(resp)
	}
	//keyPrefix := "by-dev/meta/root-coord/partitions/449657611215201437"
	if strings.Contains(keyPrefix, "root-coord/partitions") {
		result = show.ShowPartsInfo(resp)
	}
	//keyPrefix := "by-dev/meta/segment-index/449682565894505321/449682565894505322/449682565894705334/449682565894906272"
	if strings.Contains(keyPrefix, "segment-index") {
		result = show.ShowSegIndexesInfo(resp)
	}
	if result == "" {
		m.textArea.SetText("etcd key is invalid")
	} else {
		m.textArea.SetText(result)
	}

}

func (m *ui) loadUI(app fyne.App) {

	m.ipEntry = widget.NewEntry()
	m.ipEntry.SetPlaceHolder("input ip address")
	ipEntryL := widget.NewFormItem("ETCD IP", m.ipEntry)

	m.portEntry = widget.NewEntry()
	m.portEntry.SetPlaceHolder("input port")
	m.portEntry.SetText("2379")
	portEntryL := widget.NewFormItem("PORT", m.portEntry)

	m.userEntry = widget.NewEntry()
	m.userEntry.SetPlaceHolder("input username")
	userEntryL := widget.NewFormItem("USERNAME", m.userEntry)

	m.passEntry = widget.NewPasswordEntry()
	m.passEntry.SetPlaceHolder("input password")
	passEntryL := widget.NewFormItem("PASSWORD", m.passEntry)
	form1 := widget.NewForm(
		ipEntryL, userEntryL,
	)
	form2 := widget.NewForm(
		portEntryL, passEntryL,
	)
	c1 := container.NewGridWithColumns(2, form1, form2)
	m.connButton = widget.NewButton("connect", m.connect)
	m.connButton.Importance = widget.HighImportance
	c2 := container.NewGridWithColumns(4, m.connButton)

	m.labelKey = widget.NewLabel("input etcd key:")

	m.keyEntry = widget.NewEntry()
	m.keyEntry.SetPlaceHolder("input etcd key")

	m.parseButton = widget.NewButton("parse", m.parse)
	m.parseButton.Importance = widget.HighImportance
	m.parseButton.Disable()

	cbtn := container.NewGridWithColumns(4, m.parseButton)

	m.textArea = widget.NewMultiLineEntry()

	c3 := container.NewVBox(c1, c2, m.labelKey, m.keyEntry, cbtn)

	content := container.NewBorder(c3, nil, nil, nil, m.textArea)

	m.window = app.NewWindow("milvus meta parse")

	m.window.SetContent(content)
	m.window.Resize(fyne.NewSize(600, 600))
	m.window.CenterOnScreen()
	m.window.Show()
}
