// Copyright (c) 2021 PaddlePaddle Authors. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"math"

	"github.com/PaddlePaddle/PaddleDTX/dai/blockchain"
)

const (
	prefixFlTaskIndex       = "index_fltask"
	prefixFlTaskListIndex   = "index_fltask_list"
	prefixREFlTaskListIndex = "index_re_fltask_list"
	prefixNodeIndex         = "index_executor_node"
	prefixNodeNameIndex     = "index_executor_name"
	prefixNodeListIndex     = "index_executor_node_list"
)

// subByInt64Max return maxInt64 - N
func subByInt64Max(n int64) int64 {
	return math.MaxInt64 - n
}

// packFlTaskIndex pack task index for saving task using taskID
func packFlTaskIndex(taskID string) string {
	return fmt.Sprintf("%s/%s", prefixFlTaskIndex, taskID)
}

// packFlTaskListIndex pack index for saving tasks for requester, in time descending order
func packFlTaskListIndex(task blockchain.FLTask) string {
	return fmt.Sprintf("%s/%x/%d/%s", prefixFlTaskListIndex, task.Requester, subByInt64Max(task.PublishTime), task.TaskID)
}

// packExecutorTaskListIndex pack index for saving tasks that an executor involves, in time descending order
func packExecutorTaskListIndex(executor []byte, task blockchain.FLTask) string {
	return fmt.Sprintf("%s/%x/%d/%s", prefixFlTaskListIndex, executor, subByInt64Max(task.PublishTime), task.TaskID)
}

func packRequesterExecutorTaskIndex(executor []byte, task blockchain.FLTask) string {
	return fmt.Sprintf("%s/%x/%x/%d/%s", prefixREFlTaskListIndex, task.Requester, executor, subByInt64Max(task.PublishTime), task.TaskID)
}

// packFlTaskFilter pack filter index with public key for searching tasks for requester or executor
func packFlTaskFilter(rPubkey, ePubkey []byte) (string, []string) {
	if len(rPubkey) > 0 && len(ePubkey) > 0 {
		return prefixFlTaskListIndex, []string{fmt.Sprintf("%x", rPubkey), fmt.Sprintf("%x", ePubkey)}
	} else if len(rPubkey) > 0 {
		return prefixFlTaskListIndex, []string{fmt.Sprintf("%x", rPubkey)}
	} else {
		return prefixFlTaskListIndex, []string{fmt.Sprintf("%x", ePubkey)}
	}
}

// packNodeIndex pack index-id contract key for saving executor node
func packNodeIndex(nodeID []byte) string {
	return fmt.Sprintf("%s/%x", prefixNodeIndex, nodeID)
}

// packNodeStringIndex pack index-id contract key for saving executor node
func packNodeStringIndex(nodeID string) string {
	return fmt.Sprintf("%s/%s", prefixNodeIndex, nodeID)
}

// packNodeNameIndex pack index-name contract key for saving executor node
func packNodeNameIndex(name string) string {
	return fmt.Sprintf("%s/%s", prefixNodeNameIndex, name)
}

// packNodeListIndex pack filter for listing executor nodes
func packNodeListIndex(node blockchain.ExecutorNode) string {
	return fmt.Sprintf("%s/%d/%x", prefixNodeListIndex, subByInt64Max(node.RegTime), node.ID)
}
