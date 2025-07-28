package ui

import (
	"log"

	"github.com/Thales-Coutinho/AutoSync-GUI/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	components := CreateComponents(cfg)
	SetupNumberValidation(components.NumberOfBackups)

	handlers := NewHandlers(w, components)
	handlers.SetupSelectDirHandler()
	handlers.SetupBackupHandler()
	handlers.SetupCloseHandler()

	content := container.NewVBox(
		components.Label,
		components.ErrorLabel,
		components.RemoteNameRow,
		components.DirToBackupRow,
		components.NumberOfBackupsRow,
		components.DateFormatRow,
		components.BackupButton,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
