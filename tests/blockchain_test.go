package tests

import (
	"testing"
	"goai/pkg/blockchain"
)

func TestDeployModel(t *testing.T) {
	modelName := "TestModel"
	modelID := blockchain.DeployModel(modelName)

	if modelID == "" {
		t.Errorf("Expected a valid model ID, got an empty string")
	}
}
