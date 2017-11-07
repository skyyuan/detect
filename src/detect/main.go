package main

import (
	_ "detect/routers"
	"github.com/astaxie/beego"
	"detect/utils"
	"flag"
	"fmt"
	"net"
	"encoding/binary"
	"time"
	"github.com/robfig/cron"
)

var (
	laddr = flag.String("Addr", "10.10.4.201:30003", "")
)
func main() {
	utils.InitRedisPool()
	defer utils.CloseRedisPool()
	utils.InitProducer()
	defer utils.CloseProducer()
	defer recoverAllPanic()
	initdetect()
	cronHeartBeat()
	beego.SetLogFuncCall(true)
	beego.Run()


}
func recoverAllPanic(){
	if err := recover(); err != nil {
		beego.Error("Recover for panic:",err)
	}
}

func cronHeartBeat(){
	spec := "*/10 * * * * *"
	c := cron.New()
	c.AddFunc(spec, DetectHEARTBEAT)
	c.Start()
}

func initdetect() {
	flag.Parse()
	localAddr, err := net.ResolveUDPAddr("udp", *laddr)

	if err != nil {
		fmt.Println("Can't resolve address: ", err)
	}

	conn, err := net.DialUDP("udp", nil, localAddr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(`{"name": "detect", "command": "NewSystemDetect" }`))
	if err != nil {
		fmt.Println("failed:", err)
	}
	data := make([]byte, 1024)

	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("Read of ", err)
	}
	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
}

func DetectHEARTBEAT()  {
	flag.Parse()
	localAddr, err := net.ResolveUDPAddr("udp", *laddr)

	if err != nil {
		fmt.Println("Can't resolve address: ", err)
	}

	conn, err := net.DialUDP("udp", nil, localAddr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(`{"name": "detect", "command": "DetectHEARTBEAT" }`))
	if err != nil {
		fmt.Println("failed:", err)
	}
	data := make([]byte, 1024)

	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("Read of ", err)
	}
	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
}