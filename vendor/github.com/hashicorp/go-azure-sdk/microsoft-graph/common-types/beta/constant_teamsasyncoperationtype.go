package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAsyncOperationType string

const (
	TeamsAsyncOperationType_ArchiveChannel   TeamsAsyncOperationType = "archiveChannel"
	TeamsAsyncOperationType_ArchiveTeam      TeamsAsyncOperationType = "archiveTeam"
	TeamsAsyncOperationType_CloneTeam        TeamsAsyncOperationType = "cloneTeam"
	TeamsAsyncOperationType_CreateChannel    TeamsAsyncOperationType = "createChannel"
	TeamsAsyncOperationType_CreateChat       TeamsAsyncOperationType = "createChat"
	TeamsAsyncOperationType_CreateTeam       TeamsAsyncOperationType = "createTeam"
	TeamsAsyncOperationType_Invalid          TeamsAsyncOperationType = "invalid"
	TeamsAsyncOperationType_TeamifyGroup     TeamsAsyncOperationType = "teamifyGroup"
	TeamsAsyncOperationType_UnarchiveChannel TeamsAsyncOperationType = "unarchiveChannel"
	TeamsAsyncOperationType_UnarchiveTeam    TeamsAsyncOperationType = "unarchiveTeam"
)

func PossibleValuesForTeamsAsyncOperationType() []string {
	return []string{
		string(TeamsAsyncOperationType_ArchiveChannel),
		string(TeamsAsyncOperationType_ArchiveTeam),
		string(TeamsAsyncOperationType_CloneTeam),
		string(TeamsAsyncOperationType_CreateChannel),
		string(TeamsAsyncOperationType_CreateChat),
		string(TeamsAsyncOperationType_CreateTeam),
		string(TeamsAsyncOperationType_Invalid),
		string(TeamsAsyncOperationType_TeamifyGroup),
		string(TeamsAsyncOperationType_UnarchiveChannel),
		string(TeamsAsyncOperationType_UnarchiveTeam),
	}
}

func (s *TeamsAsyncOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAsyncOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAsyncOperationType(input string) (*TeamsAsyncOperationType, error) {
	vals := map[string]TeamsAsyncOperationType{
		"archivechannel":   TeamsAsyncOperationType_ArchiveChannel,
		"archiveteam":      TeamsAsyncOperationType_ArchiveTeam,
		"cloneteam":        TeamsAsyncOperationType_CloneTeam,
		"createchannel":    TeamsAsyncOperationType_CreateChannel,
		"createchat":       TeamsAsyncOperationType_CreateChat,
		"createteam":       TeamsAsyncOperationType_CreateTeam,
		"invalid":          TeamsAsyncOperationType_Invalid,
		"teamifygroup":     TeamsAsyncOperationType_TeamifyGroup,
		"unarchivechannel": TeamsAsyncOperationType_UnarchiveChannel,
		"unarchiveteam":    TeamsAsyncOperationType_UnarchiveTeam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAsyncOperationType(input)
	return &out, nil
}
