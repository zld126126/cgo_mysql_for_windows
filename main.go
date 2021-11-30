package main

/*
// -Wl,--allow-multiple-definition for multiple definition
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -Wl,--allow-multiple-definition -L${SRCDIR}/lib -lmysql
#include <Windows.h>
#include <winsock.h> // for mysql-socket
#include <stdio.h> // for c.puts
#include <string.h> // for c.strlen
#include "mysql.h"
// 自定义方法
char * GoNil = NULL;
size_t GetMYSQLROWStrLen(MYSQL_ROW row,int j){
	return strlen(row[j]);
}
MYSQL_FIELD GetMYSQLFIELDItem(MYSQL_FIELD * field,int j){
	return field[j];
}
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

const (
	maxSize = 1 << 20
)

func Pause() {
	var str string
	fmt.Println("")
	fmt.Print("请按任意键继续...")
	fmt.Scanln(&str)
	fmt.Print("程序退出...")
}

func main() {
	C.puts(C.CString(" C MYSQL 使用库函数查询…… "))
	// 使用C的函数库 初始化 MYSQL *
	mysql := C.mysql_init(nil)
	if mysql == nil {
		log.Fatal("mysql is nil")
		return
	}
	// 使用库连接 MYSQL *
	C.mysql_real_connect(mysql, C.CString("127.0.0.1"), C.CString("root"), C.CString("root"), C.CString("dongbao"), C.uint(3306), C.GoNil, C.ulong(0))
	// 查询函数 int
	C.mysql_query(mysql, C.CString("select * from users"))
	// 查询结果  MYSQL_RES *
	results := C.mysql_store_result(mysql)
	// 查询的字段数目  unsigned int
	// 查询结果 char **MYSQL_ROW
	if results == nil {
		log.Fatal("result is nil")
		return
	}
	// 查询结果数目
	num_rows := int(C.mysql_num_rows(results))
	if num_rows > 0 {
		field := C.mysql_fetch_field(results)
		cfields := (*[maxSize]C.MYSQL_FIELD)(unsafe.Pointer(field))
		num_fields := int(C.mysql_num_fields(results))

		fmt.Println("num_rows:", num_rows, "num_fields:", num_fields)

		for i := 0; i < num_rows; i++ {
			var row C.MYSQL_ROW = C.mysql_fetch_row(results)
			rowPtr := (*[maxSize]*[maxSize]byte)(unsafe.Pointer(row))

			for j := 0; j < num_fields; j++ {
				fieldName_StrLen := C.strlen(cfields[j].name)
				field_name := C.GoBytes(unsafe.Pointer(C.GetMYSQLFIELDItem(field, C.int(j)).name), C.int(fieldName_StrLen))

				fieldValue_StrLen := C.GetMYSQLROWStrLen(row, C.int(j))
				field_value := C.GoBytes(unsafe.Pointer(rowPtr[j]), C.int(fieldValue_StrLen))
				if string(field_value) != "" {
					fmt.Printf("[%d]field_name is : %s \n", j, string(field_name))
					fmt.Printf("field_name is : %s , field_value is : %s \n", string(field_name), string(field_value))
				}
			}
		}
	}

	// 释放结果
	C.mysql_free_result(results)
	// 关闭mysql
	C.mysql_close(mysql)
	C.mysql_server_end()

	Pause()
}
