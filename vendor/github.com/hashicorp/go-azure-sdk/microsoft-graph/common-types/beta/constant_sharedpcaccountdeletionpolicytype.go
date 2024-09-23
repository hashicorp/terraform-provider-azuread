package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCAccountDeletionPolicyType string

const (
	SharedPCAccountDeletionPolicyType_DiskSpaceThreshold                    SharedPCAccountDeletionPolicyType = "diskSpaceThreshold"
	SharedPCAccountDeletionPolicyType_DiskSpaceThresholdOrInactiveThreshold SharedPCAccountDeletionPolicyType = "diskSpaceThresholdOrInactiveThreshold"
	SharedPCAccountDeletionPolicyType_Immediate                             SharedPCAccountDeletionPolicyType = "immediate"
)

func PossibleValuesForSharedPCAccountDeletionPolicyType() []string {
	return []string{
		string(SharedPCAccountDeletionPolicyType_DiskSpaceThreshold),
		string(SharedPCAccountDeletionPolicyType_DiskSpaceThresholdOrInactiveThreshold),
		string(SharedPCAccountDeletionPolicyType_Immediate),
	}
}

func (s *SharedPCAccountDeletionPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCAccountDeletionPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCAccountDeletionPolicyType(input string) (*SharedPCAccountDeletionPolicyType, error) {
	vals := map[string]SharedPCAccountDeletionPolicyType{
		"diskspacethreshold":                    SharedPCAccountDeletionPolicyType_DiskSpaceThreshold,
		"diskspacethresholdorinactivethreshold": SharedPCAccountDeletionPolicyType_DiskSpaceThresholdOrInactiveThreshold,
		"immediate":                             SharedPCAccountDeletionPolicyType_Immediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCAccountDeletionPolicyType(input)
	return &out, nil
}
