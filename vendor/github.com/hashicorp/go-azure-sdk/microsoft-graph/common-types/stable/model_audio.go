package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Audio struct {
	// The title of the album for this audio file.
	Album nullable.Type[string] `json:"album,omitempty"`

	// The artist named on the album for the audio file.
	AlbumArtist nullable.Type[string] `json:"albumArtist,omitempty"`

	// The performing artist for the audio file.
	Artist nullable.Type[string] `json:"artist,omitempty"`

	// Bitrate expressed in kbps.
	Bitrate nullable.Type[int64] `json:"bitrate,omitempty"`

	// The name of the composer of the audio file.
	Composers nullable.Type[string] `json:"composers,omitempty"`

	// Copyright information for the audio file.
	Copyright nullable.Type[string] `json:"copyright,omitempty"`

	// The number of the disc this audio file came from.
	Disc nullable.Type[int64] `json:"disc,omitempty"`

	// The total number of discs in this album.
	DiscCount nullable.Type[int64] `json:"discCount,omitempty"`

	// Duration of the audio file, expressed in milliseconds
	Duration nullable.Type[int64] `json:"duration,omitempty"`

	// The genre of this audio file.
	Genre nullable.Type[string] `json:"genre,omitempty"`

	// Indicates if the file is protected with digital rights management.
	HasDrm nullable.Type[bool] `json:"hasDrm,omitempty"`

	// Indicates if the file is encoded with a variable bitrate.
	IsVariableBitrate nullable.Type[bool] `json:"isVariableBitrate,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The title of the audio file.
	Title nullable.Type[string] `json:"title,omitempty"`

	// The number of the track on the original disc for this audio file.
	Track nullable.Type[int64] `json:"track,omitempty"`

	// The total number of tracks on the original disc for this audio file.
	TrackCount nullable.Type[int64] `json:"trackCount,omitempty"`

	// The year the audio file was recorded.
	Year nullable.Type[int64] `json:"year,omitempty"`
}
