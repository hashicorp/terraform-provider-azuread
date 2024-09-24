package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExternalPartnerStatus string

const (
	CloudPCExternalPartnerStatus_Available    CloudPCExternalPartnerStatus = "available"
	CloudPCExternalPartnerStatus_Healthy      CloudPCExternalPartnerStatus = "healthy"
	CloudPCExternalPartnerStatus_NotAvailable CloudPCExternalPartnerStatus = "notAvailable"
	CloudPCExternalPartnerStatus_Unhealthy    CloudPCExternalPartnerStatus = "unhealthy"
)

func PossibleValuesForCloudPCExternalPartnerStatus() []string {
	return []string{
		string(CloudPCExternalPartnerStatus_Available),
		string(CloudPCExternalPartnerStatus_Healthy),
		string(CloudPCExternalPartnerStatus_NotAvailable),
		string(CloudPCExternalPartnerStatus_Unhealthy),
	}
}

func (s *CloudPCExternalPartnerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCExternalPartnerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCExternalPartnerStatus(input string) (*CloudPCExternalPartnerStatus, error) {
	vals := map[string]CloudPCExternalPartnerStatus{
		"available":    CloudPCExternalPartnerStatus_Available,
		"healthy":      CloudPCExternalPartnerStatus_Healthy,
		"notavailable": CloudPCExternalPartnerStatus_NotAvailable,
		"unhealthy":    CloudPCExternalPartnerStatus_Unhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExternalPartnerStatus(input)
	return &out, nil
}
