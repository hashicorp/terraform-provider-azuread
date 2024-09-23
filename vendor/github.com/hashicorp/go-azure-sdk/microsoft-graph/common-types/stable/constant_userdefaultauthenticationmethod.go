package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDefaultAuthenticationMethod string

const (
	UserDefaultAuthenticationMethod_None                 UserDefaultAuthenticationMethod = "none"
	UserDefaultAuthenticationMethod_Oath                 UserDefaultAuthenticationMethod = "oath"
	UserDefaultAuthenticationMethod_Push                 UserDefaultAuthenticationMethod = "push"
	UserDefaultAuthenticationMethod_Sms                  UserDefaultAuthenticationMethod = "sms"
	UserDefaultAuthenticationMethod_VoiceAlternateMobile UserDefaultAuthenticationMethod = "voiceAlternateMobile"
	UserDefaultAuthenticationMethod_VoiceMobile          UserDefaultAuthenticationMethod = "voiceMobile"
	UserDefaultAuthenticationMethod_VoiceOffice          UserDefaultAuthenticationMethod = "voiceOffice"
)

func PossibleValuesForUserDefaultAuthenticationMethod() []string {
	return []string{
		string(UserDefaultAuthenticationMethod_None),
		string(UserDefaultAuthenticationMethod_Oath),
		string(UserDefaultAuthenticationMethod_Push),
		string(UserDefaultAuthenticationMethod_Sms),
		string(UserDefaultAuthenticationMethod_VoiceAlternateMobile),
		string(UserDefaultAuthenticationMethod_VoiceMobile),
		string(UserDefaultAuthenticationMethod_VoiceOffice),
	}
}

func (s *UserDefaultAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserDefaultAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserDefaultAuthenticationMethod(input string) (*UserDefaultAuthenticationMethod, error) {
	vals := map[string]UserDefaultAuthenticationMethod{
		"none":                 UserDefaultAuthenticationMethod_None,
		"oath":                 UserDefaultAuthenticationMethod_Oath,
		"push":                 UserDefaultAuthenticationMethod_Push,
		"sms":                  UserDefaultAuthenticationMethod_Sms,
		"voicealternatemobile": UserDefaultAuthenticationMethod_VoiceAlternateMobile,
		"voicemobile":          UserDefaultAuthenticationMethod_VoiceMobile,
		"voiceoffice":          UserDefaultAuthenticationMethod_VoiceOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserDefaultAuthenticationMethod(input)
	return &out, nil
}
