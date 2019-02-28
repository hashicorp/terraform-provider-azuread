package guid

import (
	"github.com/gofrs/uuid"
)

func New() uuid.UUID {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return uuid
}
