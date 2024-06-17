package model

import (
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"milvusmetagui/proto/indexpb"
)

type Index struct {
	TenantID        string
	CollectionID    int64
	FieldID         int64
	IndexID         int64
	IndexName       string
	IsDeleted       bool
	CreateTime      uint64
	TypeParams      []*commonpb.KeyValuePair
	IndexParams     []*commonpb.KeyValuePair
	IsAutoIndex     bool
	UserIndexParams []*commonpb.KeyValuePair
}

func UnmarshalIndexModel(indexInfo *indexpb.FieldIndex) *Index {
	if indexInfo == nil {
		return nil
	}

	return &Index{
		CollectionID:    indexInfo.IndexInfo.GetCollectionID(),
		FieldID:         indexInfo.IndexInfo.GetFieldID(),
		IndexID:         indexInfo.IndexInfo.GetIndexID(),
		IndexName:       indexInfo.IndexInfo.GetIndexName(),
		IsDeleted:       indexInfo.GetDeleted(),
		CreateTime:      indexInfo.CreateTime,
		TypeParams:      indexInfo.IndexInfo.GetTypeParams(),
		IndexParams:     indexInfo.IndexInfo.GetIndexParams(),
		IsAutoIndex:     indexInfo.IndexInfo.GetIsAutoIndex(),
		UserIndexParams: indexInfo.IndexInfo.GetUserIndexParams(),
	}
}

func MarshalIndexModel(index *Index) *indexpb.FieldIndex {
	if index == nil {
		return nil
	}

	return &indexpb.FieldIndex{
		IndexInfo: &indexpb.IndexInfo{
			CollectionID:    index.CollectionID,
			FieldID:         index.FieldID,
			IndexName:       index.IndexName,
			IndexID:         index.IndexID,
			TypeParams:      index.TypeParams,
			IndexParams:     index.IndexParams,
			IsAutoIndex:     index.IsAutoIndex,
			UserIndexParams: index.UserIndexParams,
		},
		Deleted:    index.IsDeleted,
		CreateTime: index.CreateTime,
	}
}
