package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/Thales-Coutinho/AutoSync-GUI/internal/config"
)

type Handlers struct {
	Window     fyne.Window
	Components *AppComponents
}

func NewHandlers(w fyne.Window, components *AppComponents) *Handlers {
	return &Handlers{
		Window:     w,
		Components: components,
	}
}

func (h *Handlers) SetupSelectDirHandler() {
	h.Components.SelectDirButton.OnTapped = func() {
		dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				h.Components.DirLabel.SetText(uri.Path())
			}
		}, h.Window).Show()
	}
}

func (h *Handlers) SetupBackupHandler() {
	h.Components.BackupButton.OnTapped = func() {
		h.Components.Label.SetText("Status: Backup in progress...")
	}
}

func (h *Handlers) SetupCloseHandler() {
	h.Window.SetCloseIntercept(func() {
		if err := config.Save(config.Config{
			NumberOfBackups: config.ParseBackups(h.Components.NumberOfBackups.Text),
			BackupDir:       h.Components.DirLabel.Text,
		}); err != nil {
			log.Printf("Error saving config: %v", err)
		}
		h.Window.Close()
	})
}
