package main 

import (
	"home-control-client/app/models/config"
)

func main() {
	/*
	d1 := config.ReadAccountEmail();
	config.WriteAccountEmail("teste@teste.com");
	d2 := config.ReadAccountEmail();
	fmt.Println(d1)
	fmt.Println(d2)
	*/
	config.ListNetWorksAvailable()
}
