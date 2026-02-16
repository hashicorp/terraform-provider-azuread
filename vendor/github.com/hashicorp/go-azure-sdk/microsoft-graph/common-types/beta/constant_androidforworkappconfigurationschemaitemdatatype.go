package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppConfigurationSchemaItemDataType string

const (
	AndroidForWorkAppConfigurationSchemaItemDataType_Bool        AndroidForWorkAppConfigurationSchemaItemDataType = "bool"
	AndroidForWorkAppConfigurationSchemaItemDataType_Bundle      AndroidForWorkAppConfigurationSchemaItemDataType = "bundle"
	AndroidForWorkAppConfigurationSchemaItemDataType_BundleArray AndroidForWorkAppConfigurationSchemaItemDataType = "bundleArray"
	AndroidForWorkAppConfigurationSchemaItemDataType_Choice      AndroidForWorkAppConfigurationSchemaItemDataType = "choice"
	AndroidForWorkAppConfigurationSchemaItemDataType_Hidden      AndroidForWorkAppConfigurationSchemaItemDataType = "hidden"
	AndroidForWorkAppConfigurationSchemaItemDataType_Integer     AndroidForWorkAppConfigurationSchemaItemDataType = "integer"
	AndroidForWorkAppConfigurationSchemaItemDataType_Multiselect AndroidForWorkAppConfigurationSchemaItemDataType = "multiselect"
	AndroidForWorkAppConfigurationSchemaItemDataType_String      AndroidForWorkAppConfigurationSchemaItemDataType = "string"
)

func PossibleValuesForAndroidForWorkAppConfigurationSchemaItemDataType() []string {
	return []string{
		string(AndroidForWorkAppConfigurationSchemaItemDataType_Bool),
		string(AndroidForWorkAppConfigurationSchemaItemDataType_Bundle),
		string(AndroidForWorkAppConfigurationSchemaItemDataType_BundleArray),
		string(AndroidForWorkAppConfigurationSchemaItemDataType_Choice),
		string(AndroidForWorkAppConfigurationSchemaItemDataType_Hidden),
		string(AndroidForWorkAppConfigurationSchemaItemDataType_Integer),
		string(AndroidForWorkAppConfigurationSchemaItemDataType_Multiselect),
		string(AndroidForWorkAppConfigurationSchemaItemDataType_String),
	}
}

func (s *AndroidForWorkAppConfigurationSchemaItemDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkAppConfigurationSchemaItemDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkAppConfigurationSchemaItemDataType(input string) (*AndroidForWorkAppConfigurationSchemaItemDataType, error) {
	vals := map[string]AndroidForWorkAppConfigurationSchemaItemDataType{
		"bool":        AndroidForWorkAppConfigurationSchemaItemDataType_Bool,
		"bundle":      AndroidForWorkAppConfigurationSchemaItemDataType_Bundle,
		"bundlearray": AndroidForWorkAppConfigurationSchemaItemDataType_BundleArray,
		"choice":      AndroidForWorkAppConfigurationSchemaItemDataType_Choice,
		"hidden":      AndroidForWorkAppConfigurationSchemaItemDataType_Hidden,
		"integer":     AndroidForWorkAppConfigurationSchemaItemDataType_Integer,
		"multiselect": AndroidForWorkAppConfigurationSchemaItemDataType_Multiselect,
		"string":      AndroidForWorkAppConfigurationSchemaItemDataType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkAppConfigurationSchemaItemDataType(input)
	return &out, nil
}
