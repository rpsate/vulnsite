package utils

import(
	"fmt"
	"gopkg.in/ini.v1"
)

var(
	HostPort string
	Admin string
	AutoCreateAdmin bool
	JwtKey string
	Issuer string
	TokenExpire int
	DbDrive string
	DbHost string
	DbPort string
	DbUsername string
	DbPassword string
	DbDatabase string
	DbCharset string
)

func init() {
	//载入配置
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("error:", err)
	}
	LoadServer(file)
	LoadAdmin(file)
	LoadJwt(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File)  {
	HostPort = file.Section("server").Key("port").MustString(":8080")
}

func LoadAdmin(file *ini.File)  {
	Admin = file.Section("admin").Key("username").MustString("admin")
	AutoCreateAdmin = file.Section("admin").Key("autoCreate").MustBool(false)
}

func LoadJwt(file *ini.File) {
	JwtKey = file.Section("jwt").Key("JwtKey").MustString("")
	Issuer = file.Section("jwt").Key("Issuer").MustString("")
	TokenExpire = file.Section("jwt").Key("TokenExpire").MustInt(10)
}

func LoadDatabase(file *ini.File) {
	DbDrive = file.Section("database").Key("driver").MustString("mysql")
	DbHost = file.Section("database").Key("host").MustString("localhost")
	DbPort = file.Section("database").Key("port").MustString("3306")
	DbUsername = file.Section("database").Key("username").MustString("root")
	DbPassword = file.Section("database").Key("password").MustString("")
	DbDatabase = file.Section("database").Key("database").MustString("")
	DbCharset = file.Section("database").Key("charset").MustString("utf8mb4")
}