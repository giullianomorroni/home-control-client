package main

import (
	"github.com/tarm/goserial"
	"log"
	"time"
)

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
