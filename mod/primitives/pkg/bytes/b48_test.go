package bytes_test

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/merkle/zero"
	"github.com/stretchr/testify/require"
)

func TestB48_HashTreeRoot(t *testing.T) {
	tests := []struct {
		name  string
		input bytes.B48
		want  [32]byte
	}{
		{
			name:  "Zero bytes",
			input: bytes.B48{},
			want:  zero.Hashes[1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.HashTreeRoot()
			require.NoError(t, err)
			require.Equal(t, tt.want, result)
		})
	}
}
