/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package plugin_test

import (
	"errors"
	"testing"

	"magma/orc8r/cloud/go/obsidian/handlers"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/plugin/mocks"
	"magma/orc8r/cloud/go/registry"
	configregistry "magma/orc8r/cloud/go/services/config/registry"
	"magma/orc8r/cloud/go/services/metricsd"
	stateregistry "magma/orc8r/cloud/go/services/state/registry"
	"magma/orc8r/cloud/go/services/streamer/mconfig/factory"
	"magma/orc8r/cloud/go/services/streamer/providers"

	"github.com/stretchr/testify/assert"
)

type errorLoader struct{}

func (errorLoader) LoadPlugins() ([]plugin.OrchestratorPlugin, error) {
	return nil, errors.New("foobar")
}

type mockLoader struct {
	ret plugin.OrchestratorPlugin
}

func (m mockLoader) LoadPlugins() ([]plugin.OrchestratorPlugin, error) {
	return []plugin.OrchestratorPlugin{m.ret}, nil
}

func TestLoadAllPlugins(t *testing.T) {
	// Happy path - just make sure all functions on the plugin are called
	mockPlugin := &mocks.OrchestratorPlugin{}
	mockPlugin.On("GetServices").Return([]registry.ServiceLocation{})
	mockPlugin.On("GetConfigManagers").Return([]configregistry.ConfigManager{})
	mockPlugin.On("GetStateSerdes").Return([]stateregistry.StateSerde{})
	mockPlugin.On("GetMconfigBuilders").Return([]factory.MconfigBuilder{})
	mockPlugin.On("GetMetricsProfiles").Times(1).Return([]metricsd.MetricsProfile{})
	mockPlugin.On("GetObsidianHandlers").Return([]handlers.Handler{})
	mockPlugin.On("GetStreamerProviders").Return([]providers.StreamProvider{})
	err := plugin.LoadAllPlugins(mockLoader{ret: mockPlugin})
	assert.NoError(t, err)
	mockPlugin.AssertNumberOfCalls(t, "GetServices", 1)
	mockPlugin.AssertNumberOfCalls(t, "GetConfigManagers", 1)
	mockPlugin.AssertNumberOfCalls(t, "GetStateSerdes", 1)
	mockPlugin.AssertNumberOfCalls(t, "GetMconfigBuilders", 1)
	mockPlugin.AssertNumberOfCalls(t, "GetMetricsProfiles", 1)
	mockPlugin.AssertNumberOfCalls(t, "GetObsidianHandlers", 1)
	mockPlugin.AssertNumberOfCalls(t, "GetStreamerProviders", 1)
	mockPlugin.AssertExpectations(t)

	// Error in the middle of registration - duplicate metrics profile
	mockPlugin.On("GetMetricsProfiles").Times(1).Return(
		[]metricsd.MetricsProfile{
			{Name: "foo"},
			{Name: "foo"},
		},
	)
	err = plugin.LoadAllPlugins(mockLoader{ret: mockPlugin})
	assert.EqualError(t, err, "A metrics profile with the name foo already exists")
	mockPlugin.AssertExpectations(t)

	// Error from loader
	err = plugin.LoadAllPlugins(errorLoader{})
	assert.EqualError(t, err, "foobar")
}
