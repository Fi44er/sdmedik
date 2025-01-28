package index

import "github.com/blevesearch/bleve/v2"

func (s *service) Get() bleve.Index {
	return s.index
}
