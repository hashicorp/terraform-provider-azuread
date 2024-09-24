package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryValueType string

const (
	RegistryValueType_Binary            RegistryValueType = "binary"
	RegistryValueType_Dword             RegistryValueType = "dword"
	RegistryValueType_DwordBigEndian    RegistryValueType = "dwordBigEndian"
	RegistryValueType_DwordLittleEndian RegistryValueType = "dwordLittleEndian"
	RegistryValueType_ExpandSz          RegistryValueType = "expandSz"
	RegistryValueType_Link              RegistryValueType = "link"
	RegistryValueType_MultiSz           RegistryValueType = "multiSz"
	RegistryValueType_None              RegistryValueType = "none"
	RegistryValueType_Qword             RegistryValueType = "qword"
	RegistryValueType_QwordlittleEndian RegistryValueType = "qwordlittleEndian"
	RegistryValueType_Sz                RegistryValueType = "sz"
	RegistryValueType_Unknown           RegistryValueType = "unknown"
)

func PossibleValuesForRegistryValueType() []string {
	return []string{
		string(RegistryValueType_Binary),
		string(RegistryValueType_Dword),
		string(RegistryValueType_DwordBigEndian),
		string(RegistryValueType_DwordLittleEndian),
		string(RegistryValueType_ExpandSz),
		string(RegistryValueType_Link),
		string(RegistryValueType_MultiSz),
		string(RegistryValueType_None),
		string(RegistryValueType_Qword),
		string(RegistryValueType_QwordlittleEndian),
		string(RegistryValueType_Sz),
		string(RegistryValueType_Unknown),
	}
}

func (s *RegistryValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryValueType(input string) (*RegistryValueType, error) {
	vals := map[string]RegistryValueType{
		"binary":            RegistryValueType_Binary,
		"dword":             RegistryValueType_Dword,
		"dwordbigendian":    RegistryValueType_DwordBigEndian,
		"dwordlittleendian": RegistryValueType_DwordLittleEndian,
		"expandsz":          RegistryValueType_ExpandSz,
		"link":              RegistryValueType_Link,
		"multisz":           RegistryValueType_MultiSz,
		"none":              RegistryValueType_None,
		"qword":             RegistryValueType_Qword,
		"qwordlittleendian": RegistryValueType_QwordlittleEndian,
		"sz":                RegistryValueType_Sz,
		"unknown":           RegistryValueType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryValueType(input)
	return &out, nil
}
