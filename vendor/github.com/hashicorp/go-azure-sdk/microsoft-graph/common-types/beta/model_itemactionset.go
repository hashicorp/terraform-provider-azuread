package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActionSet struct {
	// A comment was added to the item.
	Comment *CommentAction `json:"comment,omitempty"`

	// An item was created.
	Create *CreateAction `json:"create,omitempty"`

	// An item was deleted.
	Delete *DeleteAction `json:"delete,omitempty"`

	// An item was edited.
	Edit *EditAction `json:"edit,omitempty"`

	// A user was mentioned in the item.
	Mention *MentionAction `json:"mention,omitempty"`

	// An item was moved.
	Move *MoveAction `json:"move,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// An item was renamed.
	Rename *RenameAction `json:"rename,omitempty"`

	// An item was restored.
	Restore *RestoreAction `json:"restore,omitempty"`

	// An item was shared.
	Share *ShareAction `json:"share,omitempty"`

	// An item was versioned.
	Version *VersionAction `json:"version,omitempty"`
}
