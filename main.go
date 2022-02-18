package main

import (
	"fmt"
	"io/ioutil"

	"GoMergeExcel/config"
	"GoMergeExcel/util"
)

// 获取指定目录下的所有文件和目录
func main() {
	// 加载配置
	config.Init("")
	ifMergeHeaderConfig := config.GConfig.IfMergeHeader
	fileFormatConfig := config.GConfig.FileFormat

	// 读取当前目录文件
	dir, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("open dir failed: %s\n", err.Error())
	}

	switch fileFormatConfig {
	case "xlsx":
		util.MergeXlsx(dir, ifMergeHeaderConfig)
	case "csv":
		util.MergeCsv(dir, ifMergeHeaderConfig)
	default:
		fmt.Printf("config.yaml文件配置的fileFormat仅支持, xlsx/csv")
	}
}
