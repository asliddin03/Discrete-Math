package ast

import "fmt"

type NodeType int

const (
	NodeVariable NodeType = iota
	NodeUnary
	NodeBinary
)

type Node interface {
	Type() NodeType
	String() string
}

type VarNode struct {
	Name string
}

func (v *VarNode) Type() NodeType {
	return NodeVariable
}

func (v *VarNode) String() string {
	return v.Name
}

type UnaryNode struct {
	Op   string
	Expr Node
}

func (u *UnaryNode) Type() NodeType {
	return NodeUnary
}

func (u *UnaryNode) String() string {
	return fmt.Sprintf("%s(%s)", u.Op, u.Expr.String())
}

type BinaryNode struct {
	Op          string
	Left, Right Node
}

func (b *BinaryNode) Type() NodeType {
	return NodeBinary
}

func (b *BinaryNode) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Op, b.Right.String())
}
