package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"txt-to-excel/internal/app"
	"txt-to-excel/internal/infra"
	"txt-to-excel/internal/ui"
)

func main() {
	// путь к exe
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)

	inputFile := filepath.Join(exeDir, "filtered_logs.txt")
	outputFile := filepath.Join(exeDir, "result.xlsx")

	reader := infra.NewTxtReader(inputFile, "|")
	writer := infra.NewExcelWriter(outputFile)
	uiHandler := ui.NewMessageBox()

	converter := app.NewConverter(reader, writer, uiHandler)

	if err := converter.Run(); err != nil {
		// пишем в лог
		logFile, _ := os.OpenFile(filepath.Join(exeDir, "error.log"),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()
		logger := log.New(logFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
		logger.Println(err)

		// показываем пользователю сообщение
		if errors.Is(err, os.ErrNotExist) {
			uiHandler.ShowError("Ошибка", "Файл data.txt не найден!")
		} else {
			uiHandler.ShowError("Ошибка", "Что-то пошло не так. Подробности см. в error.log")
		}
		return
	}

	//uiHandler.ShowInfo("Готово", "Файл result.xlsx успешно создан!")
}
