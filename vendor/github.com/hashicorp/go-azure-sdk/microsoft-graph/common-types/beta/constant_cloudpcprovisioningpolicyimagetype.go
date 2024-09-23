package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicyImageType string

const (
	CloudPCProvisioningPolicyImageType_Custom  CloudPCProvisioningPolicyImageType = "custom"
	CloudPCProvisioningPolicyImageType_Gallery CloudPCProvisioningPolicyImageType = "gallery"
)

func PossibleValuesForCloudPCProvisioningPolicyImageType() []string {
	return []string{
		string(CloudPCProvisioningPolicyImageType_Custom),
		string(CloudPCProvisioningPolicyImageType_Gallery),
	}
}

func (s *CloudPCProvisioningPolicyImageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCProvisioningPolicyImageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCProvisioningPolicyImageType(input string) (*CloudPCProvisioningPolicyImageType, error) {
	vals := map[string]CloudPCProvisioningPolicyImageType{
		"custom":  CloudPCProvisioningPolicyImageType_Custom,
		"gallery": CloudPCProvisioningPolicyImageType_Gallery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningPolicyImageType(input)
	return &out, nil
}
