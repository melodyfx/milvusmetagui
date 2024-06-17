package show

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestShowPartsInfo(t *testing.T) {
	c, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.230.71:2379"},
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(),
		},
	})
	keyPrefix := "by-dev/meta/root-coord/partitions/449657611215201437"
	opts := []clientv3.OpOption{
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
		clientv3.WithLimit(5000),
		clientv3.WithRange(clientv3.GetPrefixRangeEnd(keyPrefix)),
	}
	ctx := context.Background()
	resp, _ := c.Get(ctx, keyPrefix, opts...)
	ShowPartsInfo(resp)
}
