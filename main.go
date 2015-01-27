package main

import (
	"home-control-client/app"
)

func main() {	
	account := app.ReadAccount();
	command := app.LookUpUpdates(account);
	app.UpdateDeviceStatus(command, account);
}