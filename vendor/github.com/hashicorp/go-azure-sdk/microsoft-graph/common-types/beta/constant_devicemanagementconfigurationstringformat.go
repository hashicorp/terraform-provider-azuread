package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringFormat string

const (
	DeviceManagementConfigurationStringFormat_Base64     DeviceManagementConfigurationStringFormat = "base64"
	DeviceManagementConfigurationStringFormat_BashScript DeviceManagementConfigurationStringFormat = "bashScript"
	DeviceManagementConfigurationStringFormat_Binary     DeviceManagementConfigurationStringFormat = "binary"
	DeviceManagementConfigurationStringFormat_Date       DeviceManagementConfigurationStringFormat = "date"
	DeviceManagementConfigurationStringFormat_DateTime   DeviceManagementConfigurationStringFormat = "dateTime"
	DeviceManagementConfigurationStringFormat_Email      DeviceManagementConfigurationStringFormat = "email"
	DeviceManagementConfigurationStringFormat_Guid       DeviceManagementConfigurationStringFormat = "guid"
	DeviceManagementConfigurationStringFormat_Ip         DeviceManagementConfigurationStringFormat = "ip"
	DeviceManagementConfigurationStringFormat_Json       DeviceManagementConfigurationStringFormat = "json"
	DeviceManagementConfigurationStringFormat_None       DeviceManagementConfigurationStringFormat = "none"
	DeviceManagementConfigurationStringFormat_RegEx      DeviceManagementConfigurationStringFormat = "regEx"
	DeviceManagementConfigurationStringFormat_SurfaceHub DeviceManagementConfigurationStringFormat = "surfaceHub"
	DeviceManagementConfigurationStringFormat_Time       DeviceManagementConfigurationStringFormat = "time"
	DeviceManagementConfigurationStringFormat_Url        DeviceManagementConfigurationStringFormat = "url"
	DeviceManagementConfigurationStringFormat_Version    DeviceManagementConfigurationStringFormat = "version"
	DeviceManagementConfigurationStringFormat_Xml        DeviceManagementConfigurationStringFormat = "xml"
)

func PossibleValuesForDeviceManagementConfigurationStringFormat() []string {
	return []string{
		string(DeviceManagementConfigurationStringFormat_Base64),
		string(DeviceManagementConfigurationStringFormat_BashScript),
		string(DeviceManagementConfigurationStringFormat_Binary),
		string(DeviceManagementConfigurationStringFormat_Date),
		string(DeviceManagementConfigurationStringFormat_DateTime),
		string(DeviceManagementConfigurationStringFormat_Email),
		string(DeviceManagementConfigurationStringFormat_Guid),
		string(DeviceManagementConfigurationStringFormat_Ip),
		string(DeviceManagementConfigurationStringFormat_Json),
		string(DeviceManagementConfigurationStringFormat_None),
		string(DeviceManagementConfigurationStringFormat_RegEx),
		string(DeviceManagementConfigurationStringFormat_SurfaceHub),
		string(DeviceManagementConfigurationStringFormat_Time),
		string(DeviceManagementConfigurationStringFormat_Url),
		string(DeviceManagementConfigurationStringFormat_Version),
		string(DeviceManagementConfigurationStringFormat_Xml),
	}
}

func (s *DeviceManagementConfigurationStringFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationStringFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationStringFormat(input string) (*DeviceManagementConfigurationStringFormat, error) {
	vals := map[string]DeviceManagementConfigurationStringFormat{
		"base64":     DeviceManagementConfigurationStringFormat_Base64,
		"bashscript": DeviceManagementConfigurationStringFormat_BashScript,
		"binary":     DeviceManagementConfigurationStringFormat_Binary,
		"date":       DeviceManagementConfigurationStringFormat_Date,
		"datetime":   DeviceManagementConfigurationStringFormat_DateTime,
		"email":      DeviceManagementConfigurationStringFormat_Email,
		"guid":       DeviceManagementConfigurationStringFormat_Guid,
		"ip":         DeviceManagementConfigurationStringFormat_Ip,
		"json":       DeviceManagementConfigurationStringFormat_Json,
		"none":       DeviceManagementConfigurationStringFormat_None,
		"regex":      DeviceManagementConfigurationStringFormat_RegEx,
		"surfacehub": DeviceManagementConfigurationStringFormat_SurfaceHub,
		"time":       DeviceManagementConfigurationStringFormat_Time,
		"url":        DeviceManagementConfigurationStringFormat_Url,
		"version":    DeviceManagementConfigurationStringFormat_Version,
		"xml":        DeviceManagementConfigurationStringFormat_Xml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationStringFormat(input)
	return &out, nil
}
