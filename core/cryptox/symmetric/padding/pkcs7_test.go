package padding

import (
	"reflect"
	"testing"
)

func TestDePaddingPKCS7(t *testing.T) {
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
			args: args{data: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 6, 6, 6, 6, 6, 6}},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DePaddingPKCS7(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DePaddingPKCS7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaddingPKCS7(t *testing.T) {
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
				data:      []byte{104, 101, 108, 108, 111, 32, 119, 111, 114},
				blockSize: 15,
			},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 6, 6, 6, 6, 6, 6},
		},
		{
			name: "test2",
			args: args{
				data:      []byte{104, 101, 108, 108, 111, 32, 119, 111, 114},
				blockSize: 6,
			},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaddingPKCS7(tt.args.data, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaddingPKCS7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaddingPKCS5(t *testing.T) {
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
			args: args{
				data: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114},
			},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 7, 7, 7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaddingPKCS5(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaddingPKCS5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDePaddingPKCS5(t *testing.T) {
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
			args: args{
				data: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 7, 7, 7, 7, 7, 7, 7},
			},
			want: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DePaddingPKCS5(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DePaddingPKCS5() = %v, want %v", got, tt.want)
			}
		})
	}
}
