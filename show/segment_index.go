package show

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"milvusmetagui/model"
	"milvusmetagui/proto/indexpb"
)

func ShowSegIndexesInfo(resp *clientv3.GetResponse) {
	segindexes, _ := ListSegmentIndexes(resp)
	PrintSegIndexes(segindexes)
}

func ListSegmentIndexes(resp *clientv3.GetResponse) (map[string]*model.SegmentIndex, error) {
	kvmap := make(map[string]string)
	for _, kv := range resp.Kvs {
		kvmap[string(kv.Key)] = string(kv.Value)
	}
	segIndexes := make(map[string]*model.SegmentIndex)

	for k, v := range kvmap {
		segmentIndexInfo := &indexpb.SegmentIndex{}
		err := proto.Unmarshal([]byte(v), segmentIndexInfo)
		if err != nil {
			return segIndexes, err
		}

		segIndexes[k] = model.UnmarshalSegmentIndexModel(segmentIndexInfo)
	}

	return segIndexes, nil
}

func PrintSegIndexes(segIdxMap map[string]*model.SegmentIndex) {
	for k, segidx := range segIdxMap {
		fmt.Printf("===key:%s===\n", k)
		fmt.Printf("SegmentID:%d\n", segidx.SegmentID)
		fmt.Printf("CollectionID:%d\n", segidx.CollectionID)
		fmt.Printf("PartitionID:%d\n", segidx.PartitionID)
		fmt.Printf("NumRows:%d\n", segidx.NumRows)
		fmt.Printf("IndexID:%d\n", segidx.IndexID)
		fmt.Printf("BuildID:%d\n", segidx.BuildID)
		fmt.Printf("NodeID:%d\n", segidx.NodeID)
		fmt.Printf("IndexVersion:%d\n", segidx.IndexVersion)
		fmt.Printf("IndexState:%s\n", segidx.IndexState.String())
		fmt.Printf("FailReason:%s\n", segidx.FailReason)
		fmt.Printf("IsDeleted:%t\n", segidx.IsDeleted)
		fmt.Printf("CreateTime:%d\n", segidx.CreateTime)
		fmt.Printf("IndexSize:%d\n", segidx.IndexSize)
		fmt.Printf("WriteHandoff:%t\n", segidx.WriteHandoff)
		fmt.Printf("CurrentIndexVersion:%d\n", segidx.CurrentIndexVersion)
		fmt.Printf("IndexStoreVersion:%d\n", segidx.IndexStoreVersion)
		fmt.Printf("IndexFileKeys:%v\n", segidx.IndexFileKeys)
		fmt.Println()
	}
}
