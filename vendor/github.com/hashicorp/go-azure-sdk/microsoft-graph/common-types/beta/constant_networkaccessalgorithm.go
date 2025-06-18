package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAlgorithm string

const (
	NetworkaccessAlgorithm_Md5      NetworkaccessAlgorithm = "md5"
	NetworkaccessAlgorithm_Sha1     NetworkaccessAlgorithm = "sha1"
	NetworkaccessAlgorithm_Sha256   NetworkaccessAlgorithm = "sha256"
	NetworkaccessAlgorithm_Sha256ac NetworkaccessAlgorithm = "sha256ac"
)

func PossibleValuesForNetworkaccessAlgorithm() []string {
	return []string{
		string(NetworkaccessAlgorithm_Md5),
		string(NetworkaccessAlgorithm_Sha1),
		string(NetworkaccessAlgorithm_Sha256),
		string(NetworkaccessAlgorithm_Sha256ac),
	}
}

func (s *NetworkaccessAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessAlgorithm(input string) (*NetworkaccessAlgorithm, error) {
	vals := map[string]NetworkaccessAlgorithm{
		"md5":      NetworkaccessAlgorithm_Md5,
		"sha1":     NetworkaccessAlgorithm_Sha1,
		"sha256":   NetworkaccessAlgorithm_Sha256,
		"sha256ac": NetworkaccessAlgorithm_Sha256ac,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessAlgorithm(input)
	return &out, nil
}
