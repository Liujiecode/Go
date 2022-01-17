# Go学习代码
src:  |--go_learn  //go基础知识
    |--leetcode  //leetcode题目
    |--main   //主函数入口
    |--test   //代码测试
    |--usual  //go基础拓展
# 项目管理
## 项目管理——GOPATH：
1. GOPATH存储项目所在位置；GOROOT存储GO安装位置
2. go项目下一般三个文件夹bin、pkg、src，若要用到github的第三方库，需要在src下建github.com文件夹
3. 终端在github.com文件目录下``go get github.com/ajax/mysql``
4. 此时go get下载安装的第三方库在 GOPATH/src/github.com下
## 项目管理——GO Module（建议使用）：
1. 终端``go env -w GO111MODULE=on``，打开go module
2. 在GOPATH下的src下建立文件夹，然后``go mod init 文件夹名``，生成go.mod文件
3. 倘若需要导入github.com的第三方库，则当前项目下``go get github.com/ajax/mysql``,从而生成go.sum文件
4. 此时go get下载并安装的第三方库在GOPATH/pkg/mod/github.com下 