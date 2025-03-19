package excel

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

const (
	DEFAULT_SINGLE_SHEETNAME = "Sheet1"
	SHEET_PATTERN            = "sheet%d"
)

// ExcelModel Excel返回结果模型
type ExcelModel struct {
	Sheets []*ExcelSheet
}

// ExcelSheet 一个sheet的模型
type ExcelSheet struct {
	Heads     []string
	Content   [][]interface{}
	SheetName string
}

// GetSingleSheetExcel 这个地方获取只有一个sheet的excel对象
func GetSingleSheetExcel(head []string, content [][]interface{}, sheetName ...string) *ExcelModel {
	excelModel := new(ExcelModel)
	sheet := new(ExcelSheet)

	if len(sheetName) > 0 {
		sheet.SheetName = sheetName[0]
	} else {
		sheet.SheetName = DEFAULT_SINGLE_SHEETNAME
	}
	sheet.Heads = head
	sheet.Content = content
	excelModel.Sheets = []*ExcelSheet{sheet}
	return excelModel
}

// GetEmptyExcel ...
func GetEmptyExcel() *ExcelModel {
	return new(ExcelModel)
}

// OpenReader ...
func OpenReader(r io.Reader) (*ExcelModel, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, nil
	}
	sheetsData := make([]*ExcelSheet, 0, len(f.GetSheetList()))
	for _, sheet := range f.GetSheetList() {
		// 获取 Sheet 上所有单元格
		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, err
		}
		head := make([]string, 0)
		content := make([][]interface{}, 0)
		for i, row := range rows {
			contentItem := make([]interface{}, 0)
			for _, colCell := range row {
				if i == 0 {
					head = append(head, colCell)
					continue
				}
				contentItem = append(contentItem, colCell)
			}
			if len(contentItem) == 0 {
				continue
			}
			content = append(content, contentItem)
		}
		sheetsData = append(sheetsData, &ExcelSheet{
			SheetName: sheet,
			Heads:     head,
			Content:   content,
		})
	}
	return &ExcelModel{sheetsData}, nil
}

// AddSheet 这个地方添加sheet
func (e *ExcelModel) AddSheet(head []string, content [][]interface{}, sheetName ...string) *ExcelModel {
	sheet := new(ExcelSheet)
	if len(sheetName) > 0 {
		sheet.SheetName = sheetName[0]
	} else {
		sheet.SheetName = fmt.Sprintf(SHEET_PATTERN, len(e.Sheets))
	}
	sheet.Heads = head
	sheet.Content = content
	e.Sheets = append(e.Sheets, sheet)
	return e
}

// BuildExcel 进行报表的对象的转换
func (e *ExcelModel) BuildExcel() *excelize.File {
	file := excelize.NewFile()
	// 需要删除默认的sheet
	isDefaultEx := false
	for _, sheet := range e.Sheets {
		if sheet.SheetName == DEFAULT_SINGLE_SHEETNAME {
			isDefaultEx = true
		}
		file.NewSheet(sheet.SheetName)
		// 设置表头,slice的下标是从0开始，但是excel的下标是从1开始
		for col, head := range sheet.Heads {
			headAxis, _ := excelize.CoordinatesToCellName(col+1, 1)
			file.SetCellValue(sheet.SheetName, headAxis, head)
		}
		// 设置表的内容
		for row, colData := range sheet.Content {
			for idx, value := range colData {
				celAxis, _ := excelize.CoordinatesToCellName(idx+1, row+2)
				file.SetCellValue(sheet.SheetName, celAxis, value)
			}
		}
	}
	// 删除默认的sheet
	if !isDefaultEx {
		file.DeleteSheet(DEFAULT_SINGLE_SHEETNAME)
	}
	return file
}

// SaveToFile 保存excel到指定路径
func (e *ExcelModel) SaveToFile(filePath string) error {
	return e.BuildExcel().SaveAs(filePath)
}

// ToBuf 生成byte数组
func (e *ExcelModel) ToBuf() (*bytes.Buffer, error) {
	return e.BuildExcel().WriteToBuffer()
}

// Write provides a function to write to an io.Writer.
func (e *ExcelModel) Write(w io.Writer) error {
	_, err := e.BuildExcel().WriteTo(w)
	return err
}
