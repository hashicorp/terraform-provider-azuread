package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataCollectionStatus string

const (
	DataCollectionStatus_Offline DataCollectionStatus = "offline"
	DataCollectionStatus_Online  DataCollectionStatus = "online"
)

func PossibleValuesForDataCollectionStatus() []string {
	return []string{
		string(DataCollectionStatus_Offline),
		string(DataCollectionStatus_Online),
	}
}

func (s *DataCollectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataCollectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataCollectionStatus(input string) (*DataCollectionStatus, error) {
	vals := map[string]DataCollectionStatus{
		"offline": DataCollectionStatus_Offline,
		"online":  DataCollectionStatus_Online,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataCollectionStatus(input)
	return &out, nil
}
