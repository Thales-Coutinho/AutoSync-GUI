package main

import (
	"strconv"

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

	dirLabel := widget.NewLabel("Selected Directory: None")

	selectDirButton := widget.NewButton("Select Directory", func() {
		dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				dirLabel.SetText(uri.Path())
			}
		}, w).Show()
	})
	dirToBackupRow := container.NewHBox(
		selectDirButton,
		dirLabel,
	)

	numberOfBackupsLabel := widget.NewLabel("Number of Backups:")
	numberOfBackups := widget.NewEntry()
	numberOfBackupsRow := container.NewHBox(
		numberOfBackupsLabel,
		numberOfBackups,
	)

	numberOfBackups.Validator = func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		return nil
	}
	numberOfBackups.OnChanged = func(s string) {

		if s == "" {
			return
		}

		lastChar := s[len(s)-1:]
		if _, err := strconv.Atoi(lastChar); err != nil {
			numberOfBackups.SetText(s[:len(s)-1])
		}
	}

	button := widget.NewButton("Backup", func() {
		label.SetText("Status: Backup in progress...")
	})

	content := container.NewVBox(
		label,
		errorLabel,
		dirToBackupRow,
		numberOfBackupsRow,
		button,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
