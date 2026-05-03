package indexer

const (
	ChunkTypeFunction  = "function"
	ChunkTypeStruct    = "struct"
	ChunkTypeInterface = "interface"
)

const (
	LanguageGo = "go"
)

type Chunk struct {
	ID       string
	FilePath string
	Language string
	Name     string
	Type     string // function / struct / interface
	Content  string
}
