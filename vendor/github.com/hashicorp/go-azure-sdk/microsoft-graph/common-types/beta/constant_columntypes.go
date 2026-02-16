package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnTypes string

const (
	ColumnTypes_ApprovalStatus ColumnTypes = "approvalStatus"
	ColumnTypes_Boolean        ColumnTypes = "boolean"
	ColumnTypes_Calculated     ColumnTypes = "calculated"
	ColumnTypes_Choice         ColumnTypes = "choice"
	ColumnTypes_Currency       ColumnTypes = "currency"
	ColumnTypes_DateTime       ColumnTypes = "dateTime"
	ColumnTypes_Geolocation    ColumnTypes = "geolocation"
	ColumnTypes_Location       ColumnTypes = "location"
	ColumnTypes_Lookup         ColumnTypes = "lookup"
	ColumnTypes_Multichoice    ColumnTypes = "multichoice"
	ColumnTypes_Multiterm      ColumnTypes = "multiterm"
	ColumnTypes_Note           ColumnTypes = "note"
	ColumnTypes_Number         ColumnTypes = "number"
	ColumnTypes_Term           ColumnTypes = "term"
	ColumnTypes_Text           ColumnTypes = "text"
	ColumnTypes_Thumbnail      ColumnTypes = "thumbnail"
	ColumnTypes_Url            ColumnTypes = "url"
	ColumnTypes_User           ColumnTypes = "user"
)

func PossibleValuesForColumnTypes() []string {
	return []string{
		string(ColumnTypes_ApprovalStatus),
		string(ColumnTypes_Boolean),
		string(ColumnTypes_Calculated),
		string(ColumnTypes_Choice),
		string(ColumnTypes_Currency),
		string(ColumnTypes_DateTime),
		string(ColumnTypes_Geolocation),
		string(ColumnTypes_Location),
		string(ColumnTypes_Lookup),
		string(ColumnTypes_Multichoice),
		string(ColumnTypes_Multiterm),
		string(ColumnTypes_Note),
		string(ColumnTypes_Number),
		string(ColumnTypes_Term),
		string(ColumnTypes_Text),
		string(ColumnTypes_Thumbnail),
		string(ColumnTypes_Url),
		string(ColumnTypes_User),
	}
}

func (s *ColumnTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseColumnTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseColumnTypes(input string) (*ColumnTypes, error) {
	vals := map[string]ColumnTypes{
		"approvalstatus": ColumnTypes_ApprovalStatus,
		"boolean":        ColumnTypes_Boolean,
		"calculated":     ColumnTypes_Calculated,
		"choice":         ColumnTypes_Choice,
		"currency":       ColumnTypes_Currency,
		"datetime":       ColumnTypes_DateTime,
		"geolocation":    ColumnTypes_Geolocation,
		"location":       ColumnTypes_Location,
		"lookup":         ColumnTypes_Lookup,
		"multichoice":    ColumnTypes_Multichoice,
		"multiterm":      ColumnTypes_Multiterm,
		"note":           ColumnTypes_Note,
		"number":         ColumnTypes_Number,
		"term":           ColumnTypes_Term,
		"text":           ColumnTypes_Text,
		"thumbnail":      ColumnTypes_Thumbnail,
		"url":            ColumnTypes_Url,
		"user":           ColumnTypes_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ColumnTypes(input)
	return &out, nil
}
