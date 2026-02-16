package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseAction string

const (
	EdiscoveryCaseAction_AddToReviewSet     EdiscoveryCaseAction = "addToReviewSet"
	EdiscoveryCaseAction_ApplyTags          EdiscoveryCaseAction = "applyTags"
	EdiscoveryCaseAction_ContentExport      EdiscoveryCaseAction = "contentExport"
	EdiscoveryCaseAction_ConvertToPdf       EdiscoveryCaseAction = "convertToPdf"
	EdiscoveryCaseAction_EstimateStatistics EdiscoveryCaseAction = "estimateStatistics"
	EdiscoveryCaseAction_HoldUpdate         EdiscoveryCaseAction = "holdUpdate"
	EdiscoveryCaseAction_Index              EdiscoveryCaseAction = "index"
	EdiscoveryCaseAction_PurgeData          EdiscoveryCaseAction = "purgeData"
)

func PossibleValuesForEdiscoveryCaseAction() []string {
	return []string{
		string(EdiscoveryCaseAction_AddToReviewSet),
		string(EdiscoveryCaseAction_ApplyTags),
		string(EdiscoveryCaseAction_ContentExport),
		string(EdiscoveryCaseAction_ConvertToPdf),
		string(EdiscoveryCaseAction_EstimateStatistics),
		string(EdiscoveryCaseAction_HoldUpdate),
		string(EdiscoveryCaseAction_Index),
		string(EdiscoveryCaseAction_PurgeData),
	}
}

func (s *EdiscoveryCaseAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseAction(input string) (*EdiscoveryCaseAction, error) {
	vals := map[string]EdiscoveryCaseAction{
		"addtoreviewset":     EdiscoveryCaseAction_AddToReviewSet,
		"applytags":          EdiscoveryCaseAction_ApplyTags,
		"contentexport":      EdiscoveryCaseAction_ContentExport,
		"converttopdf":       EdiscoveryCaseAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryCaseAction_EstimateStatistics,
		"holdupdate":         EdiscoveryCaseAction_HoldUpdate,
		"index":              EdiscoveryCaseAction_Index,
		"purgedata":          EdiscoveryCaseAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseAction(input)
	return &out, nil
}
