package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HashAlgorithms string

const (
	HashAlgorithms_Sha1 HashAlgorithms = "sha1"
	HashAlgorithms_Sha2 HashAlgorithms = "sha2"
)

func PossibleValuesForHashAlgorithms() []string {
	return []string{
		string(HashAlgorithms_Sha1),
		string(HashAlgorithms_Sha2),
	}
}

func (s *HashAlgorithms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHashAlgorithms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHashAlgorithms(input string) (*HashAlgorithms, error) {
	vals := map[string]HashAlgorithms{
		"sha1": HashAlgorithms_Sha1,
		"sha2": HashAlgorithms_Sha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HashAlgorithms(input)
	return &out, nil
}
