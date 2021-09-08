// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package install has the side-effect of registering all builtin
// service discovery config types.
package install

import (
	_ "github.com/mkby/prometheus/discovery/aws"          // register aws
	_ "github.com/mkby/prometheus/discovery/azure"        // register azure
	_ "github.com/mkby/prometheus/discovery/consul"       // register consul
	_ "github.com/mkby/prometheus/discovery/digitalocean" // register digitalocean
	_ "github.com/mkby/prometheus/discovery/dns"          // register dns
	_ "github.com/mkby/prometheus/discovery/eureka"       // register eureka
	_ "github.com/mkby/prometheus/discovery/file"         // register file
	_ "github.com/mkby/prometheus/discovery/gce"          // register gce
	_ "github.com/mkby/prometheus/discovery/hetzner"      // register hetzner
	_ "github.com/mkby/prometheus/discovery/http"         // register http
	_ "github.com/mkby/prometheus/discovery/kubernetes"   // register kubernetes
	_ "github.com/mkby/prometheus/discovery/linode"       // register linode
	_ "github.com/mkby/prometheus/discovery/marathon"     // register marathon
	_ "github.com/mkby/prometheus/discovery/moby"         // register moby
	_ "github.com/mkby/prometheus/discovery/openstack"    // register openstack
	_ "github.com/mkby/prometheus/discovery/scaleway"     // register scaleway
	_ "github.com/mkby/prometheus/discovery/triton"       // register triton
	_ "github.com/mkby/prometheus/discovery/xds"          // register xds
	_ "github.com/mkby/prometheus/discovery/zookeeper"    // register zookeeper
)
