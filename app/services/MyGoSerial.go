package services

import (
	"github.com/tarm/goserial"
	"github.com/robfig/revel"
	"log"
	"time"
	"io/ioutil"
	"net/http"
)

/*
func main() {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
	
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$start001100111234567800"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(2000 * time.Millisecond)

	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$start001100111234567801"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(2000 * time.Millisecond)
	
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$start001100111234567800"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(2000 * time.Millisecond)
	
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$start001100111234567801"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(2000 * time.Millisecond)
	
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$start001100111234567800"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
}
*/

/*
*  updates a device from a job @seeHttpGetStatus
*/
func LookUpUpdates(account string) {
	response, err := http.Get("localhost:9000/comandos?account="+account)
    if err != nil {
        revel.TRACE.Printf("%s", err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            revel.TRACE.Printf("%s", err)
        }
        revel.TRACE.Printf("%s", string(contents))
        UpdateDeviceStatus("$start" + string(contents));
    }
}

/*
* send data to raspberry
*/
func UpdateDeviceStatus(command string) () {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte(command))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(1000 * time.Millisecond)
	s.Write([]byte("$#"))
	time.Sleep(2000 * time.Millisecond)
}