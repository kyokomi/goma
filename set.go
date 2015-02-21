package main

type set map[string]bool

func (s set) add(key string) {
	if key != "" {
		s[key] = true
	}
}

func (s set) slice() []string {
	var keys []string
	for key, val := range s {
		if val {
			keys = append(keys, key)
		}
	}
	return keys
}
