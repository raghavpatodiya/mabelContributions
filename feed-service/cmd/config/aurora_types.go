package config

// type for paginated result
type PaginatedResult[T any] struct {
	Results []T
	Page    int
	Limit   int
}
