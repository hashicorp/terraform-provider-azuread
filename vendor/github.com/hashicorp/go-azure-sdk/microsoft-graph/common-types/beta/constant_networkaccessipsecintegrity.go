package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIPSecIntegrity string

const (
	NetworkaccessIPSecIntegrity_GcmAes128 NetworkaccessIPSecIntegrity = "gcmAes128"
	NetworkaccessIPSecIntegrity_GcmAes192 NetworkaccessIPSecIntegrity = "gcmAes192"
	NetworkaccessIPSecIntegrity_GcmAes256 NetworkaccessIPSecIntegrity = "gcmAes256"
	NetworkaccessIPSecIntegrity_Sha256    NetworkaccessIPSecIntegrity = "sha256"
)

func PossibleValuesForNetworkaccessIPSecIntegrity() []string {
	return []string{
		string(NetworkaccessIPSecIntegrity_GcmAes128),
		string(NetworkaccessIPSecIntegrity_GcmAes192),
		string(NetworkaccessIPSecIntegrity_GcmAes256),
		string(NetworkaccessIPSecIntegrity_Sha256),
	}
}

func (s *NetworkaccessIPSecIntegrity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessIPSecIntegrity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessIPSecIntegrity(input string) (*NetworkaccessIPSecIntegrity, error) {
	vals := map[string]NetworkaccessIPSecIntegrity{
		"gcmaes128": NetworkaccessIPSecIntegrity_GcmAes128,
		"gcmaes192": NetworkaccessIPSecIntegrity_GcmAes192,
		"gcmaes256": NetworkaccessIPSecIntegrity_GcmAes256,
		"sha256":    NetworkaccessIPSecIntegrity_Sha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessIPSecIntegrity(input)
	return &out, nil
}
