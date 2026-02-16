package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPolicySettingType string

const (
	CloudPCPolicySettingType_Region       CloudPCPolicySettingType = "region"
	CloudPCPolicySettingType_SingleSignOn CloudPCPolicySettingType = "singleSignOn"
)

func PossibleValuesForCloudPCPolicySettingType() []string {
	return []string{
		string(CloudPCPolicySettingType_Region),
		string(CloudPCPolicySettingType_SingleSignOn),
	}
}

func (s *CloudPCPolicySettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPolicySettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPolicySettingType(input string) (*CloudPCPolicySettingType, error) {
	vals := map[string]CloudPCPolicySettingType{
		"region":       CloudPCPolicySettingType_Region,
		"singlesignon": CloudPCPolicySettingType_SingleSignOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPolicySettingType(input)
	return &out, nil
}
