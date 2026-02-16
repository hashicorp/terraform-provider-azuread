package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalActivityType string

const (
	ExternalConnectorsExternalActivityType_Commented ExternalConnectorsExternalActivityType = "commented"
	ExternalConnectorsExternalActivityType_Created   ExternalConnectorsExternalActivityType = "created"
	ExternalConnectorsExternalActivityType_Modified  ExternalConnectorsExternalActivityType = "modified"
	ExternalConnectorsExternalActivityType_Viewed    ExternalConnectorsExternalActivityType = "viewed"
)

func PossibleValuesForExternalConnectorsExternalActivityType() []string {
	return []string{
		string(ExternalConnectorsExternalActivityType_Commented),
		string(ExternalConnectorsExternalActivityType_Created),
		string(ExternalConnectorsExternalActivityType_Modified),
		string(ExternalConnectorsExternalActivityType_Viewed),
	}
}

func (s *ExternalConnectorsExternalActivityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsExternalActivityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsExternalActivityType(input string) (*ExternalConnectorsExternalActivityType, error) {
	vals := map[string]ExternalConnectorsExternalActivityType{
		"commented": ExternalConnectorsExternalActivityType_Commented,
		"created":   ExternalConnectorsExternalActivityType_Created,
		"modified":  ExternalConnectorsExternalActivityType_Modified,
		"viewed":    ExternalConnectorsExternalActivityType_Viewed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalActivityType(input)
	return &out, nil
}
