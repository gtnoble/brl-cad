package object

const (
	OPERAND = iota
	UNION
	INTERSECTION
	DIFFERENCE
	SYMMETRIC_DIFFERENCE
	COMPLEMENT
)

type combinable interface {
	node() *combinationNode
}

type combinationNode struct {
	operation    byte
	leftOperand  *combinationNode
	rightOperand *combinationNode
	object       *DbObject
}

func (c *combinationNode) node() *combinationNode {
	return c
}

func combine(leftOperand combinable, rightOperand combinable, operation byte) *combinationNode {
	return &combinationNode{
		operation:    operation,
		leftOperand:  leftOperand.node(),
		rightOperand: rightOperand.node(),
	}
}

type combinationObjectBody struct {
	objectNames []string
	operation   byte
}
