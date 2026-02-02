package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmAuthority string

const (
	MdmAuthority_Intune    MdmAuthority = "intune"
	MdmAuthority_Office365 MdmAuthority = "office365"
	MdmAuthority_Sccm      MdmAuthority = "sccm"
	MdmAuthority_Unknown   MdmAuthority = "unknown"
)

func PossibleValuesForMdmAuthority() []string {
	return []string{
		string(MdmAuthority_Intune),
		string(MdmAuthority_Office365),
		string(MdmAuthority_Sccm),
		string(MdmAuthority_Unknown),
	}
}

func (s *MdmAuthority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMdmAuthority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMdmAuthority(input string) (*MdmAuthority, error) {
	vals := map[string]MdmAuthority{
		"intune":    MdmAuthority_Intune,
		"office365": MdmAuthority_Office365,
		"sccm":      MdmAuthority_Sccm,
		"unknown":   MdmAuthority_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmAuthority(input)
	return &out, nil
}
