package function

import (
	"context"
	"testing"

	"github.com/liquidata-inc/go-mysql-server/sql"
	"github.com/stretchr/testify/require"
)

func TestConnectionID(t *testing.T) {
	require := require.New(t)

	session := sql.NewSession("", "", "", 2)
	ctx := sql.NewContext(context.Background(), sql.WithSession(session))

	f := NewConnectionID()
	result, err := f.Eval(ctx, nil)
	require.NoError(err)
	require.Equal(uint32(2), result)
}
