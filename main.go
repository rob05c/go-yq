package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func yamlValue(file string, key string) string {
	viper.SetConfigType("yaml")

	verifyKey(key)
	keyWithoutFirstDot := strings.Replace(key, ".", "", 1)

	filename := filename(file)
	viper.SetConfigName(filename)

	dir := dir(file)
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	value := fmt.Sprintf("%s", viper.Get(keyWithoutFirstDot))

	return value
}

func dir(path string) string {
	return filepath.Dir(path)
}

func filename(path string) string {
	basename := filepath.Base(path)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))

	return filepath.Base(filename)
}

func verifyKey(key string) {
	if !strings.HasPrefix(key, ".") {
		log.Fatal("Key should start with a dot, i.e.: ."+key+", but was: ", key)
	}
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("Usage: go-yq <key e.g. .foo.bar> <filename e.g. input.yaml>")
	}

	key := os.Args[1]
	yamlFile := os.Args[2]

	value := yamlValue(yamlFile, key)
	if value == "" {
		log.Fatal("File: ", yamlFile, " does not contain key: ", key)
	}
	fmt.Println(value)
}
