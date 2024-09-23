package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPreviewType string

const (
	PlannerPreviewType_Automatic   PlannerPreviewType = "automatic"
	PlannerPreviewType_Checklist   PlannerPreviewType = "checklist"
	PlannerPreviewType_Description PlannerPreviewType = "description"
	PlannerPreviewType_NoPreview   PlannerPreviewType = "noPreview"
	PlannerPreviewType_Reference   PlannerPreviewType = "reference"
)

func PossibleValuesForPlannerPreviewType() []string {
	return []string{
		string(PlannerPreviewType_Automatic),
		string(PlannerPreviewType_Checklist),
		string(PlannerPreviewType_Description),
		string(PlannerPreviewType_NoPreview),
		string(PlannerPreviewType_Reference),
	}
}

func (s *PlannerPreviewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPreviewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPreviewType(input string) (*PlannerPreviewType, error) {
	vals := map[string]PlannerPreviewType{
		"automatic":   PlannerPreviewType_Automatic,
		"checklist":   PlannerPreviewType_Checklist,
		"description": PlannerPreviewType_Description,
		"nopreview":   PlannerPreviewType_NoPreview,
		"reference":   PlannerPreviewType_Reference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPreviewType(input)
	return &out, nil
}
