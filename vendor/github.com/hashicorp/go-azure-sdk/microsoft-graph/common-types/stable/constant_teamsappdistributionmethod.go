package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDistributionMethod string

const (
	TeamsAppDistributionMethod_Organization TeamsAppDistributionMethod = "organization"
	TeamsAppDistributionMethod_Sideloaded   TeamsAppDistributionMethod = "sideloaded"
	TeamsAppDistributionMethod_Store        TeamsAppDistributionMethod = "store"
)

func PossibleValuesForTeamsAppDistributionMethod() []string {
	return []string{
		string(TeamsAppDistributionMethod_Organization),
		string(TeamsAppDistributionMethod_Sideloaded),
		string(TeamsAppDistributionMethod_Store),
	}
}

func (s *TeamsAppDistributionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDistributionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDistributionMethod(input string) (*TeamsAppDistributionMethod, error) {
	vals := map[string]TeamsAppDistributionMethod{
		"organization": TeamsAppDistributionMethod_Organization,
		"sideloaded":   TeamsAppDistributionMethod_Sideloaded,
		"store":        TeamsAppDistributionMethod_Store,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDistributionMethod(input)
	return &out, nil
}
