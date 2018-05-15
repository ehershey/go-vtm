// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type ServiceLevelMonitorStatistics struct {
	Statistics struct {
		ResponseMean *int    `json:"response_mean"`
		CurrentConns *int    `json:"current_conns"`
		TotalConn    *int    `json:"total_conn"`
		Conforming   *int    `json:"conforming"`
		TotalNonConf *int    `json:"total_non_conf"`
		ResponseMin  *int    `json:"response_min"`
		ResponseMax  *int    `json:"response_max"`
		IsOK         *string `json:"is_o_k"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetServiceLevelMonitorStatistics(name string) (*ServiceLevelMonitorStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/service_level_monitors/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(ServiceLevelMonitorStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
