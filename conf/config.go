package conf

import (
	"gopkg.in/yaml.v2"
	"lovenature/log"
	"os"
	"strings"
)

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
}

func init() {
	config, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Error("配置文件读取错误", err)
		return
	}

	var c Config
	err = yaml.Unmarshal(config, &c)
	if err != nil {
		log.Error("反序列化文件失败", err)
		return
	}

	//配置mysql
	readPath := strings.Join([]string{c.Mysql.DBUser, ":", c.Mysql.DBPassWord, "@tcp(", c.Mysql.DBHost, ":", c.Mysql.DBPort, ")/", c.Mysql.DBName, "?charset=utf8&parseTime=true"}, "")
	writePath := strings.Join([]string{c.Mysql.DBUser, ":", c.Mysql.DBPassWord, "@tcp(", c.Mysql.DBHost, ":", c.Mysql.DBPort, ")/", c.Mysql.DBName, "?charset=utf8&parseTime=true"}, "")
	Database(readPath, writePath)

	//配置Redis
	Cache(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
}
