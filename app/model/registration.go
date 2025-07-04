package model

import (
	"github.com/uptrace/bun"
)

type Registration struct {
	bun.BaseModel `bun:"table:registrations"`

	ID           string `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	ActivitiesID string `bun:"activity_id,notnull"`
	StudentsID   string `bun:"student_id,notnull"`

	Activity *Activity `bun:"rel:belongs-to,join:activity_id=id"`
	Student  *Student  `bun:"rel:belongs-to,join:student_id=id"`

	CreateUpdateUnixTimestamp
	SoftDelete
}
