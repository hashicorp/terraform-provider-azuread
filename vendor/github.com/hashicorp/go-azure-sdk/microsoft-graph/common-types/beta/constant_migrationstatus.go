package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationStatus string

const (
	MigrationStatus_AdditionalStepsRequired MigrationStatus = "additionalStepsRequired"
	MigrationStatus_NeedsReview             MigrationStatus = "needsReview"
	MigrationStatus_Ready                   MigrationStatus = "ready"
)

func PossibleValuesForMigrationStatus() []string {
	return []string{
		string(MigrationStatus_AdditionalStepsRequired),
		string(MigrationStatus_NeedsReview),
		string(MigrationStatus_Ready),
	}
}

func (s *MigrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationStatus(input string) (*MigrationStatus, error) {
	vals := map[string]MigrationStatus{
		"additionalstepsrequired": MigrationStatus_AdditionalStepsRequired,
		"needsreview":             MigrationStatus_NeedsReview,
		"ready":                   MigrationStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationStatus(input)
	return &out, nil
}
