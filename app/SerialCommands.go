package app

import (
	"github.com/tarm/goserial"
	"fmt"
	"time"
)

/*
* send data to raspberry
*/
func UpdateDeviceStatus(command, account string) () {
	if command == "$start" {
		return;
	}
	fmt.Println("Executing Command: "+command);
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}

	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
		return;
	}

	s.Write([]byte("$#"))
	time.Sleep(2000 * time.Millisecond)
	s.Write([]byte(command))
	time.Sleep(500 * time.Millisecond)
	s.Write([]byte("$#"))
	//time.Sleep(1000 * time.Millisecond)
	//s.Write([]byte("$#"))
	//time.Sleep(1000 * time.Millisecond)
	fmt.Println("Everything executed");
}
