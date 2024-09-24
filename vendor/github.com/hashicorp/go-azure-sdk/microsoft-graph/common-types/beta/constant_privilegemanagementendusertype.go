package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementEndUserType string

const (
	PrivilegeManagementEndUserType_AzureAd      PrivilegeManagementEndUserType = "azureAd"
	PrivilegeManagementEndUserType_Hybrid       PrivilegeManagementEndUserType = "hybrid"
	PrivilegeManagementEndUserType_Local        PrivilegeManagementEndUserType = "local"
	PrivilegeManagementEndUserType_Undetermined PrivilegeManagementEndUserType = "undetermined"
)

func PossibleValuesForPrivilegeManagementEndUserType() []string {
	return []string{
		string(PrivilegeManagementEndUserType_AzureAd),
		string(PrivilegeManagementEndUserType_Hybrid),
		string(PrivilegeManagementEndUserType_Local),
		string(PrivilegeManagementEndUserType_Undetermined),
	}
}

func (s *PrivilegeManagementEndUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegeManagementEndUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegeManagementEndUserType(input string) (*PrivilegeManagementEndUserType, error) {
	vals := map[string]PrivilegeManagementEndUserType{
		"azuread":      PrivilegeManagementEndUserType_AzureAd,
		"hybrid":       PrivilegeManagementEndUserType_Hybrid,
		"local":        PrivilegeManagementEndUserType_Local,
		"undetermined": PrivilegeManagementEndUserType_Undetermined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementEndUserType(input)
	return &out, nil
}
