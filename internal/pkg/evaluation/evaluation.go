// 各平台评价内容文件
package evaluation

import (
	"diana/lg"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

// Evaluation 评价内容实体类
type Evaluation struct {
	Date              string `json:"date"`               // 日期
	PlatformName      string `json:"platform_name"`      // 平台
	ProductName       string `json:"product_name"`       // 产品名称
	EvaluationContent string `json:"evaluation_content"` // 好评/差评
	EvaluationComment string `json:"evaluation_comment"` // 评价内容
}

// ParseExcelDataFile 解释 excel 的数据文件，生成 Evaluation 对象列表
func ParseExcelDataFile(filename string) ([]Evaluation, error) {
	log := lg.GetLog()

	var evaluationList []Evaluation

	f, err := excelize.OpenFile(filename)

	if err != nil {
		return nil, errors.Wrap(err, "读取 excel 文件失败")
	}

	log.WithFields(logrus.Fields{
		"filename": filename,
	}).Info("开始导入 ....")

	for _, sheetName := range f.GetSheetMap() {
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
						// 好评 / 差评
						evaluation.EvaluationContent = colCell

					case 4:
						// 评价内容
						evaluation.EvaluationComment = colCell
					}
				}

				evaluationList = append(evaluationList, evaluation)
			}
		}

	}

	log.WithFields(logrus.Fields{
		"filename": filename,
	}).Info("导入完毕 ....")

	return evaluationList, nil
}

// ExportExcelResultFile 导出结果文件
func ExportExcelResultFile(evaluationList []Evaluation) error {

	return nil
}
