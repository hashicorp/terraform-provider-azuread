package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStatusErrorCategory string

const (
	ProvisioningStatusErrorCategory_Failure           ProvisioningStatusErrorCategory = "failure"
	ProvisioningStatusErrorCategory_NonServiceFailure ProvisioningStatusErrorCategory = "nonServiceFailure"
	ProvisioningStatusErrorCategory_Success           ProvisioningStatusErrorCategory = "success"
)

func PossibleValuesForProvisioningStatusErrorCategory() []string {
	return []string{
		string(ProvisioningStatusErrorCategory_Failure),
		string(ProvisioningStatusErrorCategory_NonServiceFailure),
		string(ProvisioningStatusErrorCategory_Success),
	}
}

func (s *ProvisioningStatusErrorCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningStatusErrorCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningStatusErrorCategory(input string) (*ProvisioningStatusErrorCategory, error) {
	vals := map[string]ProvisioningStatusErrorCategory{
		"failure":           ProvisioningStatusErrorCategory_Failure,
		"nonservicefailure": ProvisioningStatusErrorCategory_NonServiceFailure,
		"success":           ProvisioningStatusErrorCategory_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStatusErrorCategory(input)
	return &out, nil
}
