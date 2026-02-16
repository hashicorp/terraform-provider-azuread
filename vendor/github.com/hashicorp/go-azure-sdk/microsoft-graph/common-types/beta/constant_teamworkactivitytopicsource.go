package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkActivityTopicSource string

const (
	TeamworkActivityTopicSource_EntityUrl TeamworkActivityTopicSource = "entityUrl"
	TeamworkActivityTopicSource_Text      TeamworkActivityTopicSource = "text"
)

func PossibleValuesForTeamworkActivityTopicSource() []string {
	return []string{
		string(TeamworkActivityTopicSource_EntityUrl),
		string(TeamworkActivityTopicSource_Text),
	}
}

func (s *TeamworkActivityTopicSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkActivityTopicSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkActivityTopicSource(input string) (*TeamworkActivityTopicSource, error) {
	vals := map[string]TeamworkActivityTopicSource{
		"entityurl": TeamworkActivityTopicSource_EntityUrl,
		"text":      TeamworkActivityTopicSource_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkActivityTopicSource(input)
	return &out, nil
}
