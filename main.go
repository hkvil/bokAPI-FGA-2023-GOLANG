package main

import "bookapi/routers"

func main(){
	var PORT = ":8080"

	// routers.CarServer().Run(PORT)
	routers.BookServer().Run(PORT)
}