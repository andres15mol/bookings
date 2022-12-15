package forms

type errors map[string][]string

//Add adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}
 //Get returns the firs error message
func (e errors) Get(filed string) string {
	es := e[filed]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
