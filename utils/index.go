package utils

type Index map[string]int

func (idx Index) Add(docs []document) {
	for _, docs : range docs{
		for _, token := analyze(doc.Text){
			ids := idx[token]
		}
	}
}
