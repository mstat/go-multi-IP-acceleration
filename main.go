package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"math/rand"
	"net"
	"sort"
)

type configType struct {
	port      string
	localIp   []string
	allowedIp map[string]bool
}

var config configType

var p = fmt.Println

func Fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	ini, err := goconfig.LoadConfigFile("goproxy.ini")
	Fatal(err)
	config.port, _ = ini.GetValue("port", "port_socks5")
	localIp, _ := ini.GetSection("local ip")
	allowedIp, _ := ini.GetSection("allowed ip")

	for _, value := range localIp {
		config.localIp = append(config.localIp, value)
	}
	config.allowedIp = make(map[string]bool)
	for _, value := range allowedIp {
		config.allowedIp[value] = true
	}
	sort.Strings(config.localIp)
	p("Port :", config.port)
	p("Local IP :")
	for _, ip := range config.localIp {
		p("\t" + ip)
	}
	p("Allowed IP :")
	for ip, _ := range config.allowedIp {
		p("\t" + ip)
	}
	if len(config.localIp) == 0 {
		log.Fatalln("The number of localIP must be greater than 0")
	}
}

func randLocalAddr() (r *net.TCPAddr) {
	ip := config.localIp[rand.Int()%len(config.localIp)]
	r, err := net.ResolveTCPAddr("tcp4", ip+":0")
	Fatal(err)
	p("********local ip", r)
	return
}

func main() {
	conf := &Config{}
	server, err := New(conf)
	Fatal(err)
	p("proxy start")
	err = server.ListenAndServe("tcp", "0.0.0.0:"+config.port)
	Fatal(err)
}
