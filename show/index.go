package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	"milvusmetagui/proto/indexpb"
	"milvusmetagui/utils"
	"strings"
)

func ShowIndexesInfo(resp *clientv3.GetResponse) string {
	indexes, _ := ListIndexes(resp)
	return PrintIndexes(indexes)
}

func ListIndexes(resp *clientv3.GetResponse) (map[string]*model.Index, error) {
	kvmap := make(map[string]string)
	for _, kv := range resp.Kvs {
		kvmap[string(kv.Key)] = string(kv.Value)
	}
	indexes := make(map[string]*model.Index)
	for k, v := range kvmap {
		meta := &indexpb.FieldIndex{}
		err := proto.Unmarshal([]byte(v), meta)
		if err != nil {
			return nil, err
		}

		indexes[k] = model.UnmarshalIndexModel(meta)
	}

	return indexes, nil
}

func PrintIndexes(indexMap map[string]*model.Index) string {
	var builder strings.Builder
	for k, index := range indexMap {
		builder.WriteString(fmt.Sprintf("===key:%s===\n", k))
		builder.WriteString(fmt.Sprintf("TenantID:%s\n", index.TenantID))
		builder.WriteString(fmt.Sprintf("CollectionID:%d\n", index.CollectionID))
		builder.WriteString(fmt.Sprintf("FieldID:%d\n", index.FieldID))
		builder.WriteString(fmt.Sprintf("IndexID:%d\n", index.IndexID))
		builder.WriteString(fmt.Sprintf("IndexName:%s\n", index.IndexName))
		builder.WriteString(fmt.Sprintf("IsDeleted:%t\n", index.IsDeleted))
		p, l := utils.ParseTS(index.CreateTime)
		timestr := fmt.Sprintf("physicalTime:%s,logicalTime:%d", p, l)
		builder.WriteString(fmt.Sprintf("CreateTime:%d(%s)\n", index.CreateTime, timestr))
		builder.WriteString(fmt.Sprintf("IsAutoIndex:%t\n", index.IsAutoIndex))
		builder.WriteString(fmt.Sprintf("TypeParams:\n"))
		for _, param := range index.TypeParams {
			builder.WriteString(fmt.Sprintf("    %s:%s\n", param.GetKey(), param.GetValue()))
		}
		builder.WriteString(fmt.Sprintf("IndexParams:\n"))
		for _, param := range index.IndexParams {
			builder.WriteString(fmt.Sprintf("    %s:%s\n", param.GetKey(), param.GetValue()))
		}
		builder.WriteString(fmt.Sprintf("UserIndexParams:\n"))
		for _, param := range index.UserIndexParams {
			builder.WriteString(fmt.Sprintf("    %s:%s\n", param.GetKey(), param.GetValue()))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
