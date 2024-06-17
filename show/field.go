package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
)

func ShowFieldsInfo(resp *clientv3.GetResponse) {
	fields, _ := listFields(resp)
	PrintFields(fields)
}

func listFields(resp *clientv3.GetResponse) (map[string]*model.Field, error) {
	kvmap := make(map[string]string)
	for _, kv := range resp.Kvs {
		kvmap[string(kv.Key)] = string(kv.Value)
	}

	fields := make(map[string]*model.Field)
	for k, v := range kvmap {
		partitionMeta := &schemapb.FieldSchema{}
		err := proto.Unmarshal([]byte(v), partitionMeta)
		if err != nil {
			return nil, err
		}
		fields[k] = model.UnmarshalFieldModel(partitionMeta)
	}
	return fields, nil
}

func PrintFields(fieldsMap map[string]*model.Field) {
	for k, field := range fieldsMap {
		fmt.Printf("===key:%s===\n", k)
		fmt.Printf("FieldID:%d\n", field.FieldID)
		fmt.Printf("Name:%s\n", field.Name)
		fmt.Printf("IsPrimaryKey:%t\n", field.IsPrimaryKey)
		fmt.Printf("Description:%s\n", field.Description)
		fmt.Printf("DataType:%s\n", field.DataType.String())
		fmt.Printf("AutoID:%t\n", field.AutoID)
		fmt.Printf("State:%s\n", field.State.String())
		fmt.Printf("IsDynamic:%t\n", field.IsDynamic)
		fmt.Printf("IsPartitionKey:%t\n", field.IsPartitionKey)
		fmt.Printf("IsClusteringKey:%t\n", field.IsClusteringKey)
		fmt.Printf("ElementType:%s\n", field.ElementType.String())
		if field.DefaultValue == nil {
			fmt.Printf("DefaultValue:%s\n", "")
		} else {
			fmt.Printf("DefaultValue:%s\n", "待解析")
		}
		fmt.Printf("TypeParams:\n")
		for _, param := range field.TypeParams {
			fmt.Printf("%s:%s\n", param.GetKey(), param.GetValue())
		}
		fmt.Printf("IndexParams:\n")
		for _, param := range field.IndexParams {
			fmt.Printf("%s:%s\n", param.GetKey(), param.GetValue())
		}
		fmt.Println()
	}
}
