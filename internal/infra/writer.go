package infra

import (
	"fmt"

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
		"Тип/Вид документа-основания",
		"Номер и дата документа-основания",
		"Заголовок",
		"Статус документа-основания",
		"Резолюция",
		"Контрольное/ Неконтрольное поручение",
		"Организация",
		"Номер и дата поручения",
		"Поручение",
		"Автор поручения",
		"Ответственный исполнитель",
		"Должность отв. исполнителя",
		"Соисполнители",
		"Лица для сведения",
		"Контролеры",
		"Плановый срок исполнения",
		"Фактический срок исполнения",
		"Нарушение срока, дней",
		"Статус",
		"Дата статуса",
		"Отчёт",
		"Доклад",
		"Ссылка на документ-основание",
		"Ссылка на поручение",
	}

	// Стиль для данных
	baseStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	})
	if err != nil {
		return err
	}

	// Стиль для заголовков
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "#000000",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFA500"}, // оранжевый фон
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 6}, // двойная линия
		},
	})
	if err != nil {
		return err
	}

	// Записываем заголовки
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Применяем стиль сразу ко всему ряду заголовков
	lastCol, _ := excelize.ColumnNumberToName(len(headers))
	if err := f.SetCellStyle(sheet, "A1", lastCol+"1", headerStyle); err != nil {
		return err
	}

	// Записываем данные
	for row, record := range data {
		for col, val := range record {
			cell, _ := excelize.CoordinatesToCellName(col+1, row+2)
			f.SetCellValue(sheet, cell, val)
		}
	}

	// Применяем стиль ко всему диапазону данных разом
	if len(data) > 0 {
		lastRow := len(data) + 1
		if err := f.SetCellStyle(sheet, "A2", lastCol+fmt.Sprint(lastRow), baseStyle); err != nil {
			return err
		}
	}

	// Автоширина для читаемости
	for i := 1; i <= len(headers); i++ {
		col, _ := excelize.ColumnNumberToName(i)
		f.SetColWidth(sheet, col, col, 20)
	}

	return f.SaveAs(w.filePath)
}
