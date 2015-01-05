package app

import (
	"github.com/robfig/revel"
    "home-control-client/app/jobs"
)

//** PRIVATE FUNCTIONS
func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.ActionInvoker,           // Invoke the action.
	}

	// Called to help init the application
	revel.OnAppStart(initApp)
}

// initApp contains all application level initialization
func initApp() {
    c := cron.New()
    c.AddFunc("0/10 * * * * *", jobs.MyJob())
    c.Start();
}
