package main

import (
"fmt"
"net"
"time"
"sync"
)

func main () {
var wg sync.WaitGroup
	for i:=1;i<=100;i++{
	wg.Add(1)
		go func(port int){
		defer wg.Done()
		target := fmt.Sprintf( "scanme.nmap.org:%d", port)
		conn, err:= net.DialTimeout("tcp", target, 2*time.Second)
		  if err==nil{
		  fmt.Println("Open port: ", port)
		conn.Close()
		  }
		}(i)
	}
	wg.Wait()
	fmt.Println("All ports are scnas")

}
