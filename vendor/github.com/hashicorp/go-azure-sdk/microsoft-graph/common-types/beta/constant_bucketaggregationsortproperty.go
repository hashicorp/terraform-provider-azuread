package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationSortProperty string

const (
	BucketAggregationSortProperty_Count       BucketAggregationSortProperty = "count"
	BucketAggregationSortProperty_KeyAsNumber BucketAggregationSortProperty = "keyAsNumber"
	BucketAggregationSortProperty_KeyAsString BucketAggregationSortProperty = "keyAsString"
)

func PossibleValuesForBucketAggregationSortProperty() []string {
	return []string{
		string(BucketAggregationSortProperty_Count),
		string(BucketAggregationSortProperty_KeyAsNumber),
		string(BucketAggregationSortProperty_KeyAsString),
	}
}

func (s *BucketAggregationSortProperty) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBucketAggregationSortProperty(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBucketAggregationSortProperty(input string) (*BucketAggregationSortProperty, error) {
	vals := map[string]BucketAggregationSortProperty{
		"count":       BucketAggregationSortProperty_Count,
		"keyasnumber": BucketAggregationSortProperty_KeyAsNumber,
		"keyasstring": BucketAggregationSortProperty_KeyAsString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BucketAggregationSortProperty(input)
	return &out, nil
}
