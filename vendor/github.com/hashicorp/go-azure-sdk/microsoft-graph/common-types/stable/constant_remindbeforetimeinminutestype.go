package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemindBeforeTimeInMinutesType string

const (
	RemindBeforeTimeInMinutesType_Mins15 RemindBeforeTimeInMinutesType = "mins15"
)

func PossibleValuesForRemindBeforeTimeInMinutesType() []string {
	return []string{
		string(RemindBeforeTimeInMinutesType_Mins15),
	}
}

func (s *RemindBeforeTimeInMinutesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemindBeforeTimeInMinutesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemindBeforeTimeInMinutesType(input string) (*RemindBeforeTimeInMinutesType, error) {
	vals := map[string]RemindBeforeTimeInMinutesType{
		"mins15": RemindBeforeTimeInMinutesType_Mins15,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemindBeforeTimeInMinutesType(input)
	return &out, nil
}
