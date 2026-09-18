package main

import(
"fmt"
"log"

"github.com/google/gopacket"
"github.com/google/gopacket/layers"
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
		ipLayer := packet.Layer(layers.LayerTypeIPv4)

		
		if ipLayer!= nil{
			ip, _ := ipLayer.(*layers.IPv4)
			fmt.Println("[IP] the person send %s -> Get %s\n", ip.SrcIP,ip.DstIP)
		}
	}
}
