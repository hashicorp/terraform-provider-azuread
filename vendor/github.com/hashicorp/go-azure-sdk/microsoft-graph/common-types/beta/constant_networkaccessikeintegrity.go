package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIkeIntegrity string

const (
	NetworkaccessIkeIntegrity_GcmAes128 NetworkaccessIkeIntegrity = "gcmAes128"
	NetworkaccessIkeIntegrity_GcmAes256 NetworkaccessIkeIntegrity = "gcmAes256"
	NetworkaccessIkeIntegrity_Sha256    NetworkaccessIkeIntegrity = "sha256"
	NetworkaccessIkeIntegrity_Sha384    NetworkaccessIkeIntegrity = "sha384"
)

func PossibleValuesForNetworkaccessIkeIntegrity() []string {
	return []string{
		string(NetworkaccessIkeIntegrity_GcmAes128),
		string(NetworkaccessIkeIntegrity_GcmAes256),
		string(NetworkaccessIkeIntegrity_Sha256),
		string(NetworkaccessIkeIntegrity_Sha384),
	}
}

func (s *NetworkaccessIkeIntegrity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessIkeIntegrity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessIkeIntegrity(input string) (*NetworkaccessIkeIntegrity, error) {
	vals := map[string]NetworkaccessIkeIntegrity{
		"gcmaes128": NetworkaccessIkeIntegrity_GcmAes128,
		"gcmaes256": NetworkaccessIkeIntegrity_GcmAes256,
		"sha256":    NetworkaccessIkeIntegrity_Sha256,
		"sha384":    NetworkaccessIkeIntegrity_Sha384,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessIkeIntegrity(input)
	return &out, nil
}
