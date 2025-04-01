package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Password string `toml:"password"`
	Nickname string `toml:"nickname"`
	RealName string `toml:"realname"`
}

func loadConfig() (*Config, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(filepath.Dir(exePath), "config.toml")

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	server := "coco.local:6667"
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Could not connect to server")
		return
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	writer.WriteString("PASS " + config.Password + "\r\n")
	writer.WriteString("NICK " + config.Nickname + "\r\n")
	writer.WriteString("USER " + config.Nickname + " 0 * :" + config.RealName + "\r\n")
	writer.WriteString("JOIN #global\r\n")
	writer.Flush()

	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Read error: ", err)
				return
			}
			fmt.Print("<<", line)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		writer.WriteString(text + "\r\n")
		writer.Flush()
	}
}
