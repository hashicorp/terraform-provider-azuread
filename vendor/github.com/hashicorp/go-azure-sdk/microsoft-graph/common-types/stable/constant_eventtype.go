package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventType string

const (
	EventType_Exception      EventType = "exception"
	EventType_Occurrence     EventType = "occurrence"
	EventType_SeriesMaster   EventType = "seriesMaster"
	EventType_SingleInstance EventType = "singleInstance"
)

func PossibleValuesForEventType() []string {
	return []string{
		string(EventType_Exception),
		string(EventType_Occurrence),
		string(EventType_SeriesMaster),
		string(EventType_SingleInstance),
	}
}

func (s *EventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventType(input string) (*EventType, error) {
	vals := map[string]EventType{
		"exception":      EventType_Exception,
		"occurrence":     EventType_Occurrence,
		"seriesmaster":   EventType_SeriesMaster,
		"singleinstance": EventType_SingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventType(input)
	return &out, nil
}
