package main

import (
	"goproject1/api"
	"goproject1/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
