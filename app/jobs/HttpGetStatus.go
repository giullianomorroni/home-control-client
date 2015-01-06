package jobs

import (
	srv "home-control-client/app/services"
	model "home-control-client/app/models/config"
	"github.com/robfig/revel"
	"io/ioutil"
	"net/http"
)

/* see init.go */
func MyJob() {
	account := model.ReadAccount();
	command := LookUpUpdates(account);
	srv.UpdateDeviceStatus(command);
}

/*
*  updates a device from a job @seeHttpGetStatus
*/
func LookUpUpdates(account string) (string) {
	response, err := http.Get("http://controlinside.com.br/comandos?account="+account)
    command := "";
    if err != nil {
        revel.TRACE.Printf("%s", err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            revel.TRACE.Printf("%s", err)
        }
        revel.TRACE.Printf("%s", string(contents))
        command = ("$start" + string(contents));
    }
    return command;
}
