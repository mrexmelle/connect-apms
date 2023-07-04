package event

import "time"

type Entity struct {
	ProposalId string    `json:"proposalId" bson:"proposal_id"`
	Time       time.Time `json:"time" bson:"time"`
	Status     string    `json:"status" bson:"status"`
	Actor      string    `json:"actor" bson:"actor"`
	Note       string    `json:"note" bson:"note"`
}
