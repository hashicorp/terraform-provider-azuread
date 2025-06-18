package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DestinationType string

const (
	DestinationType_InPlace DestinationType = "inPlace"
	DestinationType_New     DestinationType = "new"
)

func PossibleValuesForDestinationType() []string {
	return []string{
		string(DestinationType_InPlace),
		string(DestinationType_New),
	}
}

func (s *DestinationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDestinationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDestinationType(input string) (*DestinationType, error) {
	vals := map[string]DestinationType{
		"inplace": DestinationType_InPlace,
		"new":     DestinationType_New,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DestinationType(input)
	return &out, nil
}
