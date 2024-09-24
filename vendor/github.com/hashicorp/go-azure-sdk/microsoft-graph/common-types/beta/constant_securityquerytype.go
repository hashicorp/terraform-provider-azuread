package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityQueryType string

const (
	SecurityQueryType_Files    SecurityQueryType = "files"
	SecurityQueryType_Messages SecurityQueryType = "messages"
)

func PossibleValuesForSecurityQueryType() []string {
	return []string{
		string(SecurityQueryType_Files),
		string(SecurityQueryType_Messages),
	}
}

func (s *SecurityQueryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityQueryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityQueryType(input string) (*SecurityQueryType, error) {
	vals := map[string]SecurityQueryType{
		"files":    SecurityQueryType_Files,
		"messages": SecurityQueryType_Messages,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityQueryType(input)
	return &out, nil
}
