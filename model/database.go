package model

import (
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	pb "milvusmetagui/proto/etcdpb"
	"time"
)

type Database struct {
	TenantID    string
	ID          int64
	Name        string
	State       pb.DatabaseState
	CreatedTime uint64
	Properties  []*commonpb.KeyValuePair
}

func NewDatabase(id int64, name string, state pb.DatabaseState) *Database {
	return &Database{
		ID:          id,
		Name:        name,
		State:       state,
		CreatedTime: uint64(time.Now().UnixNano()),
		Properties:  make([]*commonpb.KeyValuePair, 0),
	}
}

func MarshalDatabaseModel(db *Database) *pb.DatabaseInfo {
	if db == nil {
		return nil
	}

	return &pb.DatabaseInfo{
		TenantId:    db.TenantID,
		Id:          db.ID,
		Name:        db.Name,
		State:       db.State,
		CreatedTime: db.CreatedTime,
		Properties:  db.Properties,
	}
}

func UnmarshalDatabaseModel(info *pb.DatabaseInfo) *Database {
	if info == nil {
		return nil
	}

	return &Database{
		Name:        info.GetName(),
		ID:          info.GetId(),
		CreatedTime: info.GetCreatedTime(),
		State:       info.GetState(),
		TenantID:    info.GetTenantId(),
		Properties:  info.GetProperties(),
	}
}
