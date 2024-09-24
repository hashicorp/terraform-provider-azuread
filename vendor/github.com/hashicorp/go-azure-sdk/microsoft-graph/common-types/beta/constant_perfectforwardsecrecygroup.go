package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerfectForwardSecrecyGroup string

const (
	PerfectForwardSecrecyGroup_Ecp256  PerfectForwardSecrecyGroup = "ecp256"
	PerfectForwardSecrecyGroup_Ecp384  PerfectForwardSecrecyGroup = "ecp384"
	PerfectForwardSecrecyGroup_Pfs1    PerfectForwardSecrecyGroup = "pfs1"
	PerfectForwardSecrecyGroup_Pfs2    PerfectForwardSecrecyGroup = "pfs2"
	PerfectForwardSecrecyGroup_Pfs2048 PerfectForwardSecrecyGroup = "pfs2048"
	PerfectForwardSecrecyGroup_Pfs24   PerfectForwardSecrecyGroup = "pfs24"
	PerfectForwardSecrecyGroup_PfsMM   PerfectForwardSecrecyGroup = "pfsMM"
)

func PossibleValuesForPerfectForwardSecrecyGroup() []string {
	return []string{
		string(PerfectForwardSecrecyGroup_Ecp256),
		string(PerfectForwardSecrecyGroup_Ecp384),
		string(PerfectForwardSecrecyGroup_Pfs1),
		string(PerfectForwardSecrecyGroup_Pfs2),
		string(PerfectForwardSecrecyGroup_Pfs2048),
		string(PerfectForwardSecrecyGroup_Pfs24),
		string(PerfectForwardSecrecyGroup_PfsMM),
	}
}

func (s *PerfectForwardSecrecyGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePerfectForwardSecrecyGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePerfectForwardSecrecyGroup(input string) (*PerfectForwardSecrecyGroup, error) {
	vals := map[string]PerfectForwardSecrecyGroup{
		"ecp256":  PerfectForwardSecrecyGroup_Ecp256,
		"ecp384":  PerfectForwardSecrecyGroup_Ecp384,
		"pfs1":    PerfectForwardSecrecyGroup_Pfs1,
		"pfs2":    PerfectForwardSecrecyGroup_Pfs2,
		"pfs2048": PerfectForwardSecrecyGroup_Pfs2048,
		"pfs24":   PerfectForwardSecrecyGroup_Pfs24,
		"pfsmm":   PerfectForwardSecrecyGroup_PfsMM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PerfectForwardSecrecyGroup(input)
	return &out, nil
}
