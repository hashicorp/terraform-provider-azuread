package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIkeEncryption string

const (
	NetworkaccessIkeEncryption_Aes128    NetworkaccessIkeEncryption = "aes128"
	NetworkaccessIkeEncryption_Aes192    NetworkaccessIkeEncryption = "aes192"
	NetworkaccessIkeEncryption_Aes256    NetworkaccessIkeEncryption = "aes256"
	NetworkaccessIkeEncryption_GcmAes128 NetworkaccessIkeEncryption = "gcmAes128"
	NetworkaccessIkeEncryption_GcmAes256 NetworkaccessIkeEncryption = "gcmAes256"
)

func PossibleValuesForNetworkaccessIkeEncryption() []string {
	return []string{
		string(NetworkaccessIkeEncryption_Aes128),
		string(NetworkaccessIkeEncryption_Aes192),
		string(NetworkaccessIkeEncryption_Aes256),
		string(NetworkaccessIkeEncryption_GcmAes128),
		string(NetworkaccessIkeEncryption_GcmAes256),
	}
}

func (s *NetworkaccessIkeEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessIkeEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessIkeEncryption(input string) (*NetworkaccessIkeEncryption, error) {
	vals := map[string]NetworkaccessIkeEncryption{
		"aes128":    NetworkaccessIkeEncryption_Aes128,
		"aes192":    NetworkaccessIkeEncryption_Aes192,
		"aes256":    NetworkaccessIkeEncryption_Aes256,
		"gcmaes128": NetworkaccessIkeEncryption_GcmAes128,
		"gcmaes256": NetworkaccessIkeEncryption_GcmAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessIkeEncryption(input)
	return &out, nil
}
