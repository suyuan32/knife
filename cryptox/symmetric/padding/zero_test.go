package padding

import (
	"reflect"
	"testing"
)

func TestPaddingZero(t *testing.T) {
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
				data:      []byte("hello world"),
				blockSize: 20,
			},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "test2",
			args: args{
				data:      []byte("Hi Ryan!"),
				blockSize: 30,
			},
			want: []byte{72, 105, 32, 82, 121, 97, 110, 33, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "test3",
			args: args{
				data:      nil,
				blockSize: 0,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaddingZero(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaddingZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDePaddingZero(t *testing.T) {
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
			args: args{data: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
		},
		{
			name: "test2",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DePaddingZero(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DePaddingZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
