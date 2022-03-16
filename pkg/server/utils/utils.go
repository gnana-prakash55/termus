package utils

import (
	"log"
	"net"
	"net/url"
	"time"

	"github.com/gnanaprakash55/termus/pkg/parsing"
	"github.com/gnanaprakash55/termus/pkg/roundrobin"
)

var urls []*url.URL

//initialize round robin algo
func Init_roundrobin() roundrobin.RoundRobin {

	var servers = parsing.GetServers()

	for i := 0; i < len(servers); i++ {
		url, _ := url.Parse(servers[i].URL)
		urls = append(urls, url)
	}

	rr, err := roundrobin.New(urls)
	if err != nil {
		panic(err)
	}

	return rr
}

func IsAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute)
	if err != nil {
		log.Printf("Unreachable to %v, error: %v", url.Host, err.Error())
		return false
	}
	defer conn.Close()
	return true
}

func HealthCheck() {
	var rr roundrobin.RoundRobin = Init_roundrobin()

	t := time.NewTicker(time.Minute * 1)

	for {
		select {
		case <-t.C:
			length := len(*rr.GetServers())
			backend := *rr.GetServers()
			for i := 0; i < length; i++ {
				pingURL := backend[i].URL
				isAlive := IsAlive(pingURL)
				backend[i].SetDead(!isAlive)
				msg := "ok"
				if !isAlive {
					msg = "dead"
				}
				log.Printf("%v checked %v by healthcheck", backend[i].URL, msg)
			}
		}
	}
}

func BadGateway() string {

	return `
	<html>
	<head>
	<style> 
	h1 { text-align: center; }
	h3 { text-align: center; }
	</style>
	</head>
	<body>

	<h1><strong>502 Bad Gateway</strong></h1>
	<hr>
	<h3>termus</h3>

	</body>
	<html>
	`

}

func Successful() string {

	return `
	<html>
	<head>
	<style> 
	h1 { text-align: center; }
	h3 { text-align: center; }
	</style>
	</head>
	<body>

	<h1><strong>termus</strong></h1>
	<hr>
	<h3>Successfully Installed</h3>

	</body>
	<html>
	`

}
