package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnIntegrityAlgorithmType string

const (
	VpnIntegrityAlgorithmType_Md5     VpnIntegrityAlgorithmType = "md5"
	VpnIntegrityAlgorithmType_Sha1160 VpnIntegrityAlgorithmType = "sha1_160"
	VpnIntegrityAlgorithmType_Sha196  VpnIntegrityAlgorithmType = "sha1_96"
	VpnIntegrityAlgorithmType_Sha2256 VpnIntegrityAlgorithmType = "sha2_256"
	VpnIntegrityAlgorithmType_Sha2384 VpnIntegrityAlgorithmType = "sha2_384"
	VpnIntegrityAlgorithmType_Sha2512 VpnIntegrityAlgorithmType = "sha2_512"
)

func PossibleValuesForVpnIntegrityAlgorithmType() []string {
	return []string{
		string(VpnIntegrityAlgorithmType_Md5),
		string(VpnIntegrityAlgorithmType_Sha1160),
		string(VpnIntegrityAlgorithmType_Sha196),
		string(VpnIntegrityAlgorithmType_Sha2256),
		string(VpnIntegrityAlgorithmType_Sha2384),
		string(VpnIntegrityAlgorithmType_Sha2512),
	}
}

func (s *VpnIntegrityAlgorithmType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnIntegrityAlgorithmType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnIntegrityAlgorithmType(input string) (*VpnIntegrityAlgorithmType, error) {
	vals := map[string]VpnIntegrityAlgorithmType{
		"md5":      VpnIntegrityAlgorithmType_Md5,
		"sha1_160": VpnIntegrityAlgorithmType_Sha1160,
		"sha1_96":  VpnIntegrityAlgorithmType_Sha196,
		"sha2_256": VpnIntegrityAlgorithmType_Sha2256,
		"sha2_384": VpnIntegrityAlgorithmType_Sha2384,
		"sha2_512": VpnIntegrityAlgorithmType_Sha2512,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnIntegrityAlgorithmType(input)
	return &out, nil
}
