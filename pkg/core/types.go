package core

// ModelType defines AI model categories
type ModelType string

const (
	Supervised   ModelType = "Supervised"
	Unsupervised ModelType = "Unsupervised"
	Reinforcement ModelType = "Reinforcement"
)
