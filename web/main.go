package main

import (
	_ "my-go-web/web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

