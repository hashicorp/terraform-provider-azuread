package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2RestrictionEnforcementType string

const (
	Fido2RestrictionEnforcementType_Allow Fido2RestrictionEnforcementType = "allow"
	Fido2RestrictionEnforcementType_Block Fido2RestrictionEnforcementType = "block"
)

func PossibleValuesForFido2RestrictionEnforcementType() []string {
	return []string{
		string(Fido2RestrictionEnforcementType_Allow),
		string(Fido2RestrictionEnforcementType_Block),
	}
}

func (s *Fido2RestrictionEnforcementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFido2RestrictionEnforcementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFido2RestrictionEnforcementType(input string) (*Fido2RestrictionEnforcementType, error) {
	vals := map[string]Fido2RestrictionEnforcementType{
		"allow": Fido2RestrictionEnforcementType_Allow,
		"block": Fido2RestrictionEnforcementType_Block,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2RestrictionEnforcementType(input)
	return &out, nil
}
