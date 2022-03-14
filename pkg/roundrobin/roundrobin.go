package roundrobin

import (
	"errors"
	"net/url"
	"sync/atomic"
)

var ErrorServerExists = errors.New("Server does not exist")

type RoundRobin interface {
	Next() *url.URL
}

type roundRobin struct {
	urls []*url.URL
	next uint32
}

//New instance of round robin
func New(urls []*url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, ErrorServerExists
	}

	return &roundRobin{
		urls: urls,
	}, nil
}

//Getting next URLs by Round Robin algo
func (r *roundRobin) Next() *url.URL {
	n := atomic.AddUint32(&r.next, 1)
	return r.urls[(int(n)-1)%len(r.urls)]
}
