package client

import (
	"fmt"

	"github.com/ceph/go-ceph/rados"
)

func NewCeph(url string) (*rados.Conn, error) {
	conn, _ := rados.NewConn()

	if err := conn.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to ceph: %v", err)
	}

	return conn, nil
}
