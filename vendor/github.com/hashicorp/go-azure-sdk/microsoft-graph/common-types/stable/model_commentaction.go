package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommentAction struct {
	// If true, this activity was a reply to an existing comment thread.
	IsReply nullable.Type[bool] `json:"isReply,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identity of the user who started the comment thread.
	ParentAuthor IdentitySet `json:"parentAuthor"`

	// The identities of the users participating in this comment thread.
	Participants *[]IdentitySet `json:"participants,omitempty"`
}

var _ json.Unmarshaler = &CommentAction{}

func (s *CommentAction) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsReply   nullable.Type[bool] `json:"isReply,omitempty"`
		ODataId   *string             `json:"@odata.id,omitempty"`
		ODataType *string             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsReply = decoded.IsReply
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CommentAction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["parentAuthor"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ParentAuthor' for 'CommentAction': %+v", err)
		}
		s.ParentAuthor = impl
	}

	if v, ok := temp["participants"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Participants into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentitySet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentitySetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Participants' for 'CommentAction': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Participants = &output
	}

	return nil
}
