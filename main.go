package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

// 定义命令行
var user string
var password string
var host string
var port int
var database string
var idlfile string
var project string
var idltype string
var camelCase bool
var ignoreTablesStr string
var ignoreTables []string
var help string

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

// for const var
// var 	gorootenv string = os.Getenv("GOROOT")
var gopathenv string = os.Getenv("GOPATH") //GOPATH 路径
var dsn string                             //数据库的 DSN

//define exec command

var cwserver string = "cwgo server --type HTTP --service $service --module $module --idl ../../idl/$idl"

//var cwdb string = "cwgo model --db_type mysql  --type_tag true --out_dir biz/model --dsn $dsn"

func writeidl(data IdlData) {

	templateFile := "thrift.template"
	// 读取模板文件
	if idltype != "t" {
		templateFile = "proto.template"
	}

	paths := strings.Split(gopathenv, string(os.PathListSeparator))
	tmpfile := ""
	tmpok := false
	_, err := os.Stat(templateFile)
	if err != nil {
		for _, path := range paths {
			tmpfile = filepath.Join(path, "mysql2idl", "thrift.template")
			if idltype != "t" {
				tmpfile = filepath.Join(path, "mysql2idl", "proto.template")
			}

			_, err = os.Stat(tmpfile)
			if err == nil {
				templateFile = tmpfile
				tmpok = true
				break
			}
		}
	} else {
		tmpok = true
	}

	if !tmpok {
		log.Fatal("template file not found!")
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
		fmt.Println("err on Parse idl file:", err)
	}

	// 创建输出文件
	_, err = os.Stat("idl")
	if err != nil {
		err = os.MkdirAll("idl", os.ModePerm)
		if err != nil {
			fmt.Println("err on create idl fold:", err)
		}
	}

	outputFile := "idl/" + data.TableName + ".thrift"
	if idltype != "t" {
		outputFile = "idl/" + data.TableName + ".proto"
	}

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

	//如果是protobuf文件还要把api.proto复制到 idl下
	if idltype != "t" {
		copyProtobufInclude()
	}

	log.Printf("idl successfully rendered to %s\n", outputFile)

}

func copyProtobufInclude() {
	files := []string{"api.proto", "base.proto"}

	err := os.MkdirAll("idl/include/", 0755)
	if err != nil {
		fmt.Println("Error creating directory idl/include", err)
		return
	}

	for _, file := range files {
		in, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error copying file:", file, err)
			return
		}
		err = os.WriteFile("idl/include/"+file, in, 0644)
		if err != nil {
			fmt.Println("Error copying file:", file, err)
			return
		}
	}
}

