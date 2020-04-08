package query

type Condition struct {
	Field map[string]interface{}
	Operator int
	Group bool // Wrapped into round parenthesis.
}

func (c *Condition) SetFields(key string, value interface{}) {

}

func (c *Condition) SetOperator(operator int) {
	panic("implement me")
}

func (c *Condition) GetCondition() string {
	panic("implement me")
}

func NewCondition(group bool) ConditionService {
	return &Condition{Group:group}
}

