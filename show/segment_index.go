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

func ShowSegIndexesInfo(resp *clientv3.GetResponse) string {
	segindexes, _ := ListSegmentIndexes(resp)
	return PrintSegIndexes(segindexes)
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

func PrintSegIndexes(segIdxMap map[string]*model.SegmentIndex) string {
	var builder strings.Builder
	for k, segidx := range segIdxMap {
		builder.WriteString(fmt.Sprintf("===key:%s===\n", k))
		builder.WriteString(fmt.Sprintf("SegmentID:%d\n", segidx.SegmentID))
		builder.WriteString(fmt.Sprintf("CollectionID:%d\n", segidx.CollectionID))
		builder.WriteString(fmt.Sprintf("PartitionID:%d\n", segidx.PartitionID))
		builder.WriteString(fmt.Sprintf("NumRows:%d\n", segidx.NumRows))
		builder.WriteString(fmt.Sprintf("IndexID:%d\n", segidx.IndexID))
		builder.WriteString(fmt.Sprintf("BuildID:%d\n", segidx.BuildID))
		builder.WriteString(fmt.Sprintf("NodeID:%d\n", segidx.NodeID))
		builder.WriteString(fmt.Sprintf("IndexVersion:%d\n", segidx.IndexVersion))
		builder.WriteString(fmt.Sprintf("IndexState:%s\n", segidx.IndexState.String()))
		builder.WriteString(fmt.Sprintf("FailReason:%s\n", segidx.FailReason))
		builder.WriteString(fmt.Sprintf("IsDeleted:%t\n", segidx.IsDeleted))
		p, l := utils.ParseTS(segidx.CreateTime)
		timestr := fmt.Sprintf("physicalTime:%s,logicalTime:%d", p, l)
		builder.WriteString(fmt.Sprintf("CreateTime:%d(%s)\n", segidx.CreateTime, timestr))
		builder.WriteString(fmt.Sprintf("IndexSize:%d\n", segidx.IndexSize))
		builder.WriteString(fmt.Sprintf("WriteHandoff:%t\n", segidx.WriteHandoff))
		builder.WriteString(fmt.Sprintf("CurrentIndexVersion:%d\n", segidx.CurrentIndexVersion))
		builder.WriteString(fmt.Sprintf("IndexStoreVersion:%d\n", segidx.IndexStoreVersion))
		builder.WriteString(fmt.Sprintf("IndexFileKeys:%v\n", segidx.IndexFileKeys))
		builder.WriteString("\n")
	}
	return builder.String()
}
