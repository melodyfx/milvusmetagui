package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	"strings"
)

func ShowFieldsInfo(resp *clientv3.GetResponse) string {
	fields, _ := listFields(resp)
	return PrintFields(fields)
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

func PrintFields(fieldsMap map[string]*model.Field) string {
	var builder strings.Builder
	for k, field := range fieldsMap {
		builder.WriteString(fmt.Sprintf("===key:%s===\n", k))
		builder.WriteString(fmt.Sprintf("FieldID:%d\n", field.FieldID))
		builder.WriteString(fmt.Sprintf("Name:%s\n", field.Name))
		builder.WriteString(fmt.Sprintf("IsPrimaryKey:%t\n", field.IsPrimaryKey))
		builder.WriteString(fmt.Sprintf("Description:%s\n", field.Description))
		builder.WriteString(fmt.Sprintf("DataType:%s\n", field.DataType.String()))
		builder.WriteString(fmt.Sprintf("AutoID:%t\n", field.AutoID))
		builder.WriteString(fmt.Sprintf("State:%s\n", field.State.String()))
		builder.WriteString(fmt.Sprintf("IsDynamic:%t\n", field.IsDynamic))
		builder.WriteString(fmt.Sprintf("IsPartitionKey:%t\n", field.IsPartitionKey))
		builder.WriteString(fmt.Sprintf("IsClusteringKey:%t\n", field.IsClusteringKey))
		builder.WriteString(fmt.Sprintf("ElementType:%s\n", field.ElementType.String()))
		if field.DefaultValue == nil {
			builder.WriteString(fmt.Sprintf("DefaultValue:%s\n", ""))
		} else {
			builder.WriteString(fmt.Sprintf("DefaultValue:%s\n", "待解析"))
		}
		builder.WriteString(fmt.Sprintf("TypeParams:\n"))
		for _, param := range field.TypeParams {
			builder.WriteString(fmt.Sprintf("    %s:%s\n", param.GetKey(), param.GetValue()))
		}
		builder.WriteString(fmt.Sprintf("IndexParams:\n"))
		for _, param := range field.IndexParams {
			builder.WriteString(fmt.Sprintf("    %s:%s\n", param.GetKey(), param.GetValue()))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
