package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAction string

const (
	UserAction_RegisterOrJoinDevices       UserAction = "registerOrJoinDevices"
	UserAction_RegisterSecurityInformation UserAction = "registerSecurityInformation"
)

func PossibleValuesForUserAction() []string {
	return []string{
		string(UserAction_RegisterOrJoinDevices),
		string(UserAction_RegisterSecurityInformation),
	}
}

func (s *UserAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserAction(input string) (*UserAction, error) {
	vals := map[string]UserAction{
		"registerorjoindevices":       UserAction_RegisterOrJoinDevices,
		"registersecurityinformation": UserAction_RegisterSecurityInformation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserAction(input)
	return &out, nil
}
