package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAlertType string

const (
	NetworkaccessAlertType_CrossTenantAnomaly             NetworkaccessAlertType = "crossTenantAnomaly"
	NetworkaccessAlertType_DeviceTokenInconsistency       NetworkaccessAlertType = "deviceTokenInconsistency"
	NetworkaccessAlertType_Dlp                            NetworkaccessAlertType = "dlp"
	NetworkaccessAlertType_Malware                        NetworkaccessAlertType = "malware"
	NetworkaccessAlertType_PatientZero                    NetworkaccessAlertType = "patientZero"
	NetworkaccessAlertType_SuspiciousProcess              NetworkaccessAlertType = "suspiciousProcess"
	NetworkaccessAlertType_ThreatIntelligenceTransactions NetworkaccessAlertType = "threatIntelligenceTransactions"
	NetworkaccessAlertType_UnhealthyConnectors            NetworkaccessAlertType = "unhealthyConnectors"
	NetworkaccessAlertType_UnhealthyRemoteNetworks        NetworkaccessAlertType = "unhealthyRemoteNetworks"
	NetworkaccessAlertType_WebContentBlocked              NetworkaccessAlertType = "webContentBlocked"
)

func PossibleValuesForNetworkaccessAlertType() []string {
	return []string{
		string(NetworkaccessAlertType_CrossTenantAnomaly),
		string(NetworkaccessAlertType_DeviceTokenInconsistency),
		string(NetworkaccessAlertType_Dlp),
		string(NetworkaccessAlertType_Malware),
		string(NetworkaccessAlertType_PatientZero),
		string(NetworkaccessAlertType_SuspiciousProcess),
		string(NetworkaccessAlertType_ThreatIntelligenceTransactions),
		string(NetworkaccessAlertType_UnhealthyConnectors),
		string(NetworkaccessAlertType_UnhealthyRemoteNetworks),
		string(NetworkaccessAlertType_WebContentBlocked),
	}
}

func (s *NetworkaccessAlertType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessAlertType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessAlertType(input string) (*NetworkaccessAlertType, error) {
	vals := map[string]NetworkaccessAlertType{
		"crosstenantanomaly":             NetworkaccessAlertType_CrossTenantAnomaly,
		"devicetokeninconsistency":       NetworkaccessAlertType_DeviceTokenInconsistency,
		"dlp":                            NetworkaccessAlertType_Dlp,
		"malware":                        NetworkaccessAlertType_Malware,
		"patientzero":                    NetworkaccessAlertType_PatientZero,
		"suspiciousprocess":              NetworkaccessAlertType_SuspiciousProcess,
		"threatintelligencetransactions": NetworkaccessAlertType_ThreatIntelligenceTransactions,
		"unhealthyconnectors":            NetworkaccessAlertType_UnhealthyConnectors,
		"unhealthyremotenetworks":        NetworkaccessAlertType_UnhealthyRemoteNetworks,
		"webcontentblocked":              NetworkaccessAlertType_WebContentBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessAlertType(input)
	return &out, nil
}
