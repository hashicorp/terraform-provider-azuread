package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiffieHellmanGroup string

const (
	DiffieHellmanGroup_Ecp256  DiffieHellmanGroup = "ecp256"
	DiffieHellmanGroup_Ecp384  DiffieHellmanGroup = "ecp384"
	DiffieHellmanGroup_Group1  DiffieHellmanGroup = "group1"
	DiffieHellmanGroup_Group14 DiffieHellmanGroup = "group14"
	DiffieHellmanGroup_Group2  DiffieHellmanGroup = "group2"
	DiffieHellmanGroup_Group24 DiffieHellmanGroup = "group24"
)

func PossibleValuesForDiffieHellmanGroup() []string {
	return []string{
		string(DiffieHellmanGroup_Ecp256),
		string(DiffieHellmanGroup_Ecp384),
		string(DiffieHellmanGroup_Group1),
		string(DiffieHellmanGroup_Group14),
		string(DiffieHellmanGroup_Group2),
		string(DiffieHellmanGroup_Group24),
	}
}

func (s *DiffieHellmanGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiffieHellmanGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiffieHellmanGroup(input string) (*DiffieHellmanGroup, error) {
	vals := map[string]DiffieHellmanGroup{
		"ecp256":  DiffieHellmanGroup_Ecp256,
		"ecp384":  DiffieHellmanGroup_Ecp384,
		"group1":  DiffieHellmanGroup_Group1,
		"group14": DiffieHellmanGroup_Group14,
		"group2":  DiffieHellmanGroup_Group2,
		"group24": DiffieHellmanGroup_Group24,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiffieHellmanGroup(input)
	return &out, nil
}
