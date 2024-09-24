package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptWith string

const (
	EncryptWith_Template          EncryptWith = "template"
	EncryptWith_UserDefinedRights EncryptWith = "userDefinedRights"
)

func PossibleValuesForEncryptWith() []string {
	return []string{
		string(EncryptWith_Template),
		string(EncryptWith_UserDefinedRights),
	}
}

func (s *EncryptWith) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptWith(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptWith(input string) (*EncryptWith, error) {
	vals := map[string]EncryptWith{
		"template":          EncryptWith_Template,
		"userdefinedrights": EncryptWith_UserDefinedRights,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptWith(input)
	return &out, nil
}
