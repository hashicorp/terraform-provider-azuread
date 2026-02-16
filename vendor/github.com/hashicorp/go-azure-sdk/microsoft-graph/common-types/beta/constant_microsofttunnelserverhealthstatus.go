package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelServerHealthStatus string

const (
	MicrosoftTunnelServerHealthStatus_Healthy           MicrosoftTunnelServerHealthStatus = "healthy"
	MicrosoftTunnelServerHealthStatus_Offline           MicrosoftTunnelServerHealthStatus = "offline"
	MicrosoftTunnelServerHealthStatus_Unhealthy         MicrosoftTunnelServerHealthStatus = "unhealthy"
	MicrosoftTunnelServerHealthStatus_Unknown           MicrosoftTunnelServerHealthStatus = "unknown"
	MicrosoftTunnelServerHealthStatus_UpgradeFailed     MicrosoftTunnelServerHealthStatus = "upgradeFailed"
	MicrosoftTunnelServerHealthStatus_UpgradeInProgress MicrosoftTunnelServerHealthStatus = "upgradeInProgress"
	MicrosoftTunnelServerHealthStatus_Warning           MicrosoftTunnelServerHealthStatus = "warning"
)

func PossibleValuesForMicrosoftTunnelServerHealthStatus() []string {
	return []string{
		string(MicrosoftTunnelServerHealthStatus_Healthy),
		string(MicrosoftTunnelServerHealthStatus_Offline),
		string(MicrosoftTunnelServerHealthStatus_Unhealthy),
		string(MicrosoftTunnelServerHealthStatus_Unknown),
		string(MicrosoftTunnelServerHealthStatus_UpgradeFailed),
		string(MicrosoftTunnelServerHealthStatus_UpgradeInProgress),
		string(MicrosoftTunnelServerHealthStatus_Warning),
	}
}

func (s *MicrosoftTunnelServerHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftTunnelServerHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftTunnelServerHealthStatus(input string) (*MicrosoftTunnelServerHealthStatus, error) {
	vals := map[string]MicrosoftTunnelServerHealthStatus{
		"healthy":           MicrosoftTunnelServerHealthStatus_Healthy,
		"offline":           MicrosoftTunnelServerHealthStatus_Offline,
		"unhealthy":         MicrosoftTunnelServerHealthStatus_Unhealthy,
		"unknown":           MicrosoftTunnelServerHealthStatus_Unknown,
		"upgradefailed":     MicrosoftTunnelServerHealthStatus_UpgradeFailed,
		"upgradeinprogress": MicrosoftTunnelServerHealthStatus_UpgradeInProgress,
		"warning":           MicrosoftTunnelServerHealthStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTunnelServerHealthStatus(input)
	return &out, nil
}
