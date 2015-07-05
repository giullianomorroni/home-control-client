package main

import (
	"home-control-client/app"
	"os"
)

func main() {	
	account := app.ReadAccount();
	command := app.LookUpUpdates(account);
	app.UpdateDeviceStatus(command, account);
	os.Exit(0);
}
