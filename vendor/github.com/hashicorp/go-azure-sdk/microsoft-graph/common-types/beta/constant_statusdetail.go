package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusDetail string

const (
	StatusDetail_Approved  StatusDetail = "approved"
	StatusDetail_Canceled  StatusDetail = "canceled"
	StatusDetail_Completed StatusDetail = "completed"
	StatusDetail_Rejected  StatusDetail = "rejected"
	StatusDetail_Submitted StatusDetail = "submitted"
)

func PossibleValuesForStatusDetail() []string {
	return []string{
		string(StatusDetail_Approved),
		string(StatusDetail_Canceled),
		string(StatusDetail_Completed),
		string(StatusDetail_Rejected),
		string(StatusDetail_Submitted),
	}
}

func (s *StatusDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusDetail(input string) (*StatusDetail, error) {
	vals := map[string]StatusDetail{
		"approved":  StatusDetail_Approved,
		"canceled":  StatusDetail_Canceled,
		"completed": StatusDetail_Completed,
		"rejected":  StatusDetail_Rejected,
		"submitted": StatusDetail_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusDetail(input)
	return &out, nil
}
