package indexer

type Chunk struct {
	FilePath string
	Name     string
	Type     string // function / struct / interface
	Content  string
}
