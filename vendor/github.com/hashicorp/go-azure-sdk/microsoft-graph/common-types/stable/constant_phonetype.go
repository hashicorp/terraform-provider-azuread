package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneType string

const (
	PhoneType_Assistant   PhoneType = "assistant"
	PhoneType_Business    PhoneType = "business"
	PhoneType_BusinessFax PhoneType = "businessFax"
	PhoneType_Home        PhoneType = "home"
	PhoneType_HomeFax     PhoneType = "homeFax"
	PhoneType_Mobile      PhoneType = "mobile"
	PhoneType_Other       PhoneType = "other"
	PhoneType_OtherFax    PhoneType = "otherFax"
	PhoneType_Pager       PhoneType = "pager"
	PhoneType_Radio       PhoneType = "radio"
)

func PossibleValuesForPhoneType() []string {
	return []string{
		string(PhoneType_Assistant),
		string(PhoneType_Business),
		string(PhoneType_BusinessFax),
		string(PhoneType_Home),
		string(PhoneType_HomeFax),
		string(PhoneType_Mobile),
		string(PhoneType_Other),
		string(PhoneType_OtherFax),
		string(PhoneType_Pager),
		string(PhoneType_Radio),
	}
}

func (s *PhoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhoneType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhoneType(input string) (*PhoneType, error) {
	vals := map[string]PhoneType{
		"assistant":   PhoneType_Assistant,
		"business":    PhoneType_Business,
		"businessfax": PhoneType_BusinessFax,
		"home":        PhoneType_Home,
		"homefax":     PhoneType_HomeFax,
		"mobile":      PhoneType_Mobile,
		"other":       PhoneType_Other,
		"otherfax":    PhoneType_OtherFax,
		"pager":       PhoneType_Pager,
		"radio":       PhoneType_Radio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneType(input)
	return &out, nil
}
