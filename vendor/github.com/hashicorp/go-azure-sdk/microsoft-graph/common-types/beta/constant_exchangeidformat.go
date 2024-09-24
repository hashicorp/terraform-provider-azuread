package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeIdFormat string

const (
	ExchangeIdFormat_EntryId              ExchangeIdFormat = "entryId"
	ExchangeIdFormat_EwsId                ExchangeIdFormat = "ewsId"
	ExchangeIdFormat_ImmutableEntryId     ExchangeIdFormat = "immutableEntryId"
	ExchangeIdFormat_RestId               ExchangeIdFormat = "restId"
	ExchangeIdFormat_RestImmutableEntryId ExchangeIdFormat = "restImmutableEntryId"
)

func PossibleValuesForExchangeIdFormat() []string {
	return []string{
		string(ExchangeIdFormat_EntryId),
		string(ExchangeIdFormat_EwsId),
		string(ExchangeIdFormat_ImmutableEntryId),
		string(ExchangeIdFormat_RestId),
		string(ExchangeIdFormat_RestImmutableEntryId),
	}
}

func (s *ExchangeIdFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExchangeIdFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExchangeIdFormat(input string) (*ExchangeIdFormat, error) {
	vals := map[string]ExchangeIdFormat{
		"entryid":              ExchangeIdFormat_EntryId,
		"ewsid":                ExchangeIdFormat_EwsId,
		"immutableentryid":     ExchangeIdFormat_ImmutableEntryId,
		"restid":               ExchangeIdFormat_RestId,
		"restimmutableentryid": ExchangeIdFormat_RestImmutableEntryId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExchangeIdFormat(input)
	return &out, nil
}
