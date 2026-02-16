package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceAnnouncementBase = ServiceUpdateMessage{}

type ServiceUpdateMessage struct {
	// The expected deadline of the action for the message.
	ActionRequiredByDateTime nullable.Type[string] `json:"actionRequiredByDateTime,omitempty"`

	// A collection of serviceAnnouncementAttachments.
	Attachments *[]ServiceAnnouncementAttachment `json:"attachments,omitempty"`

	// The zip file that contains all attachments for a message.
	AttachmentsArchive nullable.Type[string] `json:"attachmentsArchive,omitempty"`

	Body     *ItemBody              `json:"body,omitempty"`
	Category *ServiceUpdateCategory `json:"category,omitempty"`

	// Indicates whether the message has any attachment.
	HasAttachments *bool `json:"hasAttachments,omitempty"`

	// Indicates whether the message describes a major update for the service.
	IsMajorChange nullable.Type[bool] `json:"isMajorChange,omitempty"`

	// The affected services by the service message.
	Services *[]string `json:"services,omitempty"`

	Severity *ServiceUpdateSeverity `json:"severity,omitempty"`

	// A collection of tags for the service message. Tags are provided by the service team/support team who post the message
	// to tell whether this message contains privacy data, or whether this message is for a service new feature update, and
	// so on.
	Tags *[]string `json:"tags,omitempty"`

	// Represents user viewpoints data of the service message. This data includes message status such as whether the user
	// has archived, read, or marked the message as favorite. This property is null when accessed with application
	// permissions.
	ViewPoint *ServiceUpdateMessageViewpoint `json:"viewPoint,omitempty"`

	// Fields inherited from ServiceAnnouncementBase

	// More details about service event. This property doesn't support filters.
	Details *[]KeyValuePair `json:"details,omitempty"`

	// The end time of the service event.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The last modified time of the service event.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The start time of the service event.
	StartDateTime *string `json:"startDateTime,omitempty"`

	// The title of the service event.
	Title *string `json:"title,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ServiceUpdateMessage) ServiceAnnouncementBase() BaseServiceAnnouncementBaseImpl {
	return BaseServiceAnnouncementBaseImpl{
		Details:              s.Details,
		EndDateTime:          s.EndDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		StartDateTime:        s.StartDateTime,
		Title:                s.Title,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s ServiceUpdateMessage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceUpdateMessage{}

func (s ServiceUpdateMessage) MarshalJSON() ([]byte, error) {
	type wrapper ServiceUpdateMessage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceUpdateMessage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceUpdateMessage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceUpdateMessage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceUpdateMessage: %+v", err)
	}

	return encoded, nil
}
