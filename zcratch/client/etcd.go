package client

import (
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcd(urls []string) (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   urls,
		DialTimeout: DefaultTimeout,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to etcd: %v", err)
	}

	return client, nil
}
