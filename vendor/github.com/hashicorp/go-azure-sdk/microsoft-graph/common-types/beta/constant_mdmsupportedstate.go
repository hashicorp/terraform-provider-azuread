package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmSupportedState string

const (
	MdmSupportedState_Deprecated  MdmSupportedState = "deprecated"
	MdmSupportedState_Supported   MdmSupportedState = "supported"
	MdmSupportedState_Unknown     MdmSupportedState = "unknown"
	MdmSupportedState_Unsupported MdmSupportedState = "unsupported"
)

func PossibleValuesForMdmSupportedState() []string {
	return []string{
		string(MdmSupportedState_Deprecated),
		string(MdmSupportedState_Supported),
		string(MdmSupportedState_Unknown),
		string(MdmSupportedState_Unsupported),
	}
}

func (s *MdmSupportedState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMdmSupportedState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMdmSupportedState(input string) (*MdmSupportedState, error) {
	vals := map[string]MdmSupportedState{
		"deprecated":  MdmSupportedState_Deprecated,
		"supported":   MdmSupportedState_Supported,
		"unknown":     MdmSupportedState_Unknown,
		"unsupported": MdmSupportedState_Unsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmSupportedState(input)
	return &out, nil
}
