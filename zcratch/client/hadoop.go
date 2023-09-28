package client

import (
	"fmt"

	"github.com/colinmarc/hdfs"
)

func NewHadoop(urls []string) (*hdfs.Client, error) {
	options := hdfs.ClientOptions{
		Addresses: urls,
	}

	client, err := hdfs.NewClient(options)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to hadoop: %v", err)
	}

	return client, nil
}
