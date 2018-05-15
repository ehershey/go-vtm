// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type NodesPerPoolNodeStatistics struct {
	Statistics struct {
		NewConn            *int    `json:"new_conn"`
		CurrentConn        *int    `json:"current_conn"`
		PktsFromNodeHi     *int    `json:"pkts_from_node_hi"`
		PooledConn         *int    `json:"pooled_conn"`
		PktsToNodeHi       *int    `json:"pkts_to_node_hi"`
		PktsToNode         *int    `json:"pkts_to_node"`
		ResponseMean       *int    `json:"response_mean"`
		ResponseMax        *int    `json:"response_max"`
		BytesToNodeHi      *int    `json:"bytes_to_node_hi"`
		BytesFromNode      *int    `json:"bytes_from_node"`
		L4StatelessBuckets *int    `json:"l4_stateless_buckets"`
		Errors             *int    `json:"errors"`
		BytesFromNodeLo    *int    `json:"bytes_from_node_lo"`
		CurrentRequests    *int    `json:"current_requests"`
		State              *string `json:"state"`
		BytesFromNodeHi    *int    `json:"bytes_from_node_hi"`
		PktsFromNode       *int    `json:"pkts_from_node"`
		PktsFromNodeLo     *int    `json:"pkts_from_node_lo"`
		NodePort           *int    `json:"node_port"`
		BytesToNodeLo      *int    `json:"bytes_to_node_lo"`
		BytesToNode        *int    `json:"bytes_to_node"`
		Failures           *int    `json:"failures"`
		ResponseMin        *int    `json:"response_min"`
		PktsToNodeLo       *int    `json:"pkts_to_node_lo"`
		IdleConns          *int    `json:"idle_conns"`
		TotalConn          *int    `json:"total_conn"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNodesPerPoolNodeStatistics(name string) (*NodesPerPoolNodeStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/nodes/per_pool_node/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(NodesPerPoolNodeStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
