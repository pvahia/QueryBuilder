package query

// Operators
const (
	EQ = iota + 0
	NE
	LE
	GE
	LT
	GT
	AND
	OR
	BETWEEN
	LINK
	IN
	ISNULL
)

type Operator int
func (o Operator)ToString() string {
	switch o {
	case EQ:
		return "="
	case NE:
		return "!="
	case LE:
		return "<="
	case GE:
		return ">="
	case LT:
		return "<"
	case GT:
		return ">"
	default:
		return ""
	}
}

type Restriction struct {

}

func (r Restriction) Where() WhereService {
	return NewWhereClause()
}

func NewRestriction() RestrictionFactory {
	return &Restriction{}
}


