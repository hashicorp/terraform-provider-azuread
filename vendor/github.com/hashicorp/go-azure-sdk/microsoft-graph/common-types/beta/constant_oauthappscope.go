package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OAuthAppScope string

const (
	OAuthAppScope_ReadAllChat      OAuthAppScope = "readAllChat"
	OAuthAppScope_ReadAllFile      OAuthAppScope = "readAllFile"
	OAuthAppScope_ReadAndWriteMail OAuthAppScope = "readAndWriteMail"
	OAuthAppScope_ReadCalendar     OAuthAppScope = "readCalendar"
	OAuthAppScope_ReadContact      OAuthAppScope = "readContact"
	OAuthAppScope_ReadMail         OAuthAppScope = "readMail"
	OAuthAppScope_SendMail         OAuthAppScope = "sendMail"
	OAuthAppScope_Unknown          OAuthAppScope = "unknown"
)

func PossibleValuesForOAuthAppScope() []string {
	return []string{
		string(OAuthAppScope_ReadAllChat),
		string(OAuthAppScope_ReadAllFile),
		string(OAuthAppScope_ReadAndWriteMail),
		string(OAuthAppScope_ReadCalendar),
		string(OAuthAppScope_ReadContact),
		string(OAuthAppScope_ReadMail),
		string(OAuthAppScope_SendMail),
		string(OAuthAppScope_Unknown),
	}
}

func (s *OAuthAppScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOAuthAppScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOAuthAppScope(input string) (*OAuthAppScope, error) {
	vals := map[string]OAuthAppScope{
		"readallchat":      OAuthAppScope_ReadAllChat,
		"readallfile":      OAuthAppScope_ReadAllFile,
		"readandwritemail": OAuthAppScope_ReadAndWriteMail,
		"readcalendar":     OAuthAppScope_ReadCalendar,
		"readcontact":      OAuthAppScope_ReadContact,
		"readmail":         OAuthAppScope_ReadMail,
		"sendmail":         OAuthAppScope_SendMail,
		"unknown":          OAuthAppScope_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OAuthAppScope(input)
	return &out, nil
}
