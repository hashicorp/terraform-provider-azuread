package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIPSecEncryption string

const (
	NetworkaccessIPSecEncryption_GcmAes128 NetworkaccessIPSecEncryption = "gcmAes128"
	NetworkaccessIPSecEncryption_GcmAes192 NetworkaccessIPSecEncryption = "gcmAes192"
	NetworkaccessIPSecEncryption_GcmAes256 NetworkaccessIPSecEncryption = "gcmAes256"
	NetworkaccessIPSecEncryption_None      NetworkaccessIPSecEncryption = "none"
)

func PossibleValuesForNetworkaccessIPSecEncryption() []string {
	return []string{
		string(NetworkaccessIPSecEncryption_GcmAes128),
		string(NetworkaccessIPSecEncryption_GcmAes192),
		string(NetworkaccessIPSecEncryption_GcmAes256),
		string(NetworkaccessIPSecEncryption_None),
	}
}

func (s *NetworkaccessIPSecEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessIPSecEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessIPSecEncryption(input string) (*NetworkaccessIPSecEncryption, error) {
	vals := map[string]NetworkaccessIPSecEncryption{
		"gcmaes128": NetworkaccessIPSecEncryption_GcmAes128,
		"gcmaes192": NetworkaccessIPSecEncryption_GcmAes192,
		"gcmaes256": NetworkaccessIPSecEncryption_GcmAes256,
		"none":      NetworkaccessIPSecEncryption_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessIPSecEncryption(input)
	return &out, nil
}
