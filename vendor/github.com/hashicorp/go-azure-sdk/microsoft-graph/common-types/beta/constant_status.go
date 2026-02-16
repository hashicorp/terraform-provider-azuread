package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Status string

const (
	Status_Active  Status = "active"
	Status_Deleted Status = "deleted"
	Status_Ignored Status = "ignored"
	Status_Updated Status = "updated"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(Status_Active),
		string(Status_Deleted),
		string(Status_Ignored),
		string(Status_Updated),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"active":  Status_Active,
		"deleted": Status_Deleted,
		"ignored": Status_Ignored,
		"updated": Status_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
