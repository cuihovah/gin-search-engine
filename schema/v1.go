package schema

type Document struct {
	Type    string
	Content string
}

type Query struct {
	Type  string
	Index []string
}

