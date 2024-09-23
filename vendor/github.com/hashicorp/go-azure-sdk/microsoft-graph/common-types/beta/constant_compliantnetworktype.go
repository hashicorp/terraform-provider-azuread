package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliantNetworkType string

const (
	CompliantNetworkType_AllTenantCompliantNetworks CompliantNetworkType = "allTenantCompliantNetworks"
)

func PossibleValuesForCompliantNetworkType() []string {
	return []string{
		string(CompliantNetworkType_AllTenantCompliantNetworks),
	}
}

func (s *CompliantNetworkType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCompliantNetworkType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCompliantNetworkType(input string) (*CompliantNetworkType, error) {
	vals := map[string]CompliantNetworkType{
		"alltenantcompliantnetworks": CompliantNetworkType_AllTenantCompliantNetworks,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompliantNetworkType(input)
	return &out, nil
}
