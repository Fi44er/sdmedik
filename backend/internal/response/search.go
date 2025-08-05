package response

type SearchRes struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Article string `json:"article"`
	Type    string `json:"type"`
}
