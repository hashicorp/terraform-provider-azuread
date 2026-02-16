package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDefaultAuthenticationMethodType string

const (
	UserDefaultAuthenticationMethodType_Oath                 UserDefaultAuthenticationMethodType = "oath"
	UserDefaultAuthenticationMethodType_Push                 UserDefaultAuthenticationMethodType = "push"
	UserDefaultAuthenticationMethodType_Sms                  UserDefaultAuthenticationMethodType = "sms"
	UserDefaultAuthenticationMethodType_VoiceAlternateMobile UserDefaultAuthenticationMethodType = "voiceAlternateMobile"
	UserDefaultAuthenticationMethodType_VoiceMobile          UserDefaultAuthenticationMethodType = "voiceMobile"
	UserDefaultAuthenticationMethodType_VoiceOffice          UserDefaultAuthenticationMethodType = "voiceOffice"
)

func PossibleValuesForUserDefaultAuthenticationMethodType() []string {
	return []string{
		string(UserDefaultAuthenticationMethodType_Oath),
		string(UserDefaultAuthenticationMethodType_Push),
		string(UserDefaultAuthenticationMethodType_Sms),
		string(UserDefaultAuthenticationMethodType_VoiceAlternateMobile),
		string(UserDefaultAuthenticationMethodType_VoiceMobile),
		string(UserDefaultAuthenticationMethodType_VoiceOffice),
	}
}

func (s *UserDefaultAuthenticationMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserDefaultAuthenticationMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserDefaultAuthenticationMethodType(input string) (*UserDefaultAuthenticationMethodType, error) {
	vals := map[string]UserDefaultAuthenticationMethodType{
		"oath":                 UserDefaultAuthenticationMethodType_Oath,
		"push":                 UserDefaultAuthenticationMethodType_Push,
		"sms":                  UserDefaultAuthenticationMethodType_Sms,
		"voicealternatemobile": UserDefaultAuthenticationMethodType_VoiceAlternateMobile,
		"voicemobile":          UserDefaultAuthenticationMethodType_VoiceMobile,
		"voiceoffice":          UserDefaultAuthenticationMethodType_VoiceOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserDefaultAuthenticationMethodType(input)
	return &out, nil
}
