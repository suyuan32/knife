// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package padding

import (
	"reflect"
	"testing"
)

func TestDePaddingISO97971(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{data: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0x80, 0, 0, 0, 0}},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DePaddingISO97971(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DePaddingISO97971() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaddingISO97971(t *testing.T) {
	type args struct {
		data      []byte
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				data:      []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
				blockSize: 8,
			},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0x80, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaddingISO97971(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaddingISO97971() = %v, want %v", got, tt.want)
			}
		})
	}
}
