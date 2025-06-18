package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MfaType string

const (
	MfaType_Certificate            MfaType = "certificate"
	MfaType_Eotp                   MfaType = "eotp"
	MfaType_Fido                   MfaType = "fido"
	MfaType_OneWaySms              MfaType = "oneWaySms"
	MfaType_Other                  MfaType = "other"
	MfaType_PhoneAppNotification   MfaType = "phoneAppNotification"
	MfaType_PhoneAppOtp            MfaType = "phoneAppOtp"
	MfaType_TwoWaySms              MfaType = "twoWaySms"
	MfaType_TwoWaySmsOtherMobile   MfaType = "twoWaySmsOtherMobile"
	MfaType_TwoWayVoiceMobile      MfaType = "twoWayVoiceMobile"
	MfaType_TwoWayVoiceOffice      MfaType = "twoWayVoiceOffice"
	MfaType_TwoWayVoiceOtherMobile MfaType = "twoWayVoiceOtherMobile"
)

func PossibleValuesForMfaType() []string {
	return []string{
		string(MfaType_Certificate),
		string(MfaType_Eotp),
		string(MfaType_Fido),
		string(MfaType_OneWaySms),
		string(MfaType_Other),
		string(MfaType_PhoneAppNotification),
		string(MfaType_PhoneAppOtp),
		string(MfaType_TwoWaySms),
		string(MfaType_TwoWaySmsOtherMobile),
		string(MfaType_TwoWayVoiceMobile),
		string(MfaType_TwoWayVoiceOffice),
		string(MfaType_TwoWayVoiceOtherMobile),
	}
}

func (s *MfaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMfaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMfaType(input string) (*MfaType, error) {
	vals := map[string]MfaType{
		"certificate":            MfaType_Certificate,
		"eotp":                   MfaType_Eotp,
		"fido":                   MfaType_Fido,
		"onewaysms":              MfaType_OneWaySms,
		"other":                  MfaType_Other,
		"phoneappnotification":   MfaType_PhoneAppNotification,
		"phoneappotp":            MfaType_PhoneAppOtp,
		"twowaysms":              MfaType_TwoWaySms,
		"twowaysmsothermobile":   MfaType_TwoWaySmsOtherMobile,
		"twowayvoicemobile":      MfaType_TwoWayVoiceMobile,
		"twowayvoiceoffice":      MfaType_TwoWayVoiceOffice,
		"twowayvoiceothermobile": MfaType_TwoWayVoiceOtherMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MfaType(input)
	return &out, nil
}
