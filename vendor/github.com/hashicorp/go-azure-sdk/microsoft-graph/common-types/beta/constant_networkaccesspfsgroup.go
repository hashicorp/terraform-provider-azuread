package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPfsGroup string

const (
	NetworkaccessPfsGroup_Ecp256  NetworkaccessPfsGroup = "ecp256"
	NetworkaccessPfsGroup_Ecp384  NetworkaccessPfsGroup = "ecp384"
	NetworkaccessPfsGroup_None    NetworkaccessPfsGroup = "none"
	NetworkaccessPfsGroup_Pfs1    NetworkaccessPfsGroup = "pfs1"
	NetworkaccessPfsGroup_Pfs14   NetworkaccessPfsGroup = "pfs14"
	NetworkaccessPfsGroup_Pfs2    NetworkaccessPfsGroup = "pfs2"
	NetworkaccessPfsGroup_Pfs2048 NetworkaccessPfsGroup = "pfs2048"
	NetworkaccessPfsGroup_Pfs24   NetworkaccessPfsGroup = "pfs24"
	NetworkaccessPfsGroup_Pfsmm   NetworkaccessPfsGroup = "pfsmm"
)

func PossibleValuesForNetworkaccessPfsGroup() []string {
	return []string{
		string(NetworkaccessPfsGroup_Ecp256),
		string(NetworkaccessPfsGroup_Ecp384),
		string(NetworkaccessPfsGroup_None),
		string(NetworkaccessPfsGroup_Pfs1),
		string(NetworkaccessPfsGroup_Pfs14),
		string(NetworkaccessPfsGroup_Pfs2),
		string(NetworkaccessPfsGroup_Pfs2048),
		string(NetworkaccessPfsGroup_Pfs24),
		string(NetworkaccessPfsGroup_Pfsmm),
	}
}

func (s *NetworkaccessPfsGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessPfsGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessPfsGroup(input string) (*NetworkaccessPfsGroup, error) {
	vals := map[string]NetworkaccessPfsGroup{
		"ecp256":  NetworkaccessPfsGroup_Ecp256,
		"ecp384":  NetworkaccessPfsGroup_Ecp384,
		"none":    NetworkaccessPfsGroup_None,
		"pfs1":    NetworkaccessPfsGroup_Pfs1,
		"pfs14":   NetworkaccessPfsGroup_Pfs14,
		"pfs2":    NetworkaccessPfsGroup_Pfs2,
		"pfs2048": NetworkaccessPfsGroup_Pfs2048,
		"pfs24":   NetworkaccessPfsGroup_Pfs24,
		"pfsmm":   NetworkaccessPfsGroup_Pfsmm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPfsGroup(input)
	return &out, nil
}
