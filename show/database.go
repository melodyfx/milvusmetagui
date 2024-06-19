package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	pb "milvusmetagui/proto/etcdpb"
	"strings"
	"time"
)

func ShowDbsInfo(resp *clientv3.GetResponse) string {
	dbs, _ := ListDatabases(resp)
	return PrintDbs(dbs)
}

func ListDatabases(resp *clientv3.GetResponse) (map[string]*model.Database, error) {
	kvmap := make(map[string]string)
	for _, kv := range resp.Kvs {
		kvmap[string(kv.Key)] = string(kv.Value)
	}
	dbs := make(map[string]*model.Database)
	for k, val := range kvmap {
		dbMeta := &pb.DatabaseInfo{}
		proto.Unmarshal([]byte(val), dbMeta)
		dbs[k] = model.UnmarshalDatabaseModel(dbMeta)
	}
	return dbs, nil
}

func PrintDbs(dbs map[string]*model.Database) string {
	var builder strings.Builder
	for k, db := range dbs {
		builder.WriteString(fmt.Sprintf("===key:%s===\n", k))
		builder.WriteString(fmt.Sprintf("TenantID:%s\n", db.TenantID))
		builder.WriteString(fmt.Sprintf("ID:%d\n", db.ID))
		builder.WriteString(fmt.Sprintf("Name:%s\n", db.Name))
		builder.WriteString(fmt.Sprintf("State:%s\n", db.State.String()))
		t := time.Unix(0, int64(db.CreatedTime))
		formatted := t.Format("2006-01-02 15:04:05")
		builder.WriteString(fmt.Sprintf("CreatedTime:%d(%s)\n", db.CreatedTime, formatted))
		for _, props := range db.Properties {
			builder.WriteString(fmt.Sprintf("%s:%s\n", props.GetKey(), props.GetValue()))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
