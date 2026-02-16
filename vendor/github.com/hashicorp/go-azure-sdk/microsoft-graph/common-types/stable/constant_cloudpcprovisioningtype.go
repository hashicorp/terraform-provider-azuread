package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningType string

const (
	CloudPCProvisioningType_Dedicated CloudPCProvisioningType = "dedicated"
	CloudPCProvisioningType_Shared    CloudPCProvisioningType = "shared"
)

func PossibleValuesForCloudPCProvisioningType() []string {
	return []string{
		string(CloudPCProvisioningType_Dedicated),
		string(CloudPCProvisioningType_Shared),
	}
}

func (s *CloudPCProvisioningType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCProvisioningType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCProvisioningType(input string) (*CloudPCProvisioningType, error) {
	vals := map[string]CloudPCProvisioningType{
		"dedicated": CloudPCProvisioningType_Dedicated,
		"shared":    CloudPCProvisioningType_Shared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningType(input)
	return &out, nil
}
