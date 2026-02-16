package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryStatus string

const (
	AccessReviewHistoryStatus_Done       AccessReviewHistoryStatus = "done"
	AccessReviewHistoryStatus_Error      AccessReviewHistoryStatus = "error"
	AccessReviewHistoryStatus_Inprogress AccessReviewHistoryStatus = "inprogress"
	AccessReviewHistoryStatus_Requested  AccessReviewHistoryStatus = "requested"
)

func PossibleValuesForAccessReviewHistoryStatus() []string {
	return []string{
		string(AccessReviewHistoryStatus_Done),
		string(AccessReviewHistoryStatus_Error),
		string(AccessReviewHistoryStatus_Inprogress),
		string(AccessReviewHistoryStatus_Requested),
	}
}

func (s *AccessReviewHistoryStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewHistoryStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewHistoryStatus(input string) (*AccessReviewHistoryStatus, error) {
	vals := map[string]AccessReviewHistoryStatus{
		"done":       AccessReviewHistoryStatus_Done,
		"error":      AccessReviewHistoryStatus_Error,
		"inprogress": AccessReviewHistoryStatus_Inprogress,
		"requested":  AccessReviewHistoryStatus_Requested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryStatus(input)
	return &out, nil
}
