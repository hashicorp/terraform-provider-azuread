package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EngagementAsyncOperationType string

const (
	EngagementAsyncOperationType_CreateCommunity EngagementAsyncOperationType = "createCommunity"
)

func PossibleValuesForEngagementAsyncOperationType() []string {
	return []string{
		string(EngagementAsyncOperationType_CreateCommunity),
	}
}

func (s *EngagementAsyncOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEngagementAsyncOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEngagementAsyncOperationType(input string) (*EngagementAsyncOperationType, error) {
	vals := map[string]EngagementAsyncOperationType{
		"createcommunity": EngagementAsyncOperationType_CreateCommunity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EngagementAsyncOperationType(input)
	return &out, nil
}
