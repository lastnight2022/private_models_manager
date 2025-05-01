package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var (
	appConfig *Config
	initOnce  sync.Once
)

func Init(path string) {
	initOnce.Do(func() {
		err := loadFile(path)
		if err != nil {
			panic(err)
		}
	})
}

func loadFile(path string) error {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("打开配置文件出错:", err)
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&appConfig)
	if err != nil {
		fmt.Println("解析配置文件出错:", err)
		return err
	}
	return nil
}

func GetConfig() Config {
	return *appConfig
}
