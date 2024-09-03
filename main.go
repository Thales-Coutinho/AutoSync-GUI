package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	icon, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		panic(err)
	}
	myApp.SetIcon(icon)

	myWindow := myApp.NewWindow("AutoSync")

	// Crie um widget List para o menu lateral
	menu := widget.NewList(
		func() int { return 3 }, // Número de itens
		func() fyne.CanvasObject { return widget.NewLabel("Menu Item") }, // Template
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// Definir o texto de cada item
			o.(*widget.Label).SetText([]string{"Sobre", "Configurações", "Backup"}[i])
		})

	// Crie containers para as telas
	screenAbout := container.NewVBox(
		widget.NewLabel("Sobre"),
		widget.NewLabel("AutoSync é uma ferramenta para automatizar backups no seu provedor de cloud preferido de forma rápida e facil"),
	)
	screenSettings := container.NewVBox(
		widget.NewLabel("Configurações"),
	)
	screenBackup := container.NewVBox(
		widget.NewLabel("Backup"),
	)

	content := container.NewStack(screenAbout) // Conteúdo inicial

	// Adicione o comportamento de seleção no menu
	menu.OnSelected = func(id widget.ListItemID) {
		switch id {
		case 0:
			content.Objects = []fyne.CanvasObject{screenAbout}
		case 1:
			content.Objects = []fyne.CanvasObject{screenSettings}
		case 2:
			content.Objects = []fyne.CanvasObject{screenBackup}
		}
		content.Refresh()
	}

	// Crie o layout principal
	myWindow.SetContent(container.NewBorder(nil, nil, menu, nil, content))
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
