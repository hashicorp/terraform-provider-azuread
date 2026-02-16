package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeySize string

const (
	KeySize_Size1024 KeySize = "size1024"
	KeySize_Size2048 KeySize = "size2048"
	KeySize_Size4096 KeySize = "size4096"
)

func PossibleValuesForKeySize() []string {
	return []string{
		string(KeySize_Size1024),
		string(KeySize_Size2048),
		string(KeySize_Size4096),
	}
}

func (s *KeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKeySize(input string) (*KeySize, error) {
	vals := map[string]KeySize{
		"size1024": KeySize_Size1024,
		"size2048": KeySize_Size2048,
		"size4096": KeySize_Size4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeySize(input)
	return &out, nil
}
