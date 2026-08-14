package main

import (
	"fmt"
	"net"
	"time"
)
func main(){
	for port:=78;port<=80; port++{
	target := fmt.Sprintf("scanme.nmap.org:%d", port)
	conn, err:=net.DialTimeout("tcp", target, 2*time.Second)
		if err == nil {
		fmt.Println("port is open", port)
		conn.Close()
		}
	}

}
