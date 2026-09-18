package main

import(
"fmt"
"log"

"github.com/google/gopacket"
"github.com/google/gopacket/pcap"

)

func main(){

deviceName:="wlan0"

fmt.Println("Premission mode %s", deviceName)

handle, err:= pcap.OpenLive(deviceName,1600,true, pcap.BlockForever)
	if err!=nil{
	log.Fatal("Problem ", err)
	
	}
defer handle.Close()

packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets(){
		fmt.Println(packet)
	}
}
