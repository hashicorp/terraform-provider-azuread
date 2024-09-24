package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStepType string

const (
	ProvisioningStepType_Export              ProvisioningStepType = "export"
	ProvisioningStepType_Import              ProvisioningStepType = "import"
	ProvisioningStepType_Matching            ProvisioningStepType = "matching"
	ProvisioningStepType_Processing          ProvisioningStepType = "processing"
	ProvisioningStepType_ReferenceResolution ProvisioningStepType = "referenceResolution"
	ProvisioningStepType_Scoping             ProvisioningStepType = "scoping"
)

func PossibleValuesForProvisioningStepType() []string {
	return []string{
		string(ProvisioningStepType_Export),
		string(ProvisioningStepType_Import),
		string(ProvisioningStepType_Matching),
		string(ProvisioningStepType_Processing),
		string(ProvisioningStepType_ReferenceResolution),
		string(ProvisioningStepType_Scoping),
	}
}

func (s *ProvisioningStepType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningStepType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningStepType(input string) (*ProvisioningStepType, error) {
	vals := map[string]ProvisioningStepType{
		"export":              ProvisioningStepType_Export,
		"import":              ProvisioningStepType_Import,
		"matching":            ProvisioningStepType_Matching,
		"processing":          ProvisioningStepType_Processing,
		"referenceresolution": ProvisioningStepType_ReferenceResolution,
		"scoping":             ProvisioningStepType_Scoping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStepType(input)
	return &out, nil
}
