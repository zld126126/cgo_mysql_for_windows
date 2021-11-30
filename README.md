# cgo-mysql-for-windows
- [cgo-mysql-for-windows](#cgo-mysql-for-windows)
  - [1.安装环境](#1安装环境)
  - [2.编译 && 运行](#2编译--运行)
  - [2.1 dll 编译 def](#21-dll-编译-def)
  - [2.2 def && dll 编译 a](#22-def--dll-编译-a)
  - [2.3 编译运行exe](#23-编译运行exe)
  - [3.出错处理:](#3出错处理)
    - [3.1 In function `cgo_3a86e410a0e3_Cfunc_mysql_close':](#31-in-function-cgo_3a86e410a0e3_cfunc_mysql_close)
    - [3.2 cc1.exe: sorry, unimplemented: 64-bit mode not compiled in](#32-cc1exe-sorry-unimplemented-64-bit-mode-not-compiled-in)
  - [4.github地址](#4github地址)
  - [5.相关资料](#5相关资料)
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

## 4.github地址
[cgo_mysql_for_windows](https://github.com/zld126126/cgo_mysql_for_windows)

## 5.相关资料
[cgo-mysql-for-mac](https://studygolang.com/articles/02036)