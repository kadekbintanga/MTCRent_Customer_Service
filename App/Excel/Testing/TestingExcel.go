package Testing

import (
	"Service/App/Service/Constant/Path"
	excel2 "github.com/globalxtreme/gobaseconf/excel"
	"github.com/globalxtreme/gobaseconf/helpers"
	"github.com/xuri/excelize/v2"
)

type TestingExcel struct {
}

func (ex TestingExcel) Generate() error {
	sheets, properties := ex.setSheetsAndProperties()

	excel := ex.newFile(sheets, properties)
	excel = ex.modifySheet(excel)

	err := excel.Save(Path.PathExcelTesting(), "testing.xlsx")
	if err != nil {
		return err
	}

	return nil
}

func (ex TestingExcel) newFile(sheets []string, properties [][][]interface{}) excel2.Excel {
	excel := excel2.Excel{
		Sheets:     sheets,
		Properties: properties,
		IsPublic:   false,
	}

	excel.NewFile()

	return excel
}

func (ex TestingExcel) modifySheet(excel excel2.Excel) excel2.Excel {
	excel.SetStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "860A35", Style: 1},
			{Type: "right", Color: "860A35", Style: 1},
			{Type: "top", Color: "860A35", Style: 1},
			{Type: "bottom", Color: "860A35", Style: 1},
		},
	}, "A1:C1")

	excel.SetStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "A3B763", Style: 2},
			{Type: "right", Color: "A3B763", Style: 2},
			{Type: "top", Color: "A3B763", Style: 2},
			{Type: "bottom", Color: "A3B763", Style: 2},
		},
	}, "A3:C4")

	excel.MergeCells("A5:C5", "A6:A7")

	excel.SetWidthCols([]excel2.ColWidth{
		{Cells: "A", Width: 5},
		{Cells: "B", Width: 15},
		{Cells: "C", Width: 5},
	})

	excel.SetHeightRows([]excel2.RowHeight{
		{Row: 1, Height: 15},
	})

	return excel
}

func (TestingExcel) setSheetsAndProperties() ([]string, [][][]interface{}) {
	var sheets []string
	var properties [][][]interface{}

	sheets = append(sheets, "Testing 1", "Testing 2")
	for _ = range sheets {
		properties = append(properties, [][]interface{}{
			{"ID", "Name", "Age"},
		})
	}

	dataProperties := [][]int{
		{29, 23, 12},
		{19, 26, 14},
	}

	for sKey, property := range dataProperties {
		for pKey, val := range property {
			properties[sKey] = append(properties[sKey], []interface{}{pKey + 1, helpers.RandomString(10), val})
		}
	}

	return sheets, properties
}
