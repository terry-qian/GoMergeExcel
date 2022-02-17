package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"GoMergeExcel/config"

	"github.com/tealeg/xlsx"
)

// 获取指定目录下的所有文件和目录
func main() {
	// 加载配置
	config.Init("")

	// 设置路径，文件夹放在main的同级目录下
	PthSep := string(os.PathSeparator)
	Pthdir := "."

	dir, err := ioutil.ReadDir(Pthdir)
	if err != nil {
		fmt.Printf("open dir failed: %s\n", err.Error())
	}

	// 申明合并后的文件
	var newFile *xlsx.File
	var newSheet *xlsx.Sheet
	newFile = xlsx.NewFile()
	var newErr error
	newSheet, newErr = newFile.AddSheet("Sheet1")
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
		xlFile, err := xlsx.OpenFile(Pthdir + PthSep + fi.Name())
		if err != nil {
			fmt.Printf("open failed: %s\n", err)
		}
		fileList = append(fileList, xlFile)
	}

	for i, item := range fileList {
		for _, sheet := range item.Sheets {
			// fmt.Printf("Sheet Name: %s\n", sheet.Name)
			// fmt.Printf("index: %d\n", i)
			for num, row := range sheet.Rows {
				// fmt.Printf("num: %d\n", num)
				// 第二个开始跳过第一行表头，将后面的行写入新的文件
				if config.GConfig.IfMergeHeader {
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
	newErr = newFile.Save("./merge.xlsx")
	if newErr != nil {
		fmt.Printf("error: %s", newErr.Error())
	}
}
