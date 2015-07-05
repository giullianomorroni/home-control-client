package app

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"strings"
)

func LookUpUpdates(account string) (string) {
	url := "http://controlinside.com.br/comandos/"+strings.TrimSpace(account)	
response, err := http.Get(url)
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
