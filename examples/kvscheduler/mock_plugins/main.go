//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"log"

	"github.com/ligato/cn-infra/agent"

	kvs "github.com/ligato/vpp-agent/plugins/kvscheduler"
	"github.com/ligato/vpp-agent/plugins/orchestrator"

	mock_ifplugin "github.com/ligato/vpp-agent/examples/kvscheduler/mock_plugins/ifplugin"
	mock_l2plugin "github.com/ligato/vpp-agent/examples/kvscheduler/mock_plugins/l2plugin"
	"github.com/ligato/vpp-agent/examples/kvscheduler/mock_plugins/scenario"
)

/*
	This is a simple example for demonstrating kvscheduler with mock plugins.
*/
func main() {
	exampleAgent := &ExampleAgent{
		Orchestrator: &orchestrator.DefaultPlugin,
		KVScheduler:  &kvs.DefaultPlugin,
		MockIfPlugin: &mock_ifplugin.DefaultPlugin,
		MockL2Plugin: &mock_l2plugin.DefaultPlugin,
	}

	a := agent.NewAgent(
		agent.AllPlugins(exampleAgent),
	)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

// ExampleAgent is an example agent based on mock plugins demonstrating
// the KVScheduler framework.
type ExampleAgent struct {
	// mock plugins
	MockIfPlugin *mock_ifplugin.IfPlugin
	MockL2Plugin *mock_l2plugin.L2Plugin

	// agent core infrastructure - must be listed AFTER plugins
	KVScheduler  *kvs.Scheduler
	Orchestrator *orchestrator.Plugin
}

// String returns plugin name
func (a *ExampleAgent) String() string {
	return "example-agent"
}

// Init handles initialization phase.
func (a *ExampleAgent) Init() error {
	return nil
}

// AfterInit handles the phase after initialization.
func (a *ExampleAgent) AfterInit() error {
	go scenario.Run(a.KVScheduler)
	return nil
}

// Close cleans up the resources.
func (a *ExampleAgent) Close() error {
	return nil
}