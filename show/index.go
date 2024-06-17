package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	"milvusmetagui/proto/indexpb"
)

func ShowIndexesInfo(resp *clientv3.GetResponse) {
	indexes, _ := ListIndexes(resp)
	PrintIndexes(indexes)
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

func PrintIndexes(indexMap map[string]*model.Index) {
	for k, index := range indexMap {
		fmt.Printf("===key:%s===\n", k)
		fmt.Printf("TenantID:%s\n", index.TenantID)
		fmt.Printf("CollectionID:%d\n", index.CollectionID)
		fmt.Printf("FieldID:%d\n", index.FieldID)
		fmt.Printf("IndexID:%d\n", index.IndexID)
		fmt.Printf("IndexName:%s\n", index.IndexName)
		fmt.Printf("IsDeleted:%t\n", index.IsDeleted)
		fmt.Printf("CreateTime:%d\n", index.CreateTime)
		fmt.Printf("IsAutoIndex:%t\n", index.IsAutoIndex)
		fmt.Printf("TypeParams:\n")
		for _, param := range index.TypeParams {
			fmt.Printf("%s:%s\n", param.GetKey(), param.GetValue())
		}
		fmt.Printf("IndexParams:\n")
		for _, param := range index.IndexParams {
			fmt.Printf("%s:%s\n", param.GetKey(), param.GetValue())
		}
		fmt.Printf("UserIndexParams:\n")
		for _, param := range index.UserIndexParams {
			fmt.Printf("%s:%s\n", param.GetKey(), param.GetValue())
		}
		fmt.Println()
	}
}
