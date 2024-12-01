package main

type HashSet map[string]struct{}

func (h *HashSet) Add(value string) {
	if *h == nil {
		*h = make(HashSet)
	}

	(*h)[value] = struct{}{}
}

func (h *HashSet) Delete(value string) {
	if *h == nil {
		return
	}

	delete(*h, value)
}

func (h *HashSet) Has(value string) bool {
	if *h == nil {
		return false
	}

	_, exists := (*h)[value]
	return exists
}

func (h *HashSet) Entries() []string {
	entries := []string{}

	if *h == nil {
		return entries
	}

	for key := range *h {
		entries = append(entries, key)
	}

	return entries
}
