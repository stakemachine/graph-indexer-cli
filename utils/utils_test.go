package utils

import (
	"testing"
)

func TestSubgraphHexToHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "testing hash",
			args: args{
				s: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			},
			want: "",
		},
		{
			name: "testing correct hex",
			args: args{
				s: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
			},
			want: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
		},
		{
			name: "testing random string",
			args: args{
				s: "P6JfNfd4co9A7muUYQhPMJsMUojSF",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := SubgraphHexToHash(tt.args.s); got != tt.want && err != nil {
				t.Errorf("SubgraphHexToHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzSubgraphHexToHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		SubgraphHexToHash(s)
	})
}

func TestSubgraphHashToHex(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "testing correct hash",
			args: args{
				s: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			},
			want: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
		},
		{
			name: "testing wrong hash",
			args: args{
				s: "QmR",
			},
			want: "",
		},
		{
			name: "testing correct hex",
			args: args{
				s: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
			},
			want: "",
		},
		{
			name: "testing random string",
			args: args{
				s: "P6JfNfd4co9A7muUYQhPMJsMUojSF",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := SubgraphHashToHex(tt.args.s); got != tt.want && err != nil {
				t.Errorf("SubgraphHashToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzSubgraphHashToHex(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		SubgraphHashToHex(s)
	})
}

func TestCheckIdentifier(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testing empty",
			args: args{
				s: "",
			},
			want: false,
		},
		{
			name: "testing correct hash",
			args: args{
				s: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			},
			want: false,
		},
		{
			name: "testing wrong hash",
			args: args{
				s: "QmR",
			},
			want: false,
		},
		{
			name: "testing correct hex",
			args: args{
				s: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
			},
			want: true,
		},
		{
			name: "testing incorrect hex",
			args: args{
				s: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243b232lkjd3b4e8af7500",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIdentifier(tt.args.s); got != tt.want {
				t.Errorf("CheckIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzCheckIdentifier(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		CheckIdentifier(s)
	})
}

func TestReturnIdentifier(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "testing correct hash",
			args: args{
				s: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			},
			want: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
		},
		{
			name: "testing wrong hash",
			args: args{
				s: "QmR",
			},
			want: "",
		},
		{
			name: "testing correct hex",
			args: args{
				s: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
			},
			want: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
		},
		{
			name: "testing incorrect hex",
			args: args{
				s: "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243b232lkjd3b4e8af7500",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := ReturnIdentifier(tt.args.s); got != tt.want && err != nil {
				t.Errorf("ReturnIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzReturnIdentifier(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		ReturnIdentifier(s)
	})
}
