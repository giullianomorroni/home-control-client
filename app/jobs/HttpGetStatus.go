package jobs

import (
	srv "home-control-client/app/services"
	model "home-control-client/app/models/config"
	"github.com/robfig/revel"
)

/* see init.go */
func MyJob() {
	account := model.ReadAccount();
	srv.LookUpUpdates(account);
}
