package app

import (
    "io/ioutil"
	"strings"
	"os/exec"
	"fmt"
	"regexp"
)

/* FILE CONFIG
email@mail.com#
senha_app#
nome_rede#
senha_rede#
*/

func filePath() (path string){
	return "/opt/homecontrol/config/data.cfg"
}

func fileAccountPath() (path string){
	return "/opt/homecontrol/config/account.cfg"
}

func holdHoleShit(err error) (){
	if err != nil {
        panic(err)
    }
}

func readConfig(file string) (content string) {
	dat, err := ioutil.ReadFile(file)
	holdHoleShit(err)
    config := string(dat)
    return config
}

func ReadAccount() (content string) {
	data := readConfig(fileAccountPath())
	return data
}

func ReadAccountEmail() (content string) {
	data := strings.Split(readConfig(filePath()), "#")
	return data[0]
}

func ReadAccountPassword() (content string) {
	data := strings.Split(readConfig(filePath()), "#")
	return data[1]
}

func ReadWiFiName() (content string) {
	data := strings.Split(readConfig(filePath()), "#")
	return data[2]
}

func ReadWiFiPassword() (content string) {
	data := strings.Split(readConfig(filePath()), "#")
	return data[3]
}

func WriteAccountEmail(content string) () {
	data := strings.Split(readConfig(filePath()), "#")
	data[0] = content ;
	write(0, data)
}

func WriteAccountPassword(content string) () {
	data := strings.Split(readConfig(filePath()), "#")
	data[1] = content ;
	write(1, data)
}

func WriteWiFiName(content string) () {
	data := strings.Split(readConfig(filePath()), "#")
	data[2] = content ;
	write(2, data)
}

func WriteWiFiPassword(content string) () {
	data := strings.Split(readConfig(filePath()), "#")
	data[3] = content ;
	write(3, data)
}

func write(position int, data []string) () {
	aux := "";
	for i := 0; i < 4; i++ {
		aux = aux + data[i] + "#";
	}
	d1 := []byte(aux)
    err := ioutil.WriteFile(filePath(), d1, 0777)
    holdHoleShit(err);
}

func ListNetWorksAvailable() (data []string){
	out, err := exec.Command("iwlist", "wlan0", "scan").Output();
	fmt.Println(err)

	var validID = regexp.MustCompile("ESSID:.*")
	data = validID.FindAllString(string(out[:]), 5)

	return data;
}
