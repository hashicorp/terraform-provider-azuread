package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAggregationFilter string

const (
	NetworkaccessAggregationFilter_BytesReceived NetworkaccessAggregationFilter = "bytesReceived"
	NetworkaccessAggregationFilter_BytesSent     NetworkaccessAggregationFilter = "bytesSent"
	NetworkaccessAggregationFilter_Devices       NetworkaccessAggregationFilter = "devices"
	NetworkaccessAggregationFilter_TotalBytes    NetworkaccessAggregationFilter = "totalBytes"
	NetworkaccessAggregationFilter_Transactions  NetworkaccessAggregationFilter = "transactions"
	NetworkaccessAggregationFilter_Users         NetworkaccessAggregationFilter = "users"
)

func PossibleValuesForNetworkaccessAggregationFilter() []string {
	return []string{
		string(NetworkaccessAggregationFilter_BytesReceived),
		string(NetworkaccessAggregationFilter_BytesSent),
		string(NetworkaccessAggregationFilter_Devices),
		string(NetworkaccessAggregationFilter_TotalBytes),
		string(NetworkaccessAggregationFilter_Transactions),
		string(NetworkaccessAggregationFilter_Users),
	}
}

func (s *NetworkaccessAggregationFilter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessAggregationFilter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessAggregationFilter(input string) (*NetworkaccessAggregationFilter, error) {
	vals := map[string]NetworkaccessAggregationFilter{
		"bytesreceived": NetworkaccessAggregationFilter_BytesReceived,
		"bytessent":     NetworkaccessAggregationFilter_BytesSent,
		"devices":       NetworkaccessAggregationFilter_Devices,
		"totalbytes":    NetworkaccessAggregationFilter_TotalBytes,
		"transactions":  NetworkaccessAggregationFilter_Transactions,
		"users":         NetworkaccessAggregationFilter_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessAggregationFilter(input)
	return &out, nil
}
