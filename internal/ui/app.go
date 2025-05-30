package ui

import (
	"log"
	"strconv"

	"github.com/Thales-Coutinho/AutoSync-GUI/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func RunApplication() {
	a := app.New()
	w := a.NewWindow("AutoSync")
	icon, _ := fyne.LoadResourceFromPath("assets/icon.png")
	w.SetIcon(icon)

	if err := config.InitConfigDir(); err != nil {
		log.Fatal("Error creating config dir:", err)
	}
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Warning loading config: %v", err)
		cfg = config.Config{NumberOfBackups: 5}
	}

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
	numberOfBackups.SetText(strconv.Itoa(cfg.NumberOfBackups))
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

	w.SetCloseIntercept(func() {
		if err := config.Save(config.Config{
			NumberOfBackups: config.ParseBackups(numberOfBackups.Text),
			BackupDir:       dirLabel.Text,
		}); err != nil {
			log.Printf("Error saving config: %v", err)
		}
		w.Close()
	})

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
