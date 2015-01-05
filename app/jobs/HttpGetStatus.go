package jobs

import (
	srv "home-control-client/app/services"
	model "home-control-client/app/models/config"
	"github.com/robfig/revel"
	"github.com/revel/revel/modules/jobs/app/jobs"
)

type MyJob struct {}

func init() {
    revel.OnAppStart(func() { jobs.Schedule("0/10 * * * * ?", MyJob{} ) })
}

/* http://revel.github.io/manual/jobs.html */
func (j MyJob) Run() {
	account := model.ReadAccount();
	srv.LookUpUpdates(account);
}
