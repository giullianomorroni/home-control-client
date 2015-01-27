package app

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func init() {
	//account := model.ReadAccount();
	//command := LookUpUpdates(account);
	//srv.UpdateDeviceStatus(command);
}

func LookUpUpdates(account string) (string) {
	//response, err := http.Get("http://controlinside.com.br/comandos?account="+account)
	response, err := http.Get("http://localhost:9000/comandos?account="+account)
    command := "";
    if err != nil {
        fmt.Println(err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Println(err)
        }
        command = ("$start" + string(contents));
    }
    return command;
}
