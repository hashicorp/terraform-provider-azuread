package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsernameSource string

const (
	UsernameSource_PrimarySmtpAddress UsernameSource = "primarySmtpAddress"
	UsernameSource_SamAccountName     UsernameSource = "samAccountName"
	UsernameSource_UserPrincipalName  UsernameSource = "userPrincipalName"
)

func PossibleValuesForUsernameSource() []string {
	return []string{
		string(UsernameSource_PrimarySmtpAddress),
		string(UsernameSource_SamAccountName),
		string(UsernameSource_UserPrincipalName),
	}
}

func (s *UsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUsernameSource(input string) (*UsernameSource, error) {
	vals := map[string]UsernameSource{
		"primarysmtpaddress": UsernameSource_PrimarySmtpAddress,
		"samaccountname":     UsernameSource_SamAccountName,
		"userprincipalname":  UsernameSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsernameSource(input)
	return &out, nil
}
