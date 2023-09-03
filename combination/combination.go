package combination

const (
	OPERAND = iota
	UNION
	INTERSECTION
	DIFFERENCE
	SYMMETRIC_DIFFERENCE
	COMPLEMENT
)

type combinationObjectBody struct {
	objectNames []string
	operation   byte
}
