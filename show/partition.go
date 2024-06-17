package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	pb "milvusmetagui/proto/etcdpb"
)

func ShowPartsInfo(resp *clientv3.GetResponse) {
	parts, _ := ListPartitions(resp)
	PrintParts(parts)
}

func ListPartitions(resp *clientv3.GetResponse) (map[string]*model.Partition, error) {
	kvmap := make(map[string]string)
	for _, kv := range resp.Kvs {
		kvmap[string(kv.Key)] = string(kv.Value)
	}
	parts := make(map[string]*model.Partition)
	for k, v := range kvmap {
		partitionMeta := &pb.PartitionInfo{}
		err := proto.Unmarshal([]byte(v), partitionMeta)
		if err != nil {
			return nil, err
		}
		parts[k] = model.UnmarshalPartitionModel(partitionMeta)
	}
	return parts, nil
}

func PrintParts(partsMap map[string]*model.Partition) {
	for k, part := range partsMap {
		fmt.Printf("===key:%s===\n", k)
		fmt.Printf("PartitionID:%d\n", part.PartitionID)
		fmt.Printf("PartitionName:%s\n", part.PartitionName)
		fmt.Printf("PartitionCreatedTimestamp:%d\n", part.PartitionCreatedTimestamp)
		fmt.Printf("CollectionID:%d\n", part.CollectionID)
		fmt.Printf("State:%s\n", part.State.String())
		fmt.Println()
	}
}
