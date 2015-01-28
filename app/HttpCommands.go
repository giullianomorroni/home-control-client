package app

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func LookUpUpdates(account string) (string) {
	response, err := http.Get("http://controlinside.com.br/comandos/"+account)
	defer response.Body.Close()
    if err != nil {
        fmt.Println(err)
    } else {
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Println(err)
        }
	    command := ("$start" + string(contents));
	    return command;
    }
    return "";
}