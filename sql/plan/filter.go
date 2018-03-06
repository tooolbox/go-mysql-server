package plan

import "gopkg.in/src-d/go-mysql-server.v0/sql"

// Filter skips rows that don't match a certain expression.
type Filter struct {
	UnaryNode
	expression sql.Expression
}

// NewFilter creates a new filter node.
func NewFilter(expression sql.Expression, child sql.Node) *Filter {
	return &Filter{
		UnaryNode:  UnaryNode{Child: child},
		expression: expression,
	}
}

// Resolved implements the Resolvable interface.
func (p *Filter) Resolved() bool {
	return p.UnaryNode.Child.Resolved() && p.expression.Resolved()
}

// RowIter implements the Node interface.
func (p *Filter) RowIter(session sql.Session) (sql.RowIter, error) {
	i, err := p.Child.RowIter(session)
	if err != nil {
		return nil, err
	}
	return &filterIter{p.expression, i, session}, nil
}

// TransformUp implements the Transformable interface.
func (p *Filter) TransformUp(f func(sql.Node) (sql.Node, error)) (sql.Node, error) {
	child, err := p.Child.TransformUp(f)
	if err != nil {
		return nil, err
	}
	return f(NewFilter(p.expression, child))
}

// TransformExpressionsUp implements the Transformable interface.
func (p *Filter) TransformExpressionsUp(f func(sql.Expression) (sql.Expression, error)) (sql.Node, error) {
	expr, err := p.expression.TransformUp(f)
	if err != nil {
		return nil, err
	}

	child, err := p.Child.TransformExpressionsUp(f)
	if err != nil {
		return nil, err
	}

	return NewFilter(expr, child), nil
}

type filterIter struct {
	cond      sql.Expression
	childIter sql.RowIter
	session   sql.Session
}

func (i *filterIter) Next() (sql.Row, error) {
	for {
		row, err := i.childIter.Next()
		if err != nil {
			return nil, err
		}

		result, err := i.cond.Eval(i.session, row)
		if err != nil {
			return nil, err
		}

		if result == true {
			return row, nil
		}
	}
}

func (i *filterIter) Close() error {
	return i.childIter.Close()
}
