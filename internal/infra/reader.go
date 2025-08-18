package infra

import (
	"bufio"
	"os"
	"strings"
)

type TxtReader struct {
	filePath  string
	delimiter string
}

func NewTxtReader(filePath, delimiter string) *TxtReader {
	return &TxtReader{filePath: filePath, delimiter: delimiter}
}

// Read возвращает данные из текстового файла
func (r *TxtReader) Read() ([][]string, int, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var records [][]string
	emptyLines := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			// пустую строку пропускаем и считаем проблемной
			emptyLines++
			continue
		}

		parts := strings.Split(line, r.delimiter)
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		records = append(records, parts)
	}

	if err := scanner.Err(); err != nil {
		return nil, emptyLines, err
	}
	return records, emptyLines, nil
}