// 处理数据表
func readdb(d string) {

	//connect database
	db, err := sql.Open("mysql", d)
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

	// 创建输出文件
	_, err = os.Stat("idl")
	if err != nil {
		err = os.MkdirAll("idl", os.ModePerm)
		if err != nil {
			fmt.Println("err on create idl fold:", err)
		}
	}

	//清空 idl生成的目录
	err = clearDirectory("idl")
	if err != nil {
		fmt.Println("Error clearing directory:", err)
		log.Fatal(err)
	} else {
		//fmt.Println("Directory cleared successfully.")
	}

	// 遍历查询结果并打印表名
	fmt.Println("Tables in the database:")

	for rows.Next() {
		var data IdlData

		if err := rows.Scan(&data.TableName); err != nil {
			log.Fatal(err)
		}

		skipTable := false
		for _, ignoreTbl := range ignoreTables {
			if data.TableName == ignoreTbl {
				skipTable = true
				break
			}
		}
		if skipTable {
			continue
		}

		//fmt.Println("-----------------------")
		//fmt.Println("Table Name:", data.TableName)

		// 查询表的字段名称和类型
		rows, err := db.Query("DESCRIBE `" + data.TableName + "`")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// 遍历查询结果并输出字段名称和类型
		//fmt.Printf("Columns in table %s:\n", data.TableName)
		var line string
		var i int
		var j int
		var k int

		for rows.Next() {
			var columnInfo ColumnInfo
			if err := rows.Scan(&columnInfo.Field, &columnInfo.Type, &columnInfo.Null, &columnInfo.Key, &columnInfo.Default, &columnInfo.Extra); err != nil {
				log.Fatal(err)
			}

			//fmt.Printf("Name: %s, Type: %s, isNull:%s \n", columnInfo.Field, columnInfo.Type, columnInfo.Null)
			i = i + 1
			columnInfo.Index = i
			line = strconv.Itoa(int(i)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field
			if idltype != "t" {
				line = type2idltype(columnInfo.Type) + " " + columnInfo.Field + " = " + strconv.Itoa(int(i)) + ";"
			}

			data.TableItem = data.TableItem + line + "\n  "

			//如果不是created,updated,deleted字段，记录在字段表中
			//为更新数准备数据
			if strings.ToLower(columnInfo.Field) == "id" {
				k = k + 1
				if idltype != "t" {
					data.UpdateItems = data.UpdateItems + type2idltype(columnInfo.Type) + " " + columnInfo.Field + " = " + strconv.Itoa(int(k)) + ";\n  "

				} else {
					data.UpdateItems = data.UpdateItems + strconv.Itoa(int(k)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field + "  (api.body=\"" + columnInfo.Field + "\", api.form=\"" + columnInfo.Field + "\")" + "\n"
				}

			}
			//为新记录准备数据
			if strings.ToLower(columnInfo.Field) != "id" && strings.ToLower(columnInfo.Field) != "created_at" && strings.ToLower(columnInfo.Field) != "updated_at" && strings.ToLower(columnInfo.Field) != "deleted_at" {
				j = j + 1
				k = k + 1
				Newitem = append(
					Newitem,
					columnInfo.Field,
				)
				if idltype != "t" {
					data.NewItems = data.NewItems + type2idltype(columnInfo.Type) + " " + columnInfo.Field + " = " + strconv.Itoa(int(j)) + ";\n  "
					data.UpdateItems = data.UpdateItems + type2idltype(columnInfo.Type) + " " + columnInfo.Field + "  = " + strconv.Itoa(int(k)) + ";\n  "

				} else {
					data.NewItems = data.NewItems + strconv.Itoa(int(j)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field + "  (api.body=\"" + columnInfo.Field + "\", api.form=\"" + columnInfo.Field + "\")" + "\n"
					data.UpdateItems = data.UpdateItems + strconv.Itoa(int(k)) + ":" + type2idltype(columnInfo.Type) + " " + columnInfo.Field + "  (api.body=\"" + columnInfo.Field + "\", api.form=\"" + columnInfo.Field + "\")" + "\n"

				}

			}
		}

		// 检查遍历过程中是否出错
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		//set other items

		if camelCase {
			data.TableName1 = makeCamelCase(data.TableName)
		} else {
			data.TableName1 = capitalizeFirstLetter(data.TableName)
		}
		data.IdItem = "id" //设置默认的ID字段名
		data.TableItem = strings.TrimSpace(data.TableItem)

		writeidl(data)
	}

}

func main() {

	//是否生成项目

	flag.StringVar(&user, "user", "root", "the user name of database")
	flag.StringVar(&password, "password", "123456", "the user password of database")
	flag.StringVar(&host, "host", "127.0.0.1", "the host of database")
	flag.IntVar(&port, "port", 3306, "the port of database")
	flag.StringVar(&database, "database", "gorm", "the name of database")
	flag.StringVar(&idlfile, "idl", "gorm.thrift", "the name of idl file")
	flag.StringVar(&project, "project", "", "the project name for generate,if blank then will not generate project")
	flag.StringVar(&idltype, "idltype", "t", "the idl type for generate, t=thrift, others=protobuf")
	flag.StringVar(&ignoreTablesStr, "ignoretables", "", "tables to ignore when generating idl files. separate with a space.")
	flag.BoolVar(&camelCase, "camelcase", false, "flags whether to use camel or snake case for the database name")
	flag.StringVar(&help, "help", "", "help message")

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
	fmt.Println("ignoring tables: ", ignoreTablesStr)
	fmt.Println("using camel case", camelCase)

	if project != "" {
		fmt.Println("go project will generate into fold projects/", project)
	} else {
		fmt.Println("go project will not generated!")
	}

	// clean up table names so they properly match
	ignoreTables = strings.Split(ignoreTablesStr, " ")
	for i, tbl := range ignoreTables {
		ignoreTables[i] = strings.ToLower(strings.TrimSpace(tbl))
	}

	//create dns
	dsn = user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println("dsn=", dsn)
	//check database & generate idl
	fmt.Println("=========Step1: Generate IDL files==============")
	readdb(dsn)
	//use cwgo generate project
	if project != "" {
		fmt.Println("=========Step2: Generate Project files==============")
		_ = createproject(project)
	}

}

func createproject(project string) bool {
	//检查 是否存在项目目录
	project = strings.TrimSpace(project) //去除空格
	_, err := os.Stat("projects/" + project)
	if err != nil {
		err = os.MkdirAll("projects/"+project, os.ModePerm)
		if err != nil {
			fmt.Println("err on create project:", err)
			return false
		}
	} else {
		fmt.Println("project exist!, if you want to create again, delete the project first!")
		return false
	}

	//检查 cwgo是否存在  %GOROOT%/bin目录下
	cwgocmd := "cwgo"
	if os.PathSeparator == '\\' {
		cwgocmd += ".exe"
	}
	paths := strings.Split(gopathenv, string(os.PathListSeparator))
	tmpfile := ""
	tmpok := false
	_, err = os.Stat(cwgocmd)
	if err != nil {
		for _, path := range paths {
			tmpfile = filepath.Join(path, "bin", cwgocmd)
			_, err = os.Stat(tmpfile)
			if err == nil {
				cwgocmd = tmpfile
				tmpok = true
				break
			}
		}
	} else {
		tmpok = true
	}

	if !tmpok {
		fmt.Println("cwgo command is not installed , please install it as below:\n go install github.com/cloudwego/cwgo@latest")
		return false
	} else {
		fmt.Println("cwgo found!")
	}
	//进入项目目录
	err = os.Chdir(filepath.Join("projects", project))
	// 用cwgo 生成项目文件
	currentDir, err := os.Getwd()
	fmt.Println("Current project path:", currentDir)

	//执行cwgo生成代码
	genprojectservice(project)

	return true
}

func genprojectservice(project string) {
	cmdName := "cwgo"
	myidl := "users.thrift"
	idlpath := "../../idl"

	cmdArgs := []string{
		"server",
		"--type",
		"HTTP",
		"--service",
		project,
		"--module",
		project,
		"--idl",
		filepath.Join(idlpath, myidl),
	}

	dir, err := os.Open(idlpath)
	if err != nil {
		fmt.Println("Error opening idl directory:", err)
		return
	}
	defer dir.Close()

	// 读取目录下的所有文件
	idlInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, idlInfo := range idlInfos {
		myidl = idlInfo.Name()
		cmdArgs = []string{
			"server",
			"--type",
			"HTTP",
			"--service",
			project,
			"--module",
			project,
			"--idl",
			filepath.Join(idlpath, myidl),
		}
		fmt.Println(cmdArgs)
		cmd := exec.Command(cmdName, cmdArgs...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing command cwgo:", err)
			fmt.Println("cwgo args:", cmdArgs)
			//log.Fatal(err)
		}
		fmt.Println(string(output))
	}

}

func capitalizeFirstLetter(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToTitle(r[0])
	return string(r)
}

func makeCamelCase(s string) string {
	if s == "" {
		return ""
	}
	spl := strings.Split(s, "_")
	for i, word := range spl {
		word = capitalizeFirstLetter(word)
		spl[i] = word
	}
	return strings.Join(spl, "")
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
		return toproto("i8")
	}
	if strings.Contains(strings.ToLower(t), "smallint") {
		return toproto("i16")
	}
	if strings.Contains(strings.ToLower(t), "mediumint") {
		return toproto("i32")
	}
	if strings.Contains(strings.ToLower(t), "int") {

		if number > 0 && number <= 8 {
			return toproto("i16")
		}
		if number > 8 && number < 16 {
			return toproto("i16")
		}
		if number > 16 && number < 32 {
			return toproto("i32")
		}
		if number > 32 {
			return toproto("i64")
		}

		return toproto("i32")
	}
	if strings.Contains(strings.ToLower(t), "integer") {
		if number > 0 && number <= 8 {
			return toproto("i16")
		}
		if number > 8 && number < 16 {
			return toproto("i16")
		}
		if number > 16 && number < 32 {
			return toproto("i32")
		}
		if number > 32 {
			return toproto("i64")
		}

		return toproto("i32")
	}
	if strings.Contains(strings.ToLower(t), "bigint") {
		if number > 0 && number <= 8 {
			return toproto("i16")
		}
		if number > 8 && number < 16 {
			return toproto("i16")
		}
		if number > 16 && number < 32 {
			return toproto("i32")
		}
		if number > 32 {
			return toproto("i64")
		}

		return toproto("i32")
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
		return toproto("i64")
	}
	if strings.Contains(strings.ToLower(t), "datetime") {
		return toproto("i64")
	}
	if strings.Contains(strings.ToLower(t), "date") {
		return toproto("i64")
	}
	if strings.Contains(strings.ToLower(t), "time") {
		return toproto("i64")
	}

	return "unknown"
}

// 转换 thrift变量和proto变量
func toproto(str string) string {
	if idltype != "t" {
		switch str {
		case "i16":
			return "int32"
		case "i32":
			return "int32"
		case "i64":
			return "int64"
		case "binary":
			return "bytes"
		case "struct":
			return "message"
		default:
			return "unknown"
		}
	}
	return str
}

// 删除目录 dirPath下的所有文件和子目录
func clearDirectory(dirPath string) error {
	// 获取目录下的所有文件和子目录
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// 遍历目录项
	for _, entry := range entries {
		entryPath := filepath.Join(dirPath, entry.Name())

		// 删除文件或目录
		err := os.RemoveAll(entryPath)
		if err != nil {
			return err
		}
	}

	return nil
}
