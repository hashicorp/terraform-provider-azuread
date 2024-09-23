package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogDecryptionAlgorithm string

const (
	AppLogDecryptionAlgorithm_Aes256 AppLogDecryptionAlgorithm = "aes256"
)

func PossibleValuesForAppLogDecryptionAlgorithm() []string {
	return []string{
		string(AppLogDecryptionAlgorithm_Aes256),
	}
}

func (s *AppLogDecryptionAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppLogDecryptionAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppLogDecryptionAlgorithm(input string) (*AppLogDecryptionAlgorithm, error) {
	vals := map[string]AppLogDecryptionAlgorithm{
		"aes256": AppLogDecryptionAlgorithm_Aes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLogDecryptionAlgorithm(input)
	return &out, nil
}
