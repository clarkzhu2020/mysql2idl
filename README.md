# mysql2idl
## this util tools is for convert mysql table to idl files
## based on Cloudwego.io hertz
### Program by clark zhu (full stack developer)
### version 0.1  
### issue date : 2023-12-29
### email : zhuclark2020@gmail.com
### thanks to cloudwego.io tiktok toutiao and all the gophers

You could create idl(thrift format) file with this tools just from any mysql database
before you use this tools , you must know some useful tools about hertz, thriftgo, cwgo etc. 

#### ==========================================================
### Usage:
  mysql2idl -user root -password 123456 -host 127.0.0.1 -port 3306 -database gorm -idlfile user.thrift

### Parameter:  
  -user the database username  
  -password the database password  
  -host    the database service host  
  -port    the database service port  
  -database the database name  
  -idlfile   the filename of idl which will be generated.  
  
#### ==========================================================
#### Update information
2023-12-29: issue first version 0.1