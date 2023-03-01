package types

const (
	Active  StatusType = "active"
	Blocked StatusType = "blocked"
	Deleted StatusType = "deleted"
	Pending StatusType = "pending"
)

type StatusType string

func (g StatusType) String() string {
	return string(g)
}
