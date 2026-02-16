package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnDeadPeerDetectionRate string

const (
	VpnDeadPeerDetectionRate_High   VpnDeadPeerDetectionRate = "high"
	VpnDeadPeerDetectionRate_Low    VpnDeadPeerDetectionRate = "low"
	VpnDeadPeerDetectionRate_Medium VpnDeadPeerDetectionRate = "medium"
	VpnDeadPeerDetectionRate_None   VpnDeadPeerDetectionRate = "none"
)

func PossibleValuesForVpnDeadPeerDetectionRate() []string {
	return []string{
		string(VpnDeadPeerDetectionRate_High),
		string(VpnDeadPeerDetectionRate_Low),
		string(VpnDeadPeerDetectionRate_Medium),
		string(VpnDeadPeerDetectionRate_None),
	}
}

func (s *VpnDeadPeerDetectionRate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnDeadPeerDetectionRate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnDeadPeerDetectionRate(input string) (*VpnDeadPeerDetectionRate, error) {
	vals := map[string]VpnDeadPeerDetectionRate{
		"high":   VpnDeadPeerDetectionRate_High,
		"low":    VpnDeadPeerDetectionRate_Low,
		"medium": VpnDeadPeerDetectionRate_Medium,
		"none":   VpnDeadPeerDetectionRate_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnDeadPeerDetectionRate(input)
	return &out, nil
}
