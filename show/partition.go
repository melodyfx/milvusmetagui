package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	pb "milvusmetagui/proto/etcdpb"
	"milvusmetagui/utils"
	"strings"
)

func ShowPartsInfo(resp *clientv3.GetResponse) string {
	parts, _ := ListPartitions(resp)
	return PrintParts(parts)
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

func PrintParts(partsMap map[string]*model.Partition) string {
	var builder strings.Builder
	for k, part := range partsMap {
		builder.WriteString(fmt.Sprintf("===key:%s===\n", k))
		builder.WriteString(fmt.Sprintf("PartitionID:%d\n", part.PartitionID))
		builder.WriteString(fmt.Sprintf("PartitionName:%s\n", part.PartitionName))
		p, l := utils.ParseTS(part.PartitionCreatedTimestamp)
		timestr := fmt.Sprintf("physicalTime:%s,logicalTime:%d", p, l)
		builder.WriteString(fmt.Sprintf("PartitionCreatedTimestamp:%d(%s)\n", part.PartitionCreatedTimestamp, timestr))
		builder.WriteString(fmt.Sprintf("CollectionID:%d\n", part.CollectionID))
		builder.WriteString(fmt.Sprintf("State:%s\n", part.State.String()))
		builder.WriteString("\n")
	}
	return builder.String()
}
