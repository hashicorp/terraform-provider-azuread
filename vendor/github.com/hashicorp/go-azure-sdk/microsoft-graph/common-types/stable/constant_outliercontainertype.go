package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlierContainerType string

const (
	OutlierContainerType_Group OutlierContainerType = "group"
)

func PossibleValuesForOutlierContainerType() []string {
	return []string{
		string(OutlierContainerType_Group),
	}
}

func (s *OutlierContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlierContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlierContainerType(input string) (*OutlierContainerType, error) {
	vals := map[string]OutlierContainerType{
		"group": OutlierContainerType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlierContainerType(input)
	return &out, nil
}
