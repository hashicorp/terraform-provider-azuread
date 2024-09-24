package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseAction string

const (
	SecurityCaseAction_AddToReviewSet     SecurityCaseAction = "addToReviewSet"
	SecurityCaseAction_ApplyTags          SecurityCaseAction = "applyTags"
	SecurityCaseAction_ContentExport      SecurityCaseAction = "contentExport"
	SecurityCaseAction_ConvertToPdf       SecurityCaseAction = "convertToPdf"
	SecurityCaseAction_EstimateStatistics SecurityCaseAction = "estimateStatistics"
	SecurityCaseAction_ExportReport       SecurityCaseAction = "exportReport"
	SecurityCaseAction_ExportResult       SecurityCaseAction = "exportResult"
	SecurityCaseAction_HoldUpdate         SecurityCaseAction = "holdUpdate"
	SecurityCaseAction_Index              SecurityCaseAction = "index"
	SecurityCaseAction_PurgeData          SecurityCaseAction = "purgeData"
)

func PossibleValuesForSecurityCaseAction() []string {
	return []string{
		string(SecurityCaseAction_AddToReviewSet),
		string(SecurityCaseAction_ApplyTags),
		string(SecurityCaseAction_ContentExport),
		string(SecurityCaseAction_ConvertToPdf),
		string(SecurityCaseAction_EstimateStatistics),
		string(SecurityCaseAction_ExportReport),
		string(SecurityCaseAction_ExportResult),
		string(SecurityCaseAction_HoldUpdate),
		string(SecurityCaseAction_Index),
		string(SecurityCaseAction_PurgeData),
	}
}

func (s *SecurityCaseAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCaseAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCaseAction(input string) (*SecurityCaseAction, error) {
	vals := map[string]SecurityCaseAction{
		"addtoreviewset":     SecurityCaseAction_AddToReviewSet,
		"applytags":          SecurityCaseAction_ApplyTags,
		"contentexport":      SecurityCaseAction_ContentExport,
		"converttopdf":       SecurityCaseAction_ConvertToPdf,
		"estimatestatistics": SecurityCaseAction_EstimateStatistics,
		"exportreport":       SecurityCaseAction_ExportReport,
		"exportresult":       SecurityCaseAction_ExportResult,
		"holdupdate":         SecurityCaseAction_HoldUpdate,
		"index":              SecurityCaseAction_Index,
		"purgedata":          SecurityCaseAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCaseAction(input)
	return &out, nil
}
