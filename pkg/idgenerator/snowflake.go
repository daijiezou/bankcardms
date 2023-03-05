package idgenerator

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var (
	node       *sf.Node
	outputNode *sf.Node
)

func Init() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", "2023-01-01")
	if err != nil {
		return
	}
	// 转为毫秒时间戳
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(1)
	sf.NodeBits = 3
	sf.StepBits = 5
	outputNode, err = sf.NewNode(1)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
func GenOutputId() int64 {
	return outputNode.Generate().Int64()
}

func GenIDString() string {
	return node.Generate().String()
}
