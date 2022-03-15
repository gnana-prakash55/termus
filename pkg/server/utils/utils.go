package utils

import (
	"log"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/gnanaprakash55/termus/pkg/parsing"
	"github.com/gnanaprakash55/termus/pkg/roundrobin"
)

var urls []*url.URL

var mu sync.Mutex

var rr roundrobin.RoundRobin = Init_roundrobin()

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

func isAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute)
	if err != nil {
		log.Printf("Unreachable to %v, error: ", url.Host)
		return false
	}
	defer conn.Close()
	return true
}

func HealthCheck() {
	t := time.NewTicker(time.Minute * 1)

	for {
		select {
		case <-t.C:
			length := len(*rr.GetServers())
			backend := *rr.GetServers()
			for i := 0; i < length; i++ {
				pingURL := backend[i].URL
				isAlive := isAlive(pingURL)
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
