package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProductType string

const (
	CloudPCProductType_Business      CloudPCProductType = "business"
	CloudPCProductType_DevBox        CloudPCProductType = "devBox"
	CloudPCProductType_Enterprise    CloudPCProductType = "enterprise"
	CloudPCProductType_Frontline     CloudPCProductType = "frontline"
	CloudPCProductType_PowerAutomate CloudPCProductType = "powerAutomate"
)

func PossibleValuesForCloudPCProductType() []string {
	return []string{
		string(CloudPCProductType_Business),
		string(CloudPCProductType_DevBox),
		string(CloudPCProductType_Enterprise),
		string(CloudPCProductType_Frontline),
		string(CloudPCProductType_PowerAutomate),
	}
}

func (s *CloudPCProductType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCProductType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCProductType(input string) (*CloudPCProductType, error) {
	vals := map[string]CloudPCProductType{
		"business":      CloudPCProductType_Business,
		"devbox":        CloudPCProductType_DevBox,
		"enterprise":    CloudPCProductType_Enterprise,
		"frontline":     CloudPCProductType_Frontline,
		"powerautomate": CloudPCProductType_PowerAutomate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProductType(input)
	return &out, nil
}
