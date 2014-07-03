package config

import (
	"os"
	"fmt"
	"bufio"
)

func filePath() (path string){
	return "/opt/homecontrol/config/data.cfg"
}

func AccountEmail() (content string) {
    filename, err := os.Open(filePath());
    defer filename.Close();

    reader := bufio.NewReader(filename);

    line, err := reader.ReadString('\n');
	for err == nil {
        fmt.Print(line)
        line, err = reader.ReadString('\n')
    }
    return line
}

func AccountPassword(data string) (content string) {
    return data
}

func WiFiName(data string) (content string) {
    return data
}

func WiFiPassword(data string) (content string) {
    return data
}