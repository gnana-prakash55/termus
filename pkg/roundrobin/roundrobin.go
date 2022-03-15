package roundrobin

import (
	"errors"
	"net/url"
	"sync"
	"sync/atomic"
)

var ErrorServerExists = errors.New("Server does not exist")

type RoundRobin interface {
	Next() *Server
	AddServers(server *url.URL)
	GetServers() *[]Server
}

type roundRobin struct {
	servers []Server
	next    uint32
}

type Server struct {
	URL    *url.URL
	IsDead bool
	mu     sync.RWMutex
}

//New instance of round robin
func New(urls []*url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, ErrorServerExists
	}

	var rr RoundRobin

	rr = &roundRobin{}

	for i := 0; i < len(urls); i++ {
		rr.AddServers(urls[i])
	}

	return rr, nil
}

//Getting next URLs by Round Robin algo
func (r *roundRobin) Next() *Server {
	n := atomic.AddUint32(&r.next, 1)
	return &r.servers[(int(n)-1)%len(r.servers)]
}

//Adding Servers
func (r *roundRobin) AddServers(url *url.URL) {
	r.servers = append(r.servers, Server{URL: url})
}

//Getting Servers
func (r *roundRobin) GetServers() *[]Server {
	return &r.servers
}

//Set Dead updates the value of IsDead
func (s *Server) SetDead(d bool) {
	s.mu.Lock()
	s.IsDead = d
	s.mu.Unlock()
}

//GetIsDead read the value of IsDead
func (s *Server) GetIsDead() bool {
	s.mu.RLock()
	isAlive := s.IsDead
	s.mu.RUnlock()
	return isAlive
}
