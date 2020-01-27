package stream

import (
	"testing"

	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/types"
	"github.com/stretchr/testify/require"
)

func TestConvert_CheckServiceNode(t *testing.T) {
	require := require.New(t)

	a := &structs.CheckServiceNode{
		Node: &structs.Node{
			ID:         types.NodeID("123"),
			Node:       "node1",
			Address:    "1.2.3.4",
			Datacenter: "dc1",
			TaggedAddresses: map[string]string{
				"a": "b",
			},
			Meta: map[string]string{
				"c": "d",
			},
			RaftIndex: structs.RaftIndex{CreateIndex: 1, ModifyIndex: 2},
		},
		Service: &structs.NodeService{
			Kind:    structs.ServiceKind("default"),
			ID:      "asdf",
			Service: "foo",
			Tags:    []string{"a", "b"},
			Address: "2.3.4.5",
			TaggedAddresses: map[string]structs.ServiceAddress{
				"e": structs.ServiceAddress{Address: "3.4.5.6", Port: 2345},
			},
			Meta: map[string]string{
				"f": "g",
			},
			Port:              1234,
			Weights:           &structs.Weights{Passing: 1, Warning: 2},
			EnableTagOverride: true,
			Proxy: structs.ConnectProxyConfig{
				DestinationServiceName: "bar",
				DestinationServiceID:   "bar1",
				LocalServiceAddress:    "1.1.1.1",
				LocalServicePort:       9999,
				Config: map[string]interface{}{
					"a": float64(1),
					"b": true,
				},
				Upstreams: structs.Upstreams{
					structs.Upstream{
						DestinationType:      "a",
						DestinationNamespace: "default",
						DestinationName:      "baz",
						Datacenter:           "dc1",
						LocalBindAddress:     "2.2.2.2",
						LocalBindPort:        1111,
						Config: map[string]interface{}{
							"c": float64(1),
							"d": true,
						},
						MeshGateway: structs.MeshGatewayConfig{
							Mode: structs.MeshGatewayModeDefault,
						},
					},
					structs.Upstream{
						DestinationType:      "b",
						DestinationNamespace: "default",
						DestinationName:      "zoo",
						Datacenter:           "dc2",
						LocalBindAddress:     "3.3.3.3",
						LocalBindPort:        2222,
						Config: map[string]interface{}{
							"c": float64(1),
							"d": true,
						},
						MeshGateway: structs.MeshGatewayConfig{
							Mode: structs.MeshGatewayModeRemote,
						},
					},
				},
				MeshGateway: structs.MeshGatewayConfig{
					Mode: structs.MeshGatewayModeLocal,
				},
			},
			Connect: structs.ServiceConnect{
				Native: true,
				SidecarService: &structs.ServiceDefinition{
					Kind:    structs.ServiceKindConnectProxy,
					ID:      "asdf",
					Name:    "proxy",
					Tags:    []string{"c", "d"},
					Address: "3.4.5.6",
					TaggedAddresses: map[string]structs.ServiceAddress{
						"e": structs.ServiceAddress{Address: "4.5.6.7", Port: 3456},
					},
					Meta: map[string]string{
						"f": "g",
					},
					Port:              4567,
					Weights:           &structs.Weights{Passing: 3, Warning: 4},
					EnableTagOverride: true,
					Proxy: &structs.ConnectProxyConfig{
						DestinationServiceName: "baz",
						DestinationServiceID:   "baz1",
						LocalServiceAddress:    "5.5.5.5",
						LocalServicePort:       4444,
						Config: map[string]interface{}{
							"i": float64(1),
							"j": true,
						},
						Upstreams: structs.Upstreams{
							structs.Upstream{
								DestinationType:      "d",
								DestinationNamespace: "default",
								DestinationName:      "upstream1",
								Datacenter:           "dc3",
								LocalBindAddress:     "6.6.6.6",
								LocalBindPort:        5555,
								Config: map[string]interface{}{
									"k": float64(1),
									"l": true,
								},
								MeshGateway: structs.MeshGatewayConfig{
									Mode: structs.MeshGatewayModeNone,
								},
							},
							structs.Upstream{
								DestinationType:      "e",
								DestinationNamespace: "default",
								DestinationName:      "upstream2",
								Datacenter:           "dc4",
								LocalBindAddress:     "7.7.7.7",
								LocalBindPort:        6666,
								Config: map[string]interface{}{
									"m": float64(1),
									"n": true,
								},
								MeshGateway: structs.MeshGatewayConfig{
									Mode: structs.MeshGatewayModeLocal,
								},
							},
						},
						MeshGateway: structs.MeshGatewayConfig{
							Mode: structs.MeshGatewayModeDefault,
						},
					},
				},
			},
			LocallyRegisteredAsSidecar: true,
			RaftIndex:                  structs.RaftIndex{CreateIndex: 3, ModifyIndex: 4},
		},
	}

	b := ToCheckServiceNode(a)

	c := FromCheckServiceNode(b)

	require.Equal(a, &c)
}
