package main

import "fmt"

var cfg Config

func main() {
	loadCfgFile(&cfg)
	loadEnv(&cfg)
	fmt.Println(cfg)
}
