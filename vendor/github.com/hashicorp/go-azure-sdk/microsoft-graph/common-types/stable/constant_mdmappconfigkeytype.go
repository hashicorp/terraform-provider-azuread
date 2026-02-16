package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmAppConfigKeyType string

const (
	MdmAppConfigKeyType_BooleanType MdmAppConfigKeyType = "booleanType"
	MdmAppConfigKeyType_IntegerType MdmAppConfigKeyType = "integerType"
	MdmAppConfigKeyType_RealType    MdmAppConfigKeyType = "realType"
	MdmAppConfigKeyType_StringType  MdmAppConfigKeyType = "stringType"
	MdmAppConfigKeyType_TokenType   MdmAppConfigKeyType = "tokenType"
)

func PossibleValuesForMdmAppConfigKeyType() []string {
	return []string{
		string(MdmAppConfigKeyType_BooleanType),
		string(MdmAppConfigKeyType_IntegerType),
		string(MdmAppConfigKeyType_RealType),
		string(MdmAppConfigKeyType_StringType),
		string(MdmAppConfigKeyType_TokenType),
	}
}

func (s *MdmAppConfigKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMdmAppConfigKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMdmAppConfigKeyType(input string) (*MdmAppConfigKeyType, error) {
	vals := map[string]MdmAppConfigKeyType{
		"booleantype": MdmAppConfigKeyType_BooleanType,
		"integertype": MdmAppConfigKeyType_IntegerType,
		"realtype":    MdmAppConfigKeyType_RealType,
		"stringtype":  MdmAppConfigKeyType_StringType,
		"tokentype":   MdmAppConfigKeyType_TokenType,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmAppConfigKeyType(input)
	return &out, nil
}
