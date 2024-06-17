package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	pb "milvusmetagui/proto/etcdpb"
	"strings"
)

func ShowCollsInfo(resp *clientv3.GetResponse) string {
	colls, _ := ListCollections(resp)
	return PrintColls(colls)
}

func ListCollections(resp *clientv3.GetResponse) (map[string]*model.Collection, error) {
	kvmap := make(map[string]string)
	for _, kv := range resp.Kvs {
		kvmap[string(kv.Key)] = string(kv.Value)
	}

	colls := make(map[string]*model.Collection)
	for k, val := range kvmap {
		collMeta := &pb.CollectionInfo{}
		proto.Unmarshal([]byte(val), collMeta)
		colls[k] = model.UnmarshalCollectionModel(collMeta)
	}
	return colls, nil
}

func PrintColls(colls map[string]*model.Collection) string {
	var builder strings.Builder
	for k, coll := range colls {
		builder.WriteString(fmt.Sprintf("===key:%s===\n", k))
		builder.WriteString(fmt.Sprintf("TenantID:%s\n", coll.TenantID))
		builder.WriteString(fmt.Sprintf("DBID:%d\n", coll.DBID))
		builder.WriteString(fmt.Sprintf("CollectionID:%d\n", coll.CollectionID))
		builder.WriteString(fmt.Sprintf("Partitions:%s\n", ""))
		builder.WriteString(fmt.Sprintf("Name:%s\n", coll.Name))
		builder.WriteString(fmt.Sprintf("Description:%s\n", coll.Description))
		builder.WriteString(fmt.Sprintf("AutoID:%t\n", coll.AutoID))
		builder.WriteString(fmt.Sprintf("Fields:%s\n", ""))
		builder.WriteString(fmt.Sprintf("VirtualChannelNames:%v\n", coll.VirtualChannelNames))
		builder.WriteString(fmt.Sprintf("PhysicalChannelNames:%v\n", coll.PhysicalChannelNames))
		builder.WriteString(fmt.Sprintf("ShardsNum:%d\n", coll.ShardsNum))
		builder.WriteString(fmt.Sprintf("CreateTime:%d\n", coll.CreateTime))
		builder.WriteString(fmt.Sprintf("ConsistencyLevel:%s\n", coll.ConsistencyLevel.String()))
		builder.WriteString(fmt.Sprintf("Aliases:%v\n", coll.Aliases))
		builder.WriteString(fmt.Sprintf("State:%s\n", coll.State.String()))
		builder.WriteString(fmt.Sprintf("EnableDynamicField:%t\n", coll.EnableDynamicField))
		builder.WriteString(fmt.Sprintf("StartPositions:\n"))
		for _, position := range coll.StartPositions {
			builder.WriteString(fmt.Sprintf("    %s:0x%02x\n", position.GetKey(), position.GetData()))
		}
		builder.WriteString(fmt.Sprintf("Properties:\n"))
		for _, property := range coll.Properties {
			builder.WriteString(fmt.Sprintf("    %s:%s\n", property.GetKey(), property.GetValue()))
		}
	}
	return builder.String()
}
