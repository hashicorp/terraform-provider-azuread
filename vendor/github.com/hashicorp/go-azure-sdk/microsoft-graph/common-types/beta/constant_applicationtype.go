package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationType string

const (
	ApplicationType_Desktop   ApplicationType = "desktop"
	ApplicationType_Universal ApplicationType = "universal"
)

func PossibleValuesForApplicationType() []string {
	return []string{
		string(ApplicationType_Desktop),
		string(ApplicationType_Universal),
	}
}

func (s *ApplicationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationType(input string) (*ApplicationType, error) {
	vals := map[string]ApplicationType{
		"desktop":   ApplicationType_Desktop,
		"universal": ApplicationType_Universal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationType(input)
	return &out, nil
}
