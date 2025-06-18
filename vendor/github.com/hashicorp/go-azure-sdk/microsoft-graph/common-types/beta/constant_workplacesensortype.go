package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkplaceSensorType string

const (
	WorkplaceSensorType_Badge             WorkplaceSensorType = "badge"
	WorkplaceSensorType_Heartbeat         WorkplaceSensorType = "heartbeat"
	WorkplaceSensorType_InferredOccupancy WorkplaceSensorType = "inferredOccupancy"
	WorkplaceSensorType_Occupancy         WorkplaceSensorType = "occupancy"
	WorkplaceSensorType_PeopleCount       WorkplaceSensorType = "peopleCount"
)

func PossibleValuesForWorkplaceSensorType() []string {
	return []string{
		string(WorkplaceSensorType_Badge),
		string(WorkplaceSensorType_Heartbeat),
		string(WorkplaceSensorType_InferredOccupancy),
		string(WorkplaceSensorType_Occupancy),
		string(WorkplaceSensorType_PeopleCount),
	}
}

func (s *WorkplaceSensorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkplaceSensorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkplaceSensorType(input string) (*WorkplaceSensorType, error) {
	vals := map[string]WorkplaceSensorType{
		"badge":             WorkplaceSensorType_Badge,
		"heartbeat":         WorkplaceSensorType_Heartbeat,
		"inferredoccupancy": WorkplaceSensorType_InferredOccupancy,
		"occupancy":         WorkplaceSensorType_Occupancy,
		"peoplecount":       WorkplaceSensorType_PeopleCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkplaceSensorType(input)
	return &out, nil
}
