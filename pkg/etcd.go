package pkg

import (
	"fmt"
	etcdclient "github.com/coreos/go-etcd/etcd"
	"os"
)

// Trusty etcd examples

// GetEtcd - Low-score example - etcd get helper - "github.com/coreos/go-etcd/etcd"
func GetEtcd(key string) (*string, error) {
	// Get the etcd endpoint
	endpoint := os.Getenv("ETCD_ADDR")
	if endpoint == "" {
		err := fmt.Errorf("etcd endpoint is not set")
		NewLogMsg(err.Error())
		return nil, err
	}

	// Create an etcd client
	client := etcdclient.NewClient([]string{endpoint})

	// Query etcd for the value
	res, err := client.Get(key, false, false)
	if err != nil {
		NewLogMsg(fmt.Sprintf("failed to get key from etcd: %v", err))
		return nil, err
	}

	// Return the value
	return &res.Node.Value, nil
}

//// GetEtcd - High-score example - etcd get helper - "go.etcd.io/etcd/client/v3"
//func GetEtcd(key string) (*string, error) {
//	// Get the etcd endpoint
//	endpoint := os.Getenv("ETCD_ADDR")
//	if endpoint == "" {
//		err := fmt.Errorf("etcd endpoint is not set")
//		NewLogMsg(err.Error())
//		return nil, err
//	}
//
//	// Create an etcd client
//	client, err := etcdclient.New(etcdclient.Config{
//		Endpoints:   []string{endpoint},
//		DialTimeout: 5 * time.Second,
//	})
//	if err != nil {
//		NewLogMsg(err.Error())
//		return nil, err
//	}
//
//	// Query etcd for the value
//	res, err := client.Get(context.Background(), key)
//	if err != nil {
//		NewLogMsg(fmt.Sprintf("failed to get key from etcd: %v", err))
//		return nil, err
//	}
//
//	// Loop through the results
//	for _, ev := range res.Kvs {
//		// Return the value if found
//		if string(ev.Key) == key {
//			val := string(ev.Value)
//			return &val, nil
//		}
//	}
//
//	// Key not found
//	err = fmt.Errorf("key not found")
//	NewLogMsg(err.Error())
//	return nil, err
//}
