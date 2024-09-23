package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuarantineReason string

const (
	QuarantineReason_EncounteredBaseEscrowThreshold       QuarantineReason = "EncounteredBaseEscrowThreshold"
	QuarantineReason_EncounteredEscrowProportionThreshold QuarantineReason = "EncounteredEscrowProportionThreshold"
	QuarantineReason_EncounteredQuarantineException       QuarantineReason = "EncounteredQuarantineException"
	QuarantineReason_EncounteredTotalEscrowThreshold      QuarantineReason = "EncounteredTotalEscrowThreshold"
	QuarantineReason_IngestionInterrupted                 QuarantineReason = "IngestionInterrupted"
	QuarantineReason_QuarantinedOnDemand                  QuarantineReason = "QuarantinedOnDemand"
	QuarantineReason_TooManyDeletes                       QuarantineReason = "TooManyDeletes"
	QuarantineReason_Unknown                              QuarantineReason = "Unknown"
)

func PossibleValuesForQuarantineReason() []string {
	return []string{
		string(QuarantineReason_EncounteredBaseEscrowThreshold),
		string(QuarantineReason_EncounteredEscrowProportionThreshold),
		string(QuarantineReason_EncounteredQuarantineException),
		string(QuarantineReason_EncounteredTotalEscrowThreshold),
		string(QuarantineReason_IngestionInterrupted),
		string(QuarantineReason_QuarantinedOnDemand),
		string(QuarantineReason_TooManyDeletes),
		string(QuarantineReason_Unknown),
	}
}

func (s *QuarantineReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQuarantineReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQuarantineReason(input string) (*QuarantineReason, error) {
	vals := map[string]QuarantineReason{
		"encounteredbaseescrowthreshold":       QuarantineReason_EncounteredBaseEscrowThreshold,
		"encounteredescrowproportionthreshold": QuarantineReason_EncounteredEscrowProportionThreshold,
		"encounteredquarantineexception":       QuarantineReason_EncounteredQuarantineException,
		"encounteredtotalescrowthreshold":      QuarantineReason_EncounteredTotalEscrowThreshold,
		"ingestioninterrupted":                 QuarantineReason_IngestionInterrupted,
		"quarantinedondemand":                  QuarantineReason_QuarantinedOnDemand,
		"toomanydeletes":                       QuarantineReason_TooManyDeletes,
		"unknown":                              QuarantineReason_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QuarantineReason(input)
	return &out, nil
}
