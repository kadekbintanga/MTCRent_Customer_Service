package Constant

type IDNameInterface interface {
	OptionIDNames() map[int]string
	IDAndName(id int) map[string]interface{}
	Display(id int) string
}

type IDName struct{}

func (in IDName) Get(ind IDNameInterface) []map[string]interface{} {
	var results []map[string]interface{}
	for id, name := range ind.OptionIDNames() {
		results = append(results, map[string]interface{}{
			"id":   id,
			"name": name,
		})
	}

	return results
}

func (in IDName) Display(id int, ind IDNameInterface) string {
	idNames := ind.OptionIDNames()
	if name, ok := idNames[id]; ok {
		return name
	}
	return ""
}

func (in IDName) IDAndName(id int, ind IDNameInterface) map[string]interface{} {
	return map[string]interface{}{
		"id":   id,
		"name": in.Display(id, ind),
	}
}
