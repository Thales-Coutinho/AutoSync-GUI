package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("AutoSync")
	icon, _ := fyne.LoadResourceFromPath("assets/icon.png")
	w.SetIcon(icon)

	label := widget.NewLabel("AutoSync - a program to backup your files automatically and send them to the cloud.")

	errorLabel := widget.NewLabel("teste")

	dirLabel := widget.NewLabel("")

	selectButton := widget.NewButton("Select Directory", func() {
		dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				dirLabel.SetText(uri.Path())
			}
		}, w).Show()
	})
	numberOfBackups := widget.NewEntry()

	button := widget.NewButton("Backup", func() {
		label.SetText("Status: Backup in progress...")
	})

	content := container.NewVBox(
		label,
		errorLabel,
		dirLabel,
		selectButton,
		numberOfBackups,
		button,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
