package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidUsernameSource string

const (
	AndroidUsernameSource_PrimarySmtpAddress AndroidUsernameSource = "primarySmtpAddress"
	AndroidUsernameSource_SamAccountName     AndroidUsernameSource = "samAccountName"
	AndroidUsernameSource_UserPrincipalName  AndroidUsernameSource = "userPrincipalName"
	AndroidUsernameSource_Username           AndroidUsernameSource = "username"
)

func PossibleValuesForAndroidUsernameSource() []string {
	return []string{
		string(AndroidUsernameSource_PrimarySmtpAddress),
		string(AndroidUsernameSource_SamAccountName),
		string(AndroidUsernameSource_UserPrincipalName),
		string(AndroidUsernameSource_Username),
	}
}

func (s *AndroidUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidUsernameSource(input string) (*AndroidUsernameSource, error) {
	vals := map[string]AndroidUsernameSource{
		"primarysmtpaddress": AndroidUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidUsernameSource_SamAccountName,
		"userprincipalname":  AndroidUsernameSource_UserPrincipalName,
		"username":           AndroidUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidUsernameSource(input)
	return &out, nil
}
