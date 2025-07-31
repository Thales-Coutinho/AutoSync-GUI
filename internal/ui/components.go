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
	RemoteNameLabel    *widget.Label
	RemoteNameEntry    *widget.Entry
	ReotePathLabel     *widget.Label
	RemotePathEntry    *widget.Entry
	RemoteNameRow      *fyne.Container
	SelectDirButton    *widget.Button
	DirLabel           *widget.Label
	NumberOfBackups    *widget.Entry
	BackupButton       *widget.Button
	DirToBackupRow     *fyne.Container
	NumberOfBackupsRow *fyne.Container
	CheckUseDateonName *widget.Check
	DateFormatRow      *fyne.Container
	FileNameLabel      *widget.Label
	FileNameEntry      *widget.Entry
	FileNameRow        *fyne.Container
}

func CreateComponents(cfg config.Config) *AppComponents {
	label := widget.NewLabel("AutoSync - a program to backup your files automatically and send them to the cloud.")
	errorLabel := widget.NewLabel("teste")
	dirLabel := widget.NewLabel(cfg.BackupDir)

	RemoteNameLabel := widget.NewLabel("Remote Name:")
	RemoteNameEntry := container.NewGridWrap(fyne.NewSize(100, 40), widget.NewEntry())
	RemotePathLabel := widget.NewLabel("Remote Path:")
	entry := widget.NewEntry()
	entry.SetText(cfg.RemotePath)
	RemotePathEntry := container.NewGridWrap(fyne.NewSize(100, 40), entry)

	RemoteNameRow := container.NewHBox(
		RemoteNameLabel,
		RemoteNameEntry,
		RemotePathLabel,
		RemotePathEntry,
	)

	selectDirButton := widget.NewButton("Select Directory", nil)
	dirToBackupRow := container.NewHBox(selectDirButton, dirLabel)

	numberOfBackupsLabel := widget.NewLabel("Number of Backups:")
	numberOfBackups := widget.NewEntry()
	numberOfBackups.SetText(strconv.Itoa(cfg.NumberOfBackups))

	numberOfBackupsRow := container.NewHBox(
		numberOfBackupsLabel,
		numberOfBackups,
	)

	dateFormats := []string{
		"dd/MM/yyyy",
		"MM/dd/yyyy",
		"yyyy-MM-dd",
		"dd-MM-yyyy",
	}
	dateFormatLabel := widget.NewLabel("Date Format: ")
	dateFormatSelect := widget.NewSelect(dateFormats, func(selected string) {
	})
	dateFormatSelect.SetSelectedIndex(0)

	CheckUseDateonName := widget.NewCheck("Use Date in File Name", func(checked bool) {
		ChangedateFormatSelect(checked, dateFormatSelect)
	})
	dateFormatRow := container.NewHBox(
		CheckUseDateonName,
		dateFormatLabel,
		dateFormatSelect,
	)
	FileNameLabel := widget.NewLabel("File Name:")
	FileNameEntry := container.NewGridWrap(fyne.NewSize(100, 40), widget.NewEntry())
	FileNameRow := container.NewHBox(
		FileNameLabel,
		FileNameEntry,
	)

	button := widget.NewButton("Backup", nil) // callback será definido depois

	return &AppComponents{
		Label:              label,
		ErrorLabel:         errorLabel,
		RemoteNameRow:      RemoteNameRow,
		DirLabel:           dirLabel,
		SelectDirButton:    selectDirButton,
		NumberOfBackups:    numberOfBackups,
		BackupButton:       button,
		DirToBackupRow:     dirToBackupRow,
		NumberOfBackupsRow: numberOfBackupsRow,
		CheckUseDateonName: CheckUseDateonName,
		DateFormatRow:      dateFormatRow,
		FileNameRow:        FileNameRow,
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

func ChangedateFormatSelect(checked bool, dateFormatSelect *widget.Select) {
	if checked {
		dateFormatSelect.Enable()
	} else {
		dateFormatSelect.Disable()
	}
}
