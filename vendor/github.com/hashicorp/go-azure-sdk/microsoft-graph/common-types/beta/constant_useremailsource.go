package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEmailSource string

const (
	UserEmailSource_PrimarySmtpAddress UserEmailSource = "primarySmtpAddress"
	UserEmailSource_UserPrincipalName  UserEmailSource = "userPrincipalName"
)

func PossibleValuesForUserEmailSource() []string {
	return []string{
		string(UserEmailSource_PrimarySmtpAddress),
		string(UserEmailSource_UserPrincipalName),
	}
}

func (s *UserEmailSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserEmailSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserEmailSource(input string) (*UserEmailSource, error) {
	vals := map[string]UserEmailSource{
		"primarysmtpaddress": UserEmailSource_PrimarySmtpAddress,
		"userprincipalname":  UserEmailSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserEmailSource(input)
	return &out, nil
}
