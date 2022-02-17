# GoMergeExcel
### merge excel
1. 将多个excel放入一个文件夹
   
- excel格式仅支持.xlsx
   
2. 下载release各平台版本，然后放入以上文件夹

3. 在以上文件夹中新建config.yaml

   ```yaml
   ifMergeHeader: true
   ```

   > true: 合并表头，只保留第一个excel的表头
   >
   > false: 不合并表头

4. 在该文件夹下执行对应系统的二进制文件，得到合并后的excel

   - 例如：

   ```shell
   ./GoMergeExcel-darwin-M1-arm64-v1.0.0
   ```