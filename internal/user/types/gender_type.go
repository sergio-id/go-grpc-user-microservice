package types

const (
	Male    GenderType = "male"
	Female  GenderType = "female"
	Unknown GenderType = "unknown"
)

type GenderType string

func (g GenderType) String() string {
	return string(g)
}
