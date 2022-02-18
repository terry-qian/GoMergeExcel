# GoMergeExcel
### merge excel
1. 新建一个文件夹a
2. 将多个**同种格式**的excel放入文件夹a
2. 下载release各平台版本，然后放入文件夹a

3. 在文件夹a中新建配置文件config.yaml，配置规则如下:

   ```yaml
   fileFormat: csv
   ifMergeHeader: true
   ```

   > **fileFormat:**  需要合并的excel格式
   >
   > - csv
   > - xlsx
   >
   > **ifMergeHeader:** 是否合并表头
   >
   > - true: 合并表头，只保留第一个excel的表头
   >
   > - false: 不合并表头

4. 在该文件夹下执行对应系统的二进制文件，得到合并后的excel

   - windows，下载exe文件双击即可

   - mac，添加可执行权限
   
     ```shell
     chmod +x GoMergeExcel-darwin-M1-arm64-v2.0.0
     ./GoMergeExcel-darwin-M1-arm64-v2.0.0
     ```
   
     > mac下因没有签名无法执行的情况处理:
     >
     > 在**设置->安全性与隐私->通用**中允许GoMergeExcel进行执行
   