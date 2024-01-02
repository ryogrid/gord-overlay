package chord

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/taisho6339/gord/pkg/model"
	"github.com/taisho6339/gord/pkg/test"
	"testing"
	"time"
)

var mockTransport = &MockTransport{}

func TestProcess_SingleNode(t *testing.T) {
	assert.NotPanics(t, func() {
		ctx := context.Background()
		hostName := "single"
		node := NewLocalNode(hostName)
		process := NewProcess(node, mockTransport)
		assert.NoError(t, process.Start(context.Background()))

		succ, err := process.FindSuccessorByTable(ctx, model.NewHashID(hostName))
		assert.Nil(t, err)
		assert.Equal(t, hostName, succ.Reference().Host)
	})
}

//func TestProcess_MultiNodes(t *testing.T) {
//	ctx := context.Background()
//	processes := waitGenerateProcesses(ctx, 3)
//	process1, process2, process3 := processes[0], processes[1], processes[2]
//	defer process1.Shutdown()
//	defer process2.Shutdown()
//	defer process3.Shutdown()
//	var (
//		node1Name = "gord1"
//		node2Name = "gord2"
//		node3Name = "gord3"
//	)
//	testcases := []struct {
//		findingID      model.HashID
//		expectedHost   string
//		callingProcess *Process
//	}{
//		{
//			findingID:      model.BytesToHashID(big.NewInt(1).Bytes()),
//			expectedHost:   node1Name,
//			callingProcess: process1,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(1).Bytes()),
//			expectedHost:   node1Name,
//			callingProcess: process2,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(1).Bytes()),
//			expectedHost:   node1Name,
//			callingProcess: process3,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(2).Bytes()),
//			expectedHost:   node2Name,
//			callingProcess: process1,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(2).Bytes()),
//			expectedHost:   node2Name,
//			callingProcess: process2,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(2).Bytes()),
//			expectedHost:   node2Name,
//			callingProcess: process3,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(3).Bytes()),
//			expectedHost:   node3Name,
//			callingProcess: process1,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(3).Bytes()),
//			expectedHost:   node3Name,
//			callingProcess: process2,
//		},
//		{
//			findingID:      model.BytesToHashID(big.NewInt(3).Bytes()),
//			expectedHost:   node3Name,
//			callingProcess: process3,
//		},
//	}
//	for _, testcase := range testcases {
//		assert.NotPanics(t, func() {
//			t.Logf("[CASE] finding: %x, expected: %s, call node: %s", testcase.findingID, testcase.expectedHost, testcase.callingProcess.Host)
//			succ, err := testcase.callingProcess.FindSuccessorByTable(ctx, testcase.findingID)
//			assert.Nil(t, err)
//			assert.Equal(t, testcase.expectedHost, succ.Reference().Host)
//		})
//	}
//}

func TestProcess_Stabilize_SuccessorList(t *testing.T) {
	ctx := context.Background()
	processes := waitGenerateProcesses(ctx, 3)
	process1, process2, process3 := processes[0], processes[1], processes[2]
	defer process1.Shutdown()
	defer process2.Shutdown()
	defer process3.Shutdown()
	testcases := []struct {
		targetNode            RingNode
		expectedSuccessorList []RingNode
	}{
		{
			targetNode: process1.LocalNode,
			expectedSuccessorList: []RingNode{
				process2.LocalNode,
				process3.LocalNode,
				process1.LocalNode,
			},
		},
		{
			targetNode: process2.LocalNode,
			expectedSuccessorList: []RingNode{
				process3.LocalNode,
				process1.LocalNode,
			},
		},
		{
			targetNode: process3.LocalNode,
			expectedSuccessorList: []RingNode{
				process1.LocalNode,
				process2.LocalNode,
			},
		},
	}
	for _, testcase := range testcases {
		assert.NotPanics(t, func() {
			ctx := context.Background()
			successors, err := testcase.targetNode.GetSuccessors(ctx)
			assert.Nil(t, err)
			for i, suc := range testcase.expectedSuccessorList {
				assert.Equal(t, suc.Reference().ID, successors[i].Reference().ID)
			}
		})
	}
}

func TestProcess_Node_Failure(t *testing.T) {
	ctx := context.Background()
	assert.NotPanics(t, func() {
		processes := waitGenerateProcesses(ctx, 3)
		process1, process2, process3 := processes[0], processes[1], processes[2]
		process1.Shutdown()
		test.WaitCheckFuncWithTimeout(func() {
			t.Fatal("test failed by timeout.")
		}, func() bool {
			return len(process2.successors.nodes) == 2 && len(process3.successors.nodes) == 2
		}, 10*time.Second)

		assert.Equal(t, 2, len(process2.successors.nodes))
		assert.Equal(t, 2, len(process3.successors.nodes))
		for _, s := range process2.successors.nodes {
			assert.NotEqual(t, process1.ID, s.Reference().ID)
		}

		process2.Shutdown()
		test.WaitCheckFuncWithTimeout(func() {
			t.Fatal("test failed by timeout.")
		}, func() bool {
			return len(process3.successors.nodes) == 1
		}, 10*time.Second)

		suc, err := process3.successors.head()
		assert.Nil(t, err)
		assert.Equal(t, process3.ID, suc.Reference().ID)
	})
}
