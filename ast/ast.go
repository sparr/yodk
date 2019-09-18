package ast

import (
	"strconv"

	"github.com/dbaumgarten/yodk/tokenizer"
)

type Node interface {
	Accept(v Visitor) error
	Start() tokenizer.Position
	End() tokenizer.Position
}

type Programm struct {
	Lines []*Line
}

func (n *Programm) Start() tokenizer.Position {
	return n.Lines[0].Start()
}

func (n *Programm) End() tokenizer.Position {
	return n.Lines[len(n.Lines)-1].End()
}

type Line struct {
	Statements []Statement
}

func (n *Line) Start() tokenizer.Position {
	return n.Statements[0].Start()
}

func (n *Line) End() tokenizer.Position {
	return n.Statements[len(n.Statements)-1].End()
}

// Expressions

type Expression interface {
	Node
}

type StringConstant struct {
	Position tokenizer.Position
	Value    string
}

func (n *StringConstant) Start() tokenizer.Position {
	return n.Position
}

func (n *StringConstant) End() tokenizer.Position {
	pos := n.Start()
	pos.Coloumn += len(n.Value) + 2
	return pos
}

type NumberConstant struct {
	Position tokenizer.Position
	Value    string
}

func (n *NumberConstant) Start() tokenizer.Position {
	return n.Position
}

func (n *NumberConstant) End() tokenizer.Position {
	return n.Start().Add(len(n.Value))
}

type Dereference struct {
	Position    tokenizer.Position
	Variable    string
	Operator    string
	PrePost     string
	IsStatement bool
}

func (n *Dereference) Start() tokenizer.Position {
	return n.Position
}

func (n *Dereference) End() tokenizer.Position {
	return n.Start().Add(len(n.Variable) + len(n.Operator))
}

type UnaryOperation struct {
	Position tokenizer.Position
	Operator string
	Exp      Expression
}

func (n *UnaryOperation) Start() tokenizer.Position {
	return n.Position
}

func (n *UnaryOperation) End() tokenizer.Position {
	return n.Exp.End()
}

type BinaryOperation struct {
	Operator string
	Exp1     Expression
	Exp2     Expression
}

func (n *BinaryOperation) Start() tokenizer.Position {
	return n.Exp1.Start()
}

func (n *BinaryOperation) End() tokenizer.Position {
	return n.Exp2.End()
}

type FuncCall struct {
	Function string
	Argument Expression
}

func (n *FuncCall) Start() tokenizer.Position {
	return n.Argument.Start().Sub(len(n.Function) + 1)
}

func (n *FuncCall) End() tokenizer.Position {
	return n.Argument.End().Add(1)
}

// Statements

type Statement interface {
	Node
}

type Assignment struct {
	Position tokenizer.Position
	Variable string
	Value    Expression
	Operator string
}

func (n *Assignment) Start() tokenizer.Position {
	return n.Position
}

func (n *Assignment) End() tokenizer.Position {
	return n.Value.End()
}

type IfStatement struct {
	Position  tokenizer.Position
	Condition Expression
	IfBlock   []Statement
	ElseBlock []Statement
}

func (n *IfStatement) Start() tokenizer.Position {
	return n.Position
}

func (n *IfStatement) End() tokenizer.Position {
	if n.ElseBlock == nil {
		return n.IfBlock[len(n.IfBlock)-1].End().Add(3)
	}
	return n.ElseBlock[len(n.ElseBlock)-1].End().Add(3)
}

type GoToStatement struct {
	Position tokenizer.Position
	Line     int
}

func (n *GoToStatement) Start() tokenizer.Position {
	return n.Position
}

func (n *GoToStatement) End() tokenizer.Position {
	return n.Position.Add(len(strconv.Itoa(n.Line)) + 1)
}
