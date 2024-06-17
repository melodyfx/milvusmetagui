package model

import (
	pb "milvusmetagui/proto/etcdpb"
)

type Partition struct {
	PartitionID               int64
	PartitionName             string
	PartitionCreatedTimestamp uint64
	Extra                     map[string]string // deprecated.
	CollectionID              int64
	State                     pb.PartitionState
}

func MarshalPartitionModel(partition *Partition) *pb.PartitionInfo {
	return &pb.PartitionInfo{
		PartitionID:               partition.PartitionID,
		PartitionName:             partition.PartitionName,
		PartitionCreatedTimestamp: partition.PartitionCreatedTimestamp,
		CollectionId:              partition.CollectionID,
		State:                     partition.State,
	}
}

func UnmarshalPartitionModel(info *pb.PartitionInfo) *Partition {
	return &Partition{
		PartitionID:               info.GetPartitionID(),
		PartitionName:             info.GetPartitionName(),
		PartitionCreatedTimestamp: info.GetPartitionCreatedTimestamp(),
		CollectionID:              info.GetCollectionId(),
		State:                     info.GetState(),
	}
}
