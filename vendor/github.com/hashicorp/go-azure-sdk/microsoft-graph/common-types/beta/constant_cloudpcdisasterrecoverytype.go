package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDisasterRecoveryType string

const (
	CloudPCDisasterRecoveryType_CrossRegion   CloudPCDisasterRecoveryType = "crossRegion"
	CloudPCDisasterRecoveryType_NotConfigured CloudPCDisasterRecoveryType = "notConfigured"
	CloudPCDisasterRecoveryType_Premium       CloudPCDisasterRecoveryType = "premium"
)

func PossibleValuesForCloudPCDisasterRecoveryType() []string {
	return []string{
		string(CloudPCDisasterRecoveryType_CrossRegion),
		string(CloudPCDisasterRecoveryType_NotConfigured),
		string(CloudPCDisasterRecoveryType_Premium),
	}
}

func (s *CloudPCDisasterRecoveryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDisasterRecoveryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDisasterRecoveryType(input string) (*CloudPCDisasterRecoveryType, error) {
	vals := map[string]CloudPCDisasterRecoveryType{
		"crossregion":   CloudPCDisasterRecoveryType_CrossRegion,
		"notconfigured": CloudPCDisasterRecoveryType_NotConfigured,
		"premium":       CloudPCDisasterRecoveryType_Premium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDisasterRecoveryType(input)
	return &out, nil
}
