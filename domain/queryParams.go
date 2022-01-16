package domain

// A map from key to values with its own functions
type QueryParams map[string]string

func NewQueryParams() QueryParams {
	return map[string]string{}
}

func (qp QueryParams) AddQueryParam(key, value string) {
	qp[key] = value
}

// Deletes the given key with its pair, returning the associated value
func (qp QueryParams) PopQueryParam(key string) string {
	value, ok := qp[key]
	if ok {
		delete(qp, key)
	}
	return value
}
