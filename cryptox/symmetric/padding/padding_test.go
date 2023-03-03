package padding

import (
	"reflect"
	"testing"
)

func TestDePadding(t *testing.T) {
	type args struct {
		data      []byte
		method    PaddingType
		blockSize int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data:      []byte{0, 0},
				method:    No,
				blockSize: 8,
			},
			want:    []byte{0, 0},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				data:      []byte{1, 1, 0, 0},
				method:    Zero,
				blockSize: 4,
			},
			want:    []byte{1, 1},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				data:      []byte{0, 0, 128, 0},
				method:    ISO97971,
				blockSize: 4,
			},
			want:    []byte{0, 0},
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				data:      []byte{0, 0, 6, 6, 6, 6, 6, 6},
				method:    PKCS5,
				blockSize: 4,
			},
			want:    []byte{0, 0},
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				data:      []byte{0, 0, 2, 2},
				method:    PKCS7,
				blockSize: 4,
			},
			want:    []byte{0, 0},
			wantErr: false,
		},
		{
			name: "test6",
			args: args{
				data:      []byte{0, 0},
				method:    No,
				blockSize: 4,
			},
			want:    []byte{0, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DePadding(tt.args.data, tt.args.method, tt.args.blockSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("DePadding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DePadding() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadding(t *testing.T) {
	type args struct {
		data      []byte
		method    PaddingType
		blockSize int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data:      []byte{0, 0},
				method:    No,
				blockSize: 8,
			},
			want:    []byte{0, 0},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				data:      []byte{0, 0},
				method:    Zero,
				blockSize: 4,
			},
			want:    []byte{0, 0, 0, 0},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				data:      []byte{0, 0},
				method:    ISO97971,
				blockSize: 4,
			},
			want:    []byte{0, 0, 128, 0},
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				data:      []byte{0, 0},
				method:    PKCS5,
				blockSize: 4,
			},
			want:    []byte{0, 0, 6, 6, 6, 6, 6, 6},
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				data:      []byte{0, 0},
				method:    PKCS7,
				blockSize: 4,
			},
			want:    []byte{0, 0, 2, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Padding(tt.args.data, tt.args.method, tt.args.blockSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Padding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Padding() got = %v, want %v", got, tt.want)
			}
		})
	}
}
