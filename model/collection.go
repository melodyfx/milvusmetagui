package model

import (
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	pb "milvusmetagui/proto/etcdpb"
)

type Collection struct {
	TenantID             string
	DBID                 int64
	CollectionID         int64
	Partitions           []*Partition
	Name                 string
	Description          string
	AutoID               bool
	Fields               []*Field
	VirtualChannelNames  []string
	PhysicalChannelNames []string
	ShardsNum            int32
	StartPositions       []*commonpb.KeyDataPair
	CreateTime           uint64
	ConsistencyLevel     commonpb.ConsistencyLevel
	Aliases              []string // TODO: deprecate this.
	Properties           []*commonpb.KeyValuePair
	State                pb.CollectionState
	EnableDynamicField   bool
}

func UnmarshalCollectionModel(coll *pb.CollectionInfo) *Collection {
	if coll == nil {
		return nil
	}

	// backward compatible for deprecated fields
	partitions := make([]*Partition, len(coll.PartitionIDs))
	for idx := range coll.PartitionIDs {
		partitions[idx] = &Partition{
			PartitionID:               coll.PartitionIDs[idx],
			PartitionName:             coll.PartitionNames[idx],
			PartitionCreatedTimestamp: coll.PartitionCreatedTimestamps[idx],
		}
	}

	return &Collection{
		CollectionID:         coll.ID,
		DBID:                 coll.DbId,
		Name:                 coll.Schema.Name,
		Description:          coll.Schema.Description,
		AutoID:               coll.Schema.AutoID,
		Fields:               UnmarshalFieldModels(coll.GetSchema().GetFields()),
		Partitions:           partitions,
		VirtualChannelNames:  coll.VirtualChannelNames,
		PhysicalChannelNames: coll.PhysicalChannelNames,
		ShardsNum:            coll.ShardsNum,
		ConsistencyLevel:     coll.ConsistencyLevel,
		CreateTime:           coll.CreateTime,
		StartPositions:       coll.StartPositions,
		State:                coll.State,
		Properties:           coll.Properties,
		EnableDynamicField:   coll.Schema.EnableDynamicField,
	}
}
