package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkplaceSensorEventType string

const (
	WorkplaceSensorEventType_BadgeIn  WorkplaceSensorEventType = "badgeIn"
	WorkplaceSensorEventType_BadgeOut WorkplaceSensorEventType = "badgeOut"
)

func PossibleValuesForWorkplaceSensorEventType() []string {
	return []string{
		string(WorkplaceSensorEventType_BadgeIn),
		string(WorkplaceSensorEventType_BadgeOut),
	}
}

func (s *WorkplaceSensorEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkplaceSensorEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkplaceSensorEventType(input string) (*WorkplaceSensorEventType, error) {
	vals := map[string]WorkplaceSensorEventType{
		"badgein":  WorkplaceSensorEventType_BadgeIn,
		"badgeout": WorkplaceSensorEventType_BadgeOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkplaceSensorEventType(input)
	return &out, nil
}
