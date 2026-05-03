package indexer

const (
	ChunkTypeFunction  = "function"
	ChunkTypeStruct    = "struct"
	ChunkTypeInterface = "interface"
)

type Chunk struct {
	FilePath string
	Language string
	Name     string
	Type     string // function / struct / interface
	Content  string
}
