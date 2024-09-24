package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationAlertType string

const (
	IosNotificationAlertType_Banner        IosNotificationAlertType = "banner"
	IosNotificationAlertType_DeviceDefault IosNotificationAlertType = "deviceDefault"
	IosNotificationAlertType_Modal         IosNotificationAlertType = "modal"
	IosNotificationAlertType_None          IosNotificationAlertType = "none"
)

func PossibleValuesForIosNotificationAlertType() []string {
	return []string{
		string(IosNotificationAlertType_Banner),
		string(IosNotificationAlertType_DeviceDefault),
		string(IosNotificationAlertType_Modal),
		string(IosNotificationAlertType_None),
	}
}

func (s *IosNotificationAlertType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosNotificationAlertType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosNotificationAlertType(input string) (*IosNotificationAlertType, error) {
	vals := map[string]IosNotificationAlertType{
		"banner":        IosNotificationAlertType_Banner,
		"devicedefault": IosNotificationAlertType_DeviceDefault,
		"modal":         IosNotificationAlertType_Modal,
		"none":          IosNotificationAlertType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosNotificationAlertType(input)
	return &out, nil
}
