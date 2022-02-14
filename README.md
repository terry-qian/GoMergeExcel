# GoMergeExcel
### merge excel
1. 将多个excel放入一个文件夹
   - excel格式仅支持.xlsx

2. 构建可执行二进制文件，没有go编译环境可直接下载release各平台版本，然后放入以上文件夹

   ```shell
   make build
   ```

3. 在该文件夹下执行对应系统的二进制文件，得到合并后的excel
   - 例如：

   ```shell
   ./GoMergeExcel-mac-M1-arm64-v1.0.0
   ```