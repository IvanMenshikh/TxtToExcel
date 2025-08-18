package app

import "fmt"

type Reader interface {
	Read() ([][]string, int, error)
}

type Writer interface {
	Write([][]string) error
}

type UI interface {
	ShowInfo(title, msg string)
	ShowError(title, msg string)
}

type Converter struct {
	reader Reader
	writer Writer
	ui     UI
}

func NewConverter(r Reader, w Writer, u UI) *Converter {
	return &Converter{
		reader: r,
		writer: w,
		ui:     u,
	}
}

func (c *Converter) Run() error {
	data, emptyLines, err := c.reader.Read()
	if err != nil {
		return fmt.Errorf("ошибка чтения: %w", err)
	}

	if err := c.writer.Write(data); err != nil {
		return fmt.Errorf("ошибка записи: %w", err)
	}

	// показываем инфо о пустых строках
	if emptyLines > 0 {
		c.ui.ShowInfo("Импорт завершён",
			fmt.Sprintf("Файл успешно создан.\nПропущено пустых строк: %d\nПодробности — error.log", emptyLines))
	} else {
		c.ui.ShowInfo("Импорт завершён", "Файл успешно создан. Пустых строк нет.")
	}

	return nil
}

