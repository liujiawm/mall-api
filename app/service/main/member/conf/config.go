package conf

import (
	"mall/lib/database/orm"
	"os"
)

var Conf *Config

type Config struct {
	Orm  *orm.Config
	Http *HttpConfig
}

type HttpConfig struct {
	Port string
}

func Init() error {

	if os.Getenv("profile") == "production" {
		return production()
	}

	if os.Getenv("profile") == "development" {
		return development()
	}

	return local()

}

func local() error {

	Conf = &Config{
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "localhost",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf8",
			Database:     "emao_shop",
			SQLLog:       false,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "8090",
		},
	}
	return nil
}
func development() error {
	Conf = &Config{
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "172.17.222.99",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf-8",
			Database:     "bianyiquan",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "80",
		},
	}
	return nil
}

func production() error {
	Conf = &Config{
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "172.17.222.99",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf-8",
			Database:     "bianyiquan",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "80",
		},
	}
	return nil
}