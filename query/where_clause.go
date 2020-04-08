package query

type WhereClause struct {
	restriction string
}

func NewWhereClause() WhereService {
	return &WhereClause{}
}

func (w WhereClause) CreateAND(condition ConditionService) {

}

func (w WhereClause) CreateNot(condition ConditionService) string {
	panic("implement me")
}

func (w WhereClause) CreateOR(condition ConditionService) string {
	panic("implement me")
}

func (w WhereClause) GetRestriction() string {
	panic("implement me")
}