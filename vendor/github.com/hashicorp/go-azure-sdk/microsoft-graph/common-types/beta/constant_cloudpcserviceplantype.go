package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCServicePlanType string

const (
	CloudPCServicePlanType_Business   CloudPCServicePlanType = "business"
	CloudPCServicePlanType_Enterprise CloudPCServicePlanType = "enterprise"
)

func PossibleValuesForCloudPCServicePlanType() []string {
	return []string{
		string(CloudPCServicePlanType_Business),
		string(CloudPCServicePlanType_Enterprise),
	}
}

func (s *CloudPCServicePlanType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCServicePlanType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCServicePlanType(input string) (*CloudPCServicePlanType, error) {
	vals := map[string]CloudPCServicePlanType{
		"business":   CloudPCServicePlanType_Business,
		"enterprise": CloudPCServicePlanType_Enterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCServicePlanType(input)
	return &out, nil
}
