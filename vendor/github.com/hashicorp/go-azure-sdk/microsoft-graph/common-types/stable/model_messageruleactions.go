package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRuleActions struct {
	// A list of categories to be assigned to a message.
	AssignCategories *[]string `json:"assignCategories,omitempty"`

	// The ID of a folder that a message is to be copied to.
	CopyToFolder nullable.Type[string] `json:"copyToFolder,omitempty"`

	// Indicates whether a message should be moved to the Deleted Items folder.
	Delete nullable.Type[bool] `json:"delete,omitempty"`

	// The email addresses of the recipients to which a message should be forwarded as an attachment.
	ForwardAsAttachmentTo *[]Recipient `json:"forwardAsAttachmentTo,omitempty"`

	// The email addresses of the recipients to which a message should be forwarded.
	ForwardTo *[]Recipient `json:"forwardTo,omitempty"`

	// Indicates whether a message should be marked as read.
	MarkAsRead nullable.Type[bool] `json:"markAsRead,omitempty"`

	// Sets the importance of the message, which can be: low, normal, high.
	MarkImportance *Importance `json:"markImportance,omitempty"`

	// The ID of the folder that a message will be moved to.
	MoveToFolder nullable.Type[string] `json:"moveToFolder,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
	PermanentDelete nullable.Type[bool] `json:"permanentDelete,omitempty"`

	// The email addresses to which a message should be redirected.
	RedirectTo *[]Recipient `json:"redirectTo,omitempty"`

	// Indicates whether subsequent rules should be evaluated.
	StopProcessingRules nullable.Type[bool] `json:"stopProcessingRules,omitempty"`
}

var _ json.Unmarshaler = &MessageRuleActions{}

func (s *MessageRuleActions) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignCategories    *[]string             `json:"assignCategories,omitempty"`
		CopyToFolder        nullable.Type[string] `json:"copyToFolder,omitempty"`
		Delete              nullable.Type[bool]   `json:"delete,omitempty"`
		MarkAsRead          nullable.Type[bool]   `json:"markAsRead,omitempty"`
		MarkImportance      *Importance           `json:"markImportance,omitempty"`
		MoveToFolder        nullable.Type[string] `json:"moveToFolder,omitempty"`
		ODataId             *string               `json:"@odata.id,omitempty"`
		ODataType           *string               `json:"@odata.type,omitempty"`
		PermanentDelete     nullable.Type[bool]   `json:"permanentDelete,omitempty"`
		StopProcessingRules nullable.Type[bool]   `json:"stopProcessingRules,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignCategories = decoded.AssignCategories
	s.CopyToFolder = decoded.CopyToFolder
	s.Delete = decoded.Delete
	s.MarkAsRead = decoded.MarkAsRead
	s.MarkImportance = decoded.MarkImportance
	s.MoveToFolder = decoded.MoveToFolder
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PermanentDelete = decoded.PermanentDelete
	s.StopProcessingRules = decoded.StopProcessingRules

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MessageRuleActions into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["forwardAsAttachmentTo"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ForwardAsAttachmentTo into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ForwardAsAttachmentTo' for 'MessageRuleActions': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ForwardAsAttachmentTo = &output
	}

	if v, ok := temp["forwardTo"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ForwardTo into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ForwardTo' for 'MessageRuleActions': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ForwardTo = &output
	}

	if v, ok := temp["redirectTo"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RedirectTo into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RedirectTo' for 'MessageRuleActions': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RedirectTo = &output
	}

	return nil
}
