package conf

import (
	"gopkg.in/yaml.v2"
	"lovenature/log"
	"os"
	"strings"
)

var C Config

type Config struct {
	Mysql struct {
		DB         string `yaml:"db"`
		DBHost     string `yaml:"dbHost"`
		DBPort     string `yaml:"dbPort"`
		DBUser     string `yaml:"dbUser"`
		DBPassWord string `yaml:"dbPassWord"`
		DBName     string `yaml:"dbName"`
	}

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}

	QiNiu struct {
		AccessKey   string `yaml:"accessKey"`
		SecretKey   string `yaml:"secretKey"`
		Bucket      string `yaml:"bucket"`
		QiNiuServer string `yaml:"qiNiuServer"`
	}

	Es struct {
		Url      string `yaml:"url"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

func init() {
	config, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Error("配置文件读取错误", err)
		return
	}

	err = yaml.Unmarshal(config, &C)
	if err != nil {
		log.Error("反序列化文件失败", err)
		return
	}

	//配置mysql
	readPath := strings.Join([]string{C.Mysql.DBUser, ":", C.Mysql.DBPassWord, "@tcp(", C.Mysql.DBHost, ":", C.Mysql.DBPort, ")/", C.Mysql.DBName, "?charset=utf8&parseTime=true"}, "")
	writePath := strings.Join([]string{C.Mysql.DBUser, ":", C.Mysql.DBPassWord, "@tcp(", C.Mysql.DBHost, ":", C.Mysql.DBPort, ")/", C.Mysql.DBName, "?charset=utf8&parseTime=true"}, "")
	database(readPath, writePath)

	//配置Redis
	cache(C.Redis.Addr, C.Redis.Password, C.Redis.DB)

	//配置Es
	es(C.Es.Url)
}
