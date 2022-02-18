package util

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/tealeg/xlsx"
)

func MergeXlsx(dir []fs.FileInfo, ifMergeHeaderConfig bool) {

	var newErr error
	var fileList []*xlsx.File

	for _, fi := range dir {
		if !strings.Contains(fi.Name(), ".xlsx") {
			continue
		}
		// fmt.Printf("open success: %s\n", Pthdir+PthSep+fi.Name())
		if newErr != nil {
			fmt.Printf("error: %s", newErr.Error())
		}

		// 读取文件
		xlFile, err := xlsx.OpenFile("./" + fi.Name())
		if err != nil {
			fmt.Printf("open failed: %s\n", err)
		}
		fileList = append(fileList, xlFile)
	}

	// 申明合并后的文件
	var newFile *xlsx.File
	var newSheet *xlsx.Sheet
	newFile = xlsx.NewFile()
	newSheet, _ = newFile.AddSheet("Sheet1")

	for i, item := range fileList {
		for _, sheet := range item.Sheets {
			// fmt.Printf("Sheet Name: %s\n", sheet.Name)
			// fmt.Printf("index: %d\n", i)
			for num, row := range sheet.Rows {
				// fmt.Printf("num: %d\n", num)
				// 第二个开始跳过第一行表头，将后面的行写入新的文件
				if ifMergeHeaderConfig {
					if i < 1 {
						newRow := newSheet.AddRow()
						newRow.SetHeightCM(1)
						for _, cell := range row.Cells {
							text := cell.String()
							// fmt.Printf("%s\n", text)

							newCell := newRow.AddCell()
							newCell.Value = text
						}
					} else {
						if num > 0 {
							newRow := newSheet.AddRow()
							newRow.SetHeightCM(1)
							for _, cell := range row.Cells {
								text := cell.String()
								// fmt.Printf("%s\n", text)

								newCell := newRow.AddCell()
								newCell.Value = text
							}
						}
					}
				} else {
					newRow := newSheet.AddRow()
					newRow.SetHeightCM(1)
					for _, cell := range row.Cells {
						text := cell.String()
						// fmt.Printf("%s\n", text)

						newCell := newRow.AddCell()
						newCell.Value = text
					}
				}
			}
		}
	}
	// 写入文件
	err := newFile.Save("./merge.xlsx")
	if err != nil {
		panic(err)
	}
}
