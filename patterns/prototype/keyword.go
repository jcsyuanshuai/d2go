package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	value     string
	visit     int
	UpdatedAt time.Time
}

func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	bytes, err := json.Marshal(k)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bytes, &newKeyword)
	if err != nil {
		return nil
	}
	return &newKeyword
}

type Keywords map[string]*Keyword

func (keywords Keywords) Clone(updates []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range keywords {
		newKeywords[k] = v
	}

	for _, item := range updates {
		newKeywords[item.value] = item.Clone()
	}
	return newKeywords
}
