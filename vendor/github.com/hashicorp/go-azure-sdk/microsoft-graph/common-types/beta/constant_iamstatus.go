package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IamStatus string

const (
	IamStatus_Active   IamStatus = "active"
	IamStatus_Disabled IamStatus = "disabled"
	IamStatus_Inactive IamStatus = "inactive"
)

func PossibleValuesForIamStatus() []string {
	return []string{
		string(IamStatus_Active),
		string(IamStatus_Disabled),
		string(IamStatus_Inactive),
	}
}

func (s *IamStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIamStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIamStatus(input string) (*IamStatus, error) {
	vals := map[string]IamStatus{
		"active":   IamStatus_Active,
		"disabled": IamStatus_Disabled,
		"inactive": IamStatus_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IamStatus(input)
	return &out, nil
}
