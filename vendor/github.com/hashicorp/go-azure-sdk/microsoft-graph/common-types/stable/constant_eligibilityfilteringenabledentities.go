package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EligibilityFilteringEnabledEntities string

const (
	EligibilityFilteringEnabledEntities_None              EligibilityFilteringEnabledEntities = "none"
	EligibilityFilteringEnabledEntities_OfferShiftRequest EligibilityFilteringEnabledEntities = "offerShiftRequest"
	EligibilityFilteringEnabledEntities_SwapRequest       EligibilityFilteringEnabledEntities = "swapRequest"
	EligibilityFilteringEnabledEntities_TimeOffReason     EligibilityFilteringEnabledEntities = "timeOffReason"
)

func PossibleValuesForEligibilityFilteringEnabledEntities() []string {
	return []string{
		string(EligibilityFilteringEnabledEntities_None),
		string(EligibilityFilteringEnabledEntities_OfferShiftRequest),
		string(EligibilityFilteringEnabledEntities_SwapRequest),
		string(EligibilityFilteringEnabledEntities_TimeOffReason),
	}
}

func (s *EligibilityFilteringEnabledEntities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEligibilityFilteringEnabledEntities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEligibilityFilteringEnabledEntities(input string) (*EligibilityFilteringEnabledEntities, error) {
	vals := map[string]EligibilityFilteringEnabledEntities{
		"none":              EligibilityFilteringEnabledEntities_None,
		"offershiftrequest": EligibilityFilteringEnabledEntities_OfferShiftRequest,
		"swaprequest":       EligibilityFilteringEnabledEntities_SwapRequest,
		"timeoffreason":     EligibilityFilteringEnabledEntities_TimeOffReason,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EligibilityFilteringEnabledEntities(input)
	return &out, nil
}
