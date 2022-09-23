package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-ego/gse"
	"github.com/xuri/excelize/v2"
)

type Evaluation struct {
	Date              string // 日期
	PlatformName      string // 平台
	ProductName       string // 产品名称
	EvaluationContent string // 好评/差评
	EvaluationComment string // 评价内容
}

func main() {
	// orderFile := "/Users/sheldon/8月全平台顾客评价汇总表0905.xlsx"
	orderFile := "D:\\全平台顾客评价汇总表.xlsx"

	f, err := excelize.OpenFile(orderFile)
	if err != nil {
		log.Fatalf("Parse Excel Data File Error: %v", err)
	}

	var sheetName string
	index := 0

	for _, name := range f.GetSheetMap() {
		if index > 0 {
			break
		}
		sheetName = name
		index++
	}

	var evaluationList []Evaluation

	rows, _ := f.GetRows(sheetName)
	for rowIndex, row := range rows {
		if rowIndex > 0 {
			evaluation := Evaluation{}

			for colIndex, colCell := range row {
				colCell = strings.TrimSpace(colCell)

				switch colIndex {
				case 0:
					// 日期
					evaluation.Date = colCell
				case 1:
					// 平台
					evaluation.PlatformName = colCell
				case 2:
					// 产品名称
					evaluation.ProductName = colCell
				case 3:
					// 好评/差评
					evaluation.EvaluationContent = colCell
				case 4:
					// 评价内容
					evaluation.EvaluationComment = colCell
				}
			}

			evaluationList = append(evaluationList, evaluation)
		}
	}

	result := make(map[string]map[string]map[string]int)

	var seg gse.Segmenter

	seg.LoadDict()
	new, _ := gse.New("zh,testdata/test_dict3.txt", "alpha")

	for _, evaluation := range evaluationList {
		platformResult, ok := result[evaluation.ProductName]
		if !ok {
			platformResult = make(map[string]map[string]int)
		}

		result[evaluation.ProductName] = platformResult

		platform := fmt.Sprintf("%s-%s", evaluation.PlatformName, evaluation.EvaluationContent)

		contentResult, ok := platformResult[platform]
		if !ok {
			contentResult = make(map[string]int)
		}

		platformResult[platform] = contentResult

		hmm := new.Cut(evaluation.EvaluationComment, true)
		for _, c := range hmm {
			countResult, ok := contentResult[c]
			if !ok {
				contentResult[c] = 1
			} else {
				contentResult[c] = countResult + 1
			}

		}
	}

	for productName, platformResult := range result {
		f := excelize.NewFile()

		for platformName, contentResult := range platformResult {
			index := f.NewSheet(platformName)
			f.SetActiveSheet(index)

			loc := 2

			f.SetCellValue(platformName, "A1", "关键词")
			f.SetCellValue(platformName, "B1", "次数")

			for content, count := range contentResult {
				locA := fmt.Sprintf("A%d", loc)
				locB := fmt.Sprintf("B%d", loc)

				f.SetCellValue(platformName, locA, content)
				f.SetCellValue(platformName, locB, count)

				loc++
			}
		}

		excelFileName := fmt.Sprintf("D:\\结果\\%s.xlsx", productName)
		if err := f.SaveAs(excelFileName); err != nil {
			fmt.Printf("save file %s, err: %v", excelFileName, err)
		}
	}
}
