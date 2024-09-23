package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAttestationSettingStatus string

const (
	AzureAttestationSettingStatus_Disabled      AzureAttestationSettingStatus = "disabled"
	AzureAttestationSettingStatus_Enabled       AzureAttestationSettingStatus = "enabled"
	AzureAttestationSettingStatus_NotApplicable AzureAttestationSettingStatus = "notApplicable"
)

func PossibleValuesForAzureAttestationSettingStatus() []string {
	return []string{
		string(AzureAttestationSettingStatus_Disabled),
		string(AzureAttestationSettingStatus_Enabled),
		string(AzureAttestationSettingStatus_NotApplicable),
	}
}

func (s *AzureAttestationSettingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureAttestationSettingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureAttestationSettingStatus(input string) (*AzureAttestationSettingStatus, error) {
	vals := map[string]AzureAttestationSettingStatus{
		"disabled":      AzureAttestationSettingStatus_Disabled,
		"enabled":       AzureAttestationSettingStatus_Enabled,
		"notapplicable": AzureAttestationSettingStatus_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureAttestationSettingStatus(input)
	return &out, nil
}
