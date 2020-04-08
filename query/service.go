package query

type ConditionService interface {
	SetFields(key string, value interface{})
	SetOperator(operator int)
	GetCondition()string
}

type WhereService interface {
	CreateAND(condition ConditionService)
	CreateOR(condition ConditionService) string
	CreateNot(condition ConditionService) string
	GetRestriction() string
}

type RestrictionFactory interface {
	Where() WhereService
}