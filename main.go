package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

// 存储字段名称和类型的结构体
type ColumnInfo struct {
	Index   int
	Field   string
	Type    string
	Null    string
	Key     string
	Default sql.NullString
	Extra   string
}

// 数据表参数传递结构
type IdlData struct {
	TableName   string
	TableName1  string
	TableItem   string
	IdItem      string
	NewItems    string
	UpdateItems string
}

// for create new
var Newitem []string

func writeidl(data IdlData) {
	// 读取模板文件
	templateFile := "idl.template"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	// 创建输出文件
	outputFile := "idl/" + data.TableName + ".thrift"
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	// 将模板渲染到输出文件
	err = tmpl.Execute(output, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("idl successfully rendered to %s\n", outputFile)

}

// 处理数据表
func readdb(dns string) {

	//connect database
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 查询所有表
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 遍历查询结果并打印表名
	fmt.Println("Tables in the database:")

	for rows.Next() {
		var data IdlData

		if err := rows.Scan(&data.TableName); err != nil {
			log.Fatal(err)
		}

		fmt.Println("-----------------------")
		fmt.Println("Table Name:", data.TableName)

		// 查询表的字段名称和类型
		rows, err := db.Query("DESCRIBE " + data.TableName)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// 遍历查询结果并输出字段名称和类型
		fmt.Printf("Columns in table %s:\n", data.TableName)
		var line string
		var i int
		var j int
		var k int

		for rows.Next() {
			var columnInfo ColumnInfo
			if err := rows.Scan(&columnInfo.Field, &columnInfo.Type, &columnInfo.Null, &columnInfo.Key, &columnInfo.Default, &columnInfo.Extra); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Name: %s, Type: %s, isNull:%s \n", columnInfo.Field, columnInfo.Type, columnInfo.Null)
			i = i + 1
			columnInfo.Index = i
			line = strconv.Itoa(int(i)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field
			data.TableItem = data.TableItem + line + "\n"

			//如果不是created,updated,deleted字段，记录在字段表中
			//为更新数准备数据
			if strings.ToLower(columnInfo.Field) == "id" {
				k = k + 1
				data.UpdateItems = data.UpdateItems + strconv.Itoa(int(k)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field + "  (api.body=\"" + columnInfo.Field + "\", api.form=\"" + columnInfo.Field + "\")" + "\n"
			}
			//为新记录准备数据
			if strings.ToLower(columnInfo.Field) != "id" && strings.ToLower(columnInfo.Field) != "created_at" && strings.ToLower(columnInfo.Field) != "updated_at" && strings.ToLower(columnInfo.Field) != "deleted_at" {
				j = j + 1
				k = k + 1
				Newitem = append(
					Newitem,
					columnInfo.Field,
				)
				data.NewItems = data.NewItems + strconv.Itoa(int(j)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field + "  (api.body=\"" + columnInfo.Field + "\", api.form=\"" + columnInfo.Field + "\")" + "\n"
				data.UpdateItems = data.UpdateItems + strconv.Itoa(int(k)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field + "  (api.body=\"" + columnInfo.Field + "\", api.form=\"" + columnInfo.Field + "\")" + "\n"
			}
		}

		// 检查遍历过程中是否出错
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		//set other items
		data.TableName1 = capitalizeFirstLetter(data.TableName)
		data.IdItem = "id" //设置默认的ID字段名

		writeidl(data)
	}

}

func main() {
	//定义命令行
	var user string
	var password string
	var host string
	var port int
	var database string
	var idlfile string

	var dns string

	flag.StringVar(&user, "user", "root", "the user name of database")
	flag.StringVar(&password, "password", "", "the user password of database")
	flag.StringVar(&host, "host", "localhost", "the host of database")
	flag.IntVar(&port, "port", 3306, "the port of database")
	flag.StringVar(&database, "database", "gorm", "the name of database")
	flag.StringVar(&idlfile, "idl", "gorm.thrift", "the name of idl file")

	//解析命令行
	flag.Parse()

	//输出解析结果
	fmt.Println("You option list below:")
	fmt.Println("user:", user)
	fmt.Println("password:", password)
	fmt.Println("host:", host)
	fmt.Println("port:", port)
	fmt.Println("database:", database)
	fmt.Println("idl file:", idlfile)

	//create dns
	dns = user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database

	fmt.Println("=======================")

	//check database
	readdb(dns)

}

func capitalizeFirstLetter(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToTitle(r[0])
	return string(r)
}

func type2idltype(t string) string {

	//提取含有数字部分的内容
	// 定义正则表达式
	re := regexp.MustCompile(`\((\d+)\)`)
	// 在字符串中查找匹配项
	match := re.FindStringSubmatch(t)

	var number int = 0
	if len(match) < 2 {
	} else {
		numberStr := match[1]
		number, _ = strconv.Atoi(numberStr)
	}

	//整数
	if strings.Contains(strings.ToLower(t), "tinyint") {
		return "i8"
	}
	if strings.Contains(strings.ToLower(t), "smallint") {
		return "i16"
	}
	if strings.Contains(strings.ToLower(t), "mediumint") {
		return "i32"
	}
	if strings.Contains(strings.ToLower(t), "int") {

		if number > 0 && number <= 8 {
			return "i8"
		}
		if number > 8 && number < 16 {
			return "i16"
		}
		if number > 16 && number < 32 {
			return "i32"
		}
		if number > 32 {
			return "i64"
		}

		return "i64"
	}
	if strings.Contains(strings.ToLower(t), "integer") {
		if number > 0 && number <= 8 {
			return "i8"
		}
		if number > 8 && number < 16 {
			return "i16"
		}
		if number > 16 && number < 32 {
			return "i32"
		}
		if number > 32 {
			return "i64"
		}

		return "i64"
	}
	if strings.Contains(strings.ToLower(t), "bigint") {
		if number > 0 && number <= 8 {
			return "i8"
		}
		if number > 8 && number < 16 {
			return "i16"
		}
		if number > 16 && number < 32 {
			return "i32"
		}
		if number > 32 {
			return "i64"
		}

		return "i64"
	}
	//浮点数
	if strings.Contains(strings.ToLower(t), "float") {
		return "double"
	}
	if strings.Contains(strings.ToLower(t), "double") {
		return "double"
	}
	if strings.Contains(strings.ToLower(t), "decimal") {
		return "double"
	}
	//布尔
	if strings.Contains(strings.ToLower(t), "boolean") {
		return "bool"
	}
	if strings.Contains(strings.ToLower(t), "real") {
		return "bool"
	}
	//字符串
	if strings.Contains(strings.ToLower(t), "varchar") {
		return "string"
	}
	if strings.Contains(strings.ToLower(t), "text") {
		return "string"
	}
	if strings.Contains(strings.ToLower(t), "char") {
		return "string"
	}
	if strings.Contains(strings.ToLower(t), "mediumtext") {
		return "string"
	}
	if strings.Contains(strings.ToLower(t), "longtext") {
		return "string"
	}

	//时间
	if strings.Contains(strings.ToLower(t), "timestamp") {
		return "i64"
	}
	if strings.Contains(strings.ToLower(t), "datetime") {
		return "i64"
	}
	if strings.Contains(strings.ToLower(t), "date") {
		return "i64"
	}
	if strings.Contains(strings.ToLower(t), "time") {
		return "i64"
	}

	return "unknown"
}
