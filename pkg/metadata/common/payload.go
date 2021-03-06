// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

package common

// Payload handles the JSON unmarshalling of the metadata payload
type Payload struct {
	APIKey           string `json:"apiKey"`
	AgentVersion     string `json:"agentVersion"`
	UUID             string `json:"uuid"`
	InternalHostname string `json:"internalHostname"`
}
