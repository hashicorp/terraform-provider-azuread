package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingAssignedTo string

const (
	TrainingAssignedTo_AllUsers          TrainingAssignedTo = "allUsers"
	TrainingAssignedTo_ClickedPayload    TrainingAssignedTo = "clickedPayload"
	TrainingAssignedTo_Compromised       TrainingAssignedTo = "compromised"
	TrainingAssignedTo_DidNothing        TrainingAssignedTo = "didNothing"
	TrainingAssignedTo_None              TrainingAssignedTo = "none"
	TrainingAssignedTo_ReadButNotClicked TrainingAssignedTo = "readButNotClicked"
	TrainingAssignedTo_ReportedPhish     TrainingAssignedTo = "reportedPhish"
)

func PossibleValuesForTrainingAssignedTo() []string {
	return []string{
		string(TrainingAssignedTo_AllUsers),
		string(TrainingAssignedTo_ClickedPayload),
		string(TrainingAssignedTo_Compromised),
		string(TrainingAssignedTo_DidNothing),
		string(TrainingAssignedTo_None),
		string(TrainingAssignedTo_ReadButNotClicked),
		string(TrainingAssignedTo_ReportedPhish),
	}
}

func (s *TrainingAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingAssignedTo(input string) (*TrainingAssignedTo, error) {
	vals := map[string]TrainingAssignedTo{
		"allusers":          TrainingAssignedTo_AllUsers,
		"clickedpayload":    TrainingAssignedTo_ClickedPayload,
		"compromised":       TrainingAssignedTo_Compromised,
		"didnothing":        TrainingAssignedTo_DidNothing,
		"none":              TrainingAssignedTo_None,
		"readbutnotclicked": TrainingAssignedTo_ReadButNotClicked,
		"reportedphish":     TrainingAssignedTo_ReportedPhish,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingAssignedTo(input)
	return &out, nil
}
