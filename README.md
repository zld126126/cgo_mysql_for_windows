#cgo_mysql-for-windows
[cgo_mysql-for-mac](https://studygolang.com/articles/02036)
## 1.安装环境
    mingw && mingw-utils
    mysql:执行users.sql
    golang 1.13
## 2.编译 && 运行
## 2.1 dll 编译 def
    pexports libmysql.dll > libmysql.def
## 2.2 def && dll 编译 a
    dlltool.exe -D libmysql.dll -d libmysql.def -l libmysql.a -k
## 2.3 编译运行exe
    go build -x

## 3.出错处理:
### 3.1 In function `cgo_3a86e410a0e3_Cfunc_mysql_close':
    /tmp/go-build/cgo-gcc-prolog:87: undefined reference to `mysql_close@4'
解决方案:

    找到libmysql.def文件覆盖
        "mysql_close" >>>> "mysql_close@4"
        
### 3.2 cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
解决方案:

    WINDOWS:
        SET CGO_ENABLED=1
        SET GOOS=windows
        SET GOARCH=386
        
    GOLAND-Environment:
        CGO_ENABLED=1;GOOS=windows;GOARCH=386
   