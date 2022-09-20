package test

import (
	"WildFireTest/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Config(t *testing.T) {
	testConfig, err := config.LoadConfig()

	require.NoError(t, err, "Load Config Error")

	require.Equal(t, testConfig.App.Port, 7000, "Port Number Differs!")
	require.Equal(t, testConfig.App.Limit, 10, "Limit Differs!")
	require.Equal(t, testConfig.App.Count, 500, "Count Differs!")
}
