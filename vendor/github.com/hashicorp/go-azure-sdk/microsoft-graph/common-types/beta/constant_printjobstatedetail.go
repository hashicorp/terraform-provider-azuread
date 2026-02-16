package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStateDetail string

const (
	PrintJobStateDetail_CompletedSuccessfully PrintJobStateDetail = "completedSuccessfully"
	PrintJobStateDetail_CompletedWithErrors   PrintJobStateDetail = "completedWithErrors"
	PrintJobStateDetail_CompletedWithWarnings PrintJobStateDetail = "completedWithWarnings"
	PrintJobStateDetail_Interpreting          PrintJobStateDetail = "interpreting"
	PrintJobStateDetail_ReleaseWait           PrintJobStateDetail = "releaseWait"
	PrintJobStateDetail_Transforming          PrintJobStateDetail = "transforming"
	PrintJobStateDetail_UploadPending         PrintJobStateDetail = "uploadPending"
)

func PossibleValuesForPrintJobStateDetail() []string {
	return []string{
		string(PrintJobStateDetail_CompletedSuccessfully),
		string(PrintJobStateDetail_CompletedWithErrors),
		string(PrintJobStateDetail_CompletedWithWarnings),
		string(PrintJobStateDetail_Interpreting),
		string(PrintJobStateDetail_ReleaseWait),
		string(PrintJobStateDetail_Transforming),
		string(PrintJobStateDetail_UploadPending),
	}
}

func (s *PrintJobStateDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobStateDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobStateDetail(input string) (*PrintJobStateDetail, error) {
	vals := map[string]PrintJobStateDetail{
		"completedsuccessfully": PrintJobStateDetail_CompletedSuccessfully,
		"completedwitherrors":   PrintJobStateDetail_CompletedWithErrors,
		"completedwithwarnings": PrintJobStateDetail_CompletedWithWarnings,
		"interpreting":          PrintJobStateDetail_Interpreting,
		"releasewait":           PrintJobStateDetail_ReleaseWait,
		"transforming":          PrintJobStateDetail_Transforming,
		"uploadpending":         PrintJobStateDetail_UploadPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStateDetail(input)
	return &out, nil
}
