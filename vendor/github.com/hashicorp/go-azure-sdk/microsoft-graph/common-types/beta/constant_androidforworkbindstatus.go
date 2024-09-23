package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkBindStatus string

const (
	AndroidForWorkBindStatus_Bound             AndroidForWorkBindStatus = "bound"
	AndroidForWorkBindStatus_BoundAndValidated AndroidForWorkBindStatus = "boundAndValidated"
	AndroidForWorkBindStatus_NotBound          AndroidForWorkBindStatus = "notBound"
	AndroidForWorkBindStatus_Unbinding         AndroidForWorkBindStatus = "unbinding"
)

func PossibleValuesForAndroidForWorkBindStatus() []string {
	return []string{
		string(AndroidForWorkBindStatus_Bound),
		string(AndroidForWorkBindStatus_BoundAndValidated),
		string(AndroidForWorkBindStatus_NotBound),
		string(AndroidForWorkBindStatus_Unbinding),
	}
}

func (s *AndroidForWorkBindStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkBindStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkBindStatus(input string) (*AndroidForWorkBindStatus, error) {
	vals := map[string]AndroidForWorkBindStatus{
		"bound":             AndroidForWorkBindStatus_Bound,
		"boundandvalidated": AndroidForWorkBindStatus_BoundAndValidated,
		"notbound":          AndroidForWorkBindStatus_NotBound,
		"unbinding":         AndroidForWorkBindStatus_Unbinding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkBindStatus(input)
	return &out, nil
}
