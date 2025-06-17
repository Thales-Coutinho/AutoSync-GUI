package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Thales-Coutinho/AutoSync-GUI/internal/config"
)

type AppComponents struct {
	Label              *widget.Label
	ErrorLabel         *widget.Label
	DirLabel           *widget.Label
	SelectDirButton    *widget.Button
	NumberOfBackups    *widget.Entry
	BackupButton       *widget.Button
	DirToBackupRow     *fyne.Container
	NumberOfBackupsRow *fyne.Container
	DateFormatRow      *fyne.Container
}

func CreateComponents(cfg config.Config) *AppComponents {
	label := widget.NewLabel("AutoSync - a program to backup your files automatically and send them to the cloud.")
	errorLabel := widget.NewLabel("teste")
	dirLabel := widget.NewLabel("Selected Directory: None")

	selectDirButton := widget.NewButton("Select Directory", nil) // callback será definido depois
	dirToBackupRow := container.NewHBox(selectDirButton, dirLabel)

	numberOfBackupsLabel := widget.NewLabel("Number of Backups:")
	numberOfBackups := widget.NewEntry()
	numberOfBackups.SetText(strconv.Itoa(cfg.NumberOfBackups))

	numberOfBackupsRow := container.NewHBox(numberOfBackupsLabel, numberOfBackups)

	dateFormats := []string{
		"dd/MM/yyyy",
		"MM/dd/yyyy",
		"yyyy-MM-dd",
		"dd-MM-yyyy",
	}
	dateFormatLabel := widget.NewLabel("Date Format: ")
	dateFormatSelect := widget.NewSelect(dateFormats, func(selected string) {
		// Callback to handle date format selection
	})
	dateFormatSelect.SetSelectedIndex(0)

	dateFormatRow := container.NewHBox(dateFormatLabel, dateFormatSelect)

	button := widget.NewButton("Backup", nil) // callback será definido depois

	return &AppComponents{
		Label:              label,
		ErrorLabel:         errorLabel,
		DirLabel:           dirLabel,
		SelectDirButton:    selectDirButton,
		NumberOfBackups:    numberOfBackups,
		BackupButton:       button,
		DirToBackupRow:     dirToBackupRow,
		NumberOfBackupsRow: numberOfBackupsRow,
		DateFormatRow:      dateFormatRow,
	}
}

func SetupNumberValidation(entry *widget.Entry) {
	entry.Validator = func(s string) error {
		_, err := strconv.Atoi(s)
		return err
	}

	entry.OnChanged = func(s string) {
		if s == "" {
			return
		}

		lastChar := s[len(s)-1:]
		if _, err := strconv.Atoi(lastChar); err != nil {
			entry.SetText(s[:len(s)-1])
		}
	}
}
