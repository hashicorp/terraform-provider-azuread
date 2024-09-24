package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationProfileState string

const (
	EducationSynchronizationProfileState_Deleting           EducationSynchronizationProfileState = "deleting"
	EducationSynchronizationProfileState_DeletionFailed     EducationSynchronizationProfileState = "deletionFailed"
	EducationSynchronizationProfileState_Provisioned        EducationSynchronizationProfileState = "provisioned"
	EducationSynchronizationProfileState_Provisioning       EducationSynchronizationProfileState = "provisioning"
	EducationSynchronizationProfileState_ProvisioningFailed EducationSynchronizationProfileState = "provisioningFailed"
)

func PossibleValuesForEducationSynchronizationProfileState() []string {
	return []string{
		string(EducationSynchronizationProfileState_Deleting),
		string(EducationSynchronizationProfileState_DeletionFailed),
		string(EducationSynchronizationProfileState_Provisioned),
		string(EducationSynchronizationProfileState_Provisioning),
		string(EducationSynchronizationProfileState_ProvisioningFailed),
	}
}

func (s *EducationSynchronizationProfileState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationSynchronizationProfileState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationSynchronizationProfileState(input string) (*EducationSynchronizationProfileState, error) {
	vals := map[string]EducationSynchronizationProfileState{
		"deleting":           EducationSynchronizationProfileState_Deleting,
		"deletionfailed":     EducationSynchronizationProfileState_DeletionFailed,
		"provisioned":        EducationSynchronizationProfileState_Provisioned,
		"provisioning":       EducationSynchronizationProfileState_Provisioning,
		"provisioningfailed": EducationSynchronizationProfileState_ProvisioningFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSynchronizationProfileState(input)
	return &out, nil
}
