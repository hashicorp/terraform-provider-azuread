package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleSummaryStatus string

const (
	RoleSummaryStatus_Bad RoleSummaryStatus = "bad"
	RoleSummaryStatus_Ok  RoleSummaryStatus = "ok"
)

func PossibleValuesForRoleSummaryStatus() []string {
	return []string{
		string(RoleSummaryStatus_Bad),
		string(RoleSummaryStatus_Ok),
	}
}

func (s *RoleSummaryStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleSummaryStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleSummaryStatus(input string) (*RoleSummaryStatus, error) {
	vals := map[string]RoleSummaryStatus{
		"bad": RoleSummaryStatus_Bad,
		"ok":  RoleSummaryStatus_Ok,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleSummaryStatus(input)
	return &out, nil
}
