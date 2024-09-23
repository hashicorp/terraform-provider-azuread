package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenAccountType string

const (
	VppTokenAccountType_Business  VppTokenAccountType = "business"
	VppTokenAccountType_Education VppTokenAccountType = "education"
)

func PossibleValuesForVppTokenAccountType() []string {
	return []string{
		string(VppTokenAccountType_Business),
		string(VppTokenAccountType_Education),
	}
}

func (s *VppTokenAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenAccountType(input string) (*VppTokenAccountType, error) {
	vals := map[string]VppTokenAccountType{
		"business":  VppTokenAccountType_Business,
		"education": VppTokenAccountType_Education,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenAccountType(input)
	return &out, nil
}
