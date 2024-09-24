package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBandwidthCapacityInMbps string

const (
	NetworkaccessBandwidthCapacityInMbps_Mbps1000 NetworkaccessBandwidthCapacityInMbps = "mbps1000"
	NetworkaccessBandwidthCapacityInMbps_Mbps250  NetworkaccessBandwidthCapacityInMbps = "mbps250"
	NetworkaccessBandwidthCapacityInMbps_Mbps500  NetworkaccessBandwidthCapacityInMbps = "mbps500"
	NetworkaccessBandwidthCapacityInMbps_Mbps750  NetworkaccessBandwidthCapacityInMbps = "mbps750"
)

func PossibleValuesForNetworkaccessBandwidthCapacityInMbps() []string {
	return []string{
		string(NetworkaccessBandwidthCapacityInMbps_Mbps1000),
		string(NetworkaccessBandwidthCapacityInMbps_Mbps250),
		string(NetworkaccessBandwidthCapacityInMbps_Mbps500),
		string(NetworkaccessBandwidthCapacityInMbps_Mbps750),
	}
}

func (s *NetworkaccessBandwidthCapacityInMbps) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessBandwidthCapacityInMbps(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessBandwidthCapacityInMbps(input string) (*NetworkaccessBandwidthCapacityInMbps, error) {
	vals := map[string]NetworkaccessBandwidthCapacityInMbps{
		"mbps1000": NetworkaccessBandwidthCapacityInMbps_Mbps1000,
		"mbps250":  NetworkaccessBandwidthCapacityInMbps_Mbps250,
		"mbps500":  NetworkaccessBandwidthCapacityInMbps_Mbps500,
		"mbps750":  NetworkaccessBandwidthCapacityInMbps_Mbps750,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessBandwidthCapacityInMbps(input)
	return &out, nil
}
