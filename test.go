package main

import (
	"fmt"

	mc "../minecraft-query"
)

func main() {
	fmt.Println("tester")

	m := mc.Minecraft{
		Address: "localhost",
		Port: 25565,
		Timeout: 10,
	}

	ml, err := m.ConnectUDP()
	if err != nil {
		fmt.Println(fmt.Sprintf("Connect: %v", err))
		return
	}

	f := ml.Query()
	fmt.Println(fmt.Sprintf("f: %v", f))
	ml.Disconnect()
}
