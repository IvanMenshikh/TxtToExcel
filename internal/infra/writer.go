package infra

import (
	"github.com/xuri/excelize/v2"
)

type ExcelWriter struct {
	filePath string
}

func NewExcelWriter(filePath string) *ExcelWriter {
	return &ExcelWriter{filePath: filePath}
}

// Write записывает данные в Excel файл
func (w *ExcelWriter) Write(data [][]string) error {
	f := excelize.NewFile()
	sheet := f.GetSheetName(f.GetActiveSheetIndex())

	// Заголовки
	headers := []string{
		"Тип документа",
		"?",
		"Заголовок",
		"Статус",
		"?",
		"Тип поручения",
		"Организация",
		"Номер поручения",
		"Текст поручения",
		"Автор",
		"Ответственный исполнитель",
		"Должность",
		"?",
		"?",
		"?",
		"?",
		"?",
		"?",
		"Статус поручения",
		"?",
		"Отчёт",
		"?",
		"?",
		"?",
	}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Данные
	for row, record := range data {
		for col, val := range record {
			cell, _ := excelize.CoordinatesToCellName(col+1, row+2)
			f.SetCellValue(sheet, cell, val)
		}
	}

	return f.SaveAs(w.filePath)
}
