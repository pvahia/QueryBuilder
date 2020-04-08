package query

type Field struct {
	Key   string
	Value interface{}
}

type Row []*Field
func NewRow() *Row {
	return new(Row)
}

func (r *Row) SetField(key string, value interface{}) *Row{
	*r = append(*r, &Field{Key:key, Value: value})
	return r
}

func (r *Row) GetNames() (fields []string) {
	for _, v := range *r {
		fields = append(fields, v.Key)
	}
	return
}

func (r *Row) Transpile() (fields , placeholders []string, args []interface{}) {
	for _, v := range *r {
		fields = append(fields, v.Key)
		args = append(args, v.Value)
		placeholders = append(placeholders, "?")
	}
	return
}

func (r *Row) GetField() *Row {
	return r
}

type Rows []Row
func NewRows() *Rows {
	return new(Rows)
}

func (rs *Rows)SetRow(row *Row) *Rows {
	*rs = append(*rs, *row)
	return rs
}

func (rs *Rows)GetRows() *Rows {
	return rs
}

func (rs *Rows) GetValues() (args []interface{}) {
	for _, v := range *rs {
		for _, i := range v {
			args = append(args, i.Value)
		}
	}
	return
}

