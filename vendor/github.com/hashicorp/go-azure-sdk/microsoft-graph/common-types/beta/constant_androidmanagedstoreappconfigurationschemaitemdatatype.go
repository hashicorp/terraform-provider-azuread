package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationSchemaItemDataType string

const (
	AndroidManagedStoreAppConfigurationSchemaItemDataType_Bool        AndroidManagedStoreAppConfigurationSchemaItemDataType = "bool"
	AndroidManagedStoreAppConfigurationSchemaItemDataType_Bundle      AndroidManagedStoreAppConfigurationSchemaItemDataType = "bundle"
	AndroidManagedStoreAppConfigurationSchemaItemDataType_BundleArray AndroidManagedStoreAppConfigurationSchemaItemDataType = "bundleArray"
	AndroidManagedStoreAppConfigurationSchemaItemDataType_Choice      AndroidManagedStoreAppConfigurationSchemaItemDataType = "choice"
	AndroidManagedStoreAppConfigurationSchemaItemDataType_Hidden      AndroidManagedStoreAppConfigurationSchemaItemDataType = "hidden"
	AndroidManagedStoreAppConfigurationSchemaItemDataType_Integer     AndroidManagedStoreAppConfigurationSchemaItemDataType = "integer"
	AndroidManagedStoreAppConfigurationSchemaItemDataType_Multiselect AndroidManagedStoreAppConfigurationSchemaItemDataType = "multiselect"
	AndroidManagedStoreAppConfigurationSchemaItemDataType_String      AndroidManagedStoreAppConfigurationSchemaItemDataType = "string"
)

func PossibleValuesForAndroidManagedStoreAppConfigurationSchemaItemDataType() []string {
	return []string{
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_Bool),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_Bundle),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_BundleArray),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_Choice),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_Hidden),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_Integer),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_Multiselect),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataType_String),
	}
}

func (s *AndroidManagedStoreAppConfigurationSchemaItemDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAppConfigurationSchemaItemDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAppConfigurationSchemaItemDataType(input string) (*AndroidManagedStoreAppConfigurationSchemaItemDataType, error) {
	vals := map[string]AndroidManagedStoreAppConfigurationSchemaItemDataType{
		"bool":        AndroidManagedStoreAppConfigurationSchemaItemDataType_Bool,
		"bundle":      AndroidManagedStoreAppConfigurationSchemaItemDataType_Bundle,
		"bundlearray": AndroidManagedStoreAppConfigurationSchemaItemDataType_BundleArray,
		"choice":      AndroidManagedStoreAppConfigurationSchemaItemDataType_Choice,
		"hidden":      AndroidManagedStoreAppConfigurationSchemaItemDataType_Hidden,
		"integer":     AndroidManagedStoreAppConfigurationSchemaItemDataType_Integer,
		"multiselect": AndroidManagedStoreAppConfigurationSchemaItemDataType_Multiselect,
		"string":      AndroidManagedStoreAppConfigurationSchemaItemDataType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppConfigurationSchemaItemDataType(input)
	return &out, nil
}
