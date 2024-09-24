package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceContainerStatus string

const (
	EdiscoveryDataSourceContainerStatus_Active   EdiscoveryDataSourceContainerStatus = "Active"
	EdiscoveryDataSourceContainerStatus_Released EdiscoveryDataSourceContainerStatus = "Released"
)

func PossibleValuesForEdiscoveryDataSourceContainerStatus() []string {
	return []string{
		string(EdiscoveryDataSourceContainerStatus_Active),
		string(EdiscoveryDataSourceContainerStatus_Released),
	}
}

func (s *EdiscoveryDataSourceContainerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryDataSourceContainerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryDataSourceContainerStatus(input string) (*EdiscoveryDataSourceContainerStatus, error) {
	vals := map[string]EdiscoveryDataSourceContainerStatus{
		"active":   EdiscoveryDataSourceContainerStatus_Active,
		"released": EdiscoveryDataSourceContainerStatus_Released,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryDataSourceContainerStatus(input)
	return &out, nil
}
