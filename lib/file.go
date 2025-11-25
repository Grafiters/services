package lib

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

var allowedExtensions = map[string]bool{
	".csv":  true,
	".xls":  true,
	".xlsx": true,
}

func ExtractFile(file *multipart.FileHeader) (multipart.File, error) {
	ext := filepath.Ext(file.Filename)

	// Validasi extension
	if !allowedExtensions[ext] {
		return nil, fmt.Errorf("invalid file extension, allowed: .csv, .xls, .xlsx")
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	return src, nil
}

func ParseExcelFile(file multipart.File) ([][]string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	fmt.Println("==========================")
	defer f.Close()

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("no sheet found in file")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read sheet: %w", err)
	}

	fmt.Println(rows)
	fmt.Println("==========================")
	return rows, nil
}

func ParseCSVFile(file multipart.File) ([][]string, error) {
	reader := csv.NewReader(bufio.NewReader(file))
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: %w", err)
	}
	return rows, nil
}
