package nodes

import (
	"github.com/jrmanes/torch/config"
	"reflect"
	"testing"
)

func TestSetConsNodeDefault(t *testing.T) {
	type args struct {
		peer config.Peer
	}
	tests := []struct {
		name string
		args args
		want config.Peer
	}{
		{
			name: "Case 1: Tests default values",
			args: args{
				peer: config.Peer{
					NodeName: "consensus-full-1",
					NodeType: "consensus",
				},
			},
			want: config.Peer{
				NodeName:           "consensus-full-1",
				NodeType:           "consensus",
				ContainerName:      "consensus",
				ContainerSetupName: "consensus-setup",
				ConnectsAsEnvVar:   false,
				ConnectsTo:         nil,
				DnsConnections:     nil,
			},
		},
		{
			name: "Case 2: Tests default values already specified",
			args: args{
				peer: config.Peer{
					NodeName:           "consensus-full-1",
					NodeType:           "consensus",
					ContainerName:      "consensus",
					ContainerSetupName: "consensus-setup",
				},
			},
			want: config.Peer{
				NodeName:           "consensus-full-1",
				NodeType:           "consensus",
				ContainerName:      "consensus",
				ContainerSetupName: "consensus-setup",
				ConnectsAsEnvVar:   false,
				ConnectsTo:         nil,
				DnsConnections:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetConsNodeDefault(tt.args.peer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConsNodeDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
