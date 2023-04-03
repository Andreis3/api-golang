package main

import "github.com/andreis3/api-golang/configs"

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(configs.DBDriver)
}
