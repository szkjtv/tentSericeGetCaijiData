package main

import (

	// "example.com/m/v2/model"

	model "example.com/m/v2/model/UCBrowser"
	"example.com/m/v2/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	model.Dbinit()

}

func main() {
	router.Router()
	// for {
	// 	time.Sleep(0 * time.Microsecond)
	// }

}
