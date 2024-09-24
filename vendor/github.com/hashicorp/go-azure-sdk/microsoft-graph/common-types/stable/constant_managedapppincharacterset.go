package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppPinCharacterSet string

const (
	ManagedAppPinCharacterSet_AlphanumericAndSymbol ManagedAppPinCharacterSet = "alphanumericAndSymbol"
	ManagedAppPinCharacterSet_Numeric               ManagedAppPinCharacterSet = "numeric"
)

func PossibleValuesForManagedAppPinCharacterSet() []string {
	return []string{
		string(ManagedAppPinCharacterSet_AlphanumericAndSymbol),
		string(ManagedAppPinCharacterSet_Numeric),
	}
}

func (s *ManagedAppPinCharacterSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppPinCharacterSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppPinCharacterSet(input string) (*ManagedAppPinCharacterSet, error) {
	vals := map[string]ManagedAppPinCharacterSet{
		"alphanumericandsymbol": ManagedAppPinCharacterSet_AlphanumericAndSymbol,
		"numeric":               ManagedAppPinCharacterSet_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppPinCharacterSet(input)
	return &out, nil
}
