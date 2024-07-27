# mysql2idl
## this Utility  is for convert mysql table to idl files
## based on Cloudwego.io hertz
### Program by clark zhu (full stack developer)
### version 0.2  
### issue date : 2024-07-26
### email : zhuclark2020@gmail.com
### thanks to cloudwego.io tiktok toutiao and all the gophers
### the final idl file could support generate golang,python,java,rust project

You could create idl(thrift format) file with this tools just from any mysql database
before you use this tools , you must know some useful tools about hertz(golang), volo(rust) ,apache-thrift(java,python), thriftgo, cwgo etc. 

#### ==========================================================
### Usage:
  mysql2idl -user root -password 123456 -host 127.0.0.1 -port 3306 -database gorm -idlfile user.thrift  -idltype p -table user -project test
or mysql2idl -user root -password 123456 -host 127.0.0.1 -port 3306 -database gorm -project test

### Parameter:  
  -user the database username  
  -password the database password  
  -host    the database service host  
  -port    the database service port  
  -database the database name  
  -idlfile   the filename of idl which will be generated.  
  -idltype   the type of idl which will be generated default is t as thrift file , if not generate protobuf file.
  -table    the table name which will be converted to idl if no this parameter means all tables will be generated.
  -project  if this parameter is set , it will call cwgo to generate golang project.

[if you put the exec file mysql2idl or mysql2idl.exe to path and you must put api.proto and thrift.template and proto.template to go work path as  mysql2idl folder under go_work path]

#### How to use mysql2idl to generate golang hertz project,you also could generate java, python, rust project with apache-thrift and volo
#### step1:
     Example: your mysql database as below:
     host: 127.0.0.1
     port: 3306
     user: root
     password : 123456
     database : demo
     there are a table name is : users
     
     input command as below:
     
     c:\example>mysql2idl.exe -host 127.0.0.1 -port 3306 -user root -password 123456 -database demo -idlfile users.thrift -table users -project test
     (it will generate idl file users.thrift in the idl path and generate whole project files in project/test path , all done netxt step just give up)
     

#### step2:
     Generate project files in test path as below:
     c:\example>mkdir test && cd test
     c:\example>hz new -module demo -idl ../idl/users.thrift
     (then you will find all the hz projects files generate into the test path)
#### step3:
     Execute the example demo
     
     c:\example\test>go mod tidy
     c:\example\test>go run .

#### ==========================================================
#### Update information
2023-12-29: issue first version 0.1