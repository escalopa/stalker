package client

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func NewMemcache(urls []string) (*memcache.Client, error) {
	if len(urls) == 0 {
		return nil, fmt.Errorf("no memcache urls provided")
	}

	client := memcache.New(urls...)

	if err := client.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping memcache: %v", err)
	}

	return client, nil
}
