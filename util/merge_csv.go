package util

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func MergeCsv(dir []fs.FileInfo, ifMergeHeaderConfig bool) {
	var newErr error
	var fileList []*os.File

	for _, fi := range dir {
		if !strings.Contains(fi.Name(), ".csv") {
			continue
		}
		if newErr != nil {
			fmt.Printf("error: %s", newErr.Error())
		}
		// 读取文件
		file, err := os.Open("./" + fi.Name())

		if err != nil {
			fmt.Printf("open failed: %s\n", err)
		}
		fileList = append(fileList, file)
	}

	// 合并后的
	var rows [][]string

	if ifMergeHeaderConfig {
		for index, item := range fileList {
			reader := csv.NewReader(item)
			// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
			reader.FieldsPerRecord = -1
			if index < 1 {
				records, _ := reader.ReadAll()
				rows = append(rows, records...)
			} else {
				records, _ := reader.ReadAll()
				// 删除第一行表头
				records = records[1:][:]
				rows = append(rows, records...)
			}
		}

	} else {
		for _, item := range fileList {
			reader := csv.NewReader(item)
			reader.FieldsPerRecord = -1
			records, _ := reader.ReadAll()
			rows = append(rows, records...)
		}
	}

	// 写入
	newFile, err := os.Create("merge.csv")
	if err != nil {
		panic(err)
	}
	//异步管理
	defer newFile.Close()
	// 写入UTF-8 BOM
	_, err2 := newFile.WriteString("\xEF\xBB\xBF")
	if err2 != nil {
		panic(err2)
	}
	w := csv.NewWriter(newFile)
	err = w.WriteAll(rows)
	if err != nil {
		panic(err)
	}
	w.Flush()
}
