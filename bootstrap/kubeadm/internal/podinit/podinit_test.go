/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package podinit

import (
	"reflect"
	"testing"

	"sigs.k8s.io/cluster-api/bootstrap/kubeadm/internal/cloudinit"
)

func TestNewInitControlPlane(t *testing.T) {
	type args struct {
		input *cloudinit.ControlPlaneInput
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInitControlPlane(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInitControlPlane() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInitControlPlane() = %v, want %v", got, tt.want)
			}
		})
	}
}
