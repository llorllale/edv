/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package edv

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestController_New(t *testing.T) {
	controller, err := New(nil, "")
	require.NoError(t, err)
	require.NotNil(t, controller)
}

func TestController_GetOperations(t *testing.T) {
	controller, err := New(nil, "")
	require.NoError(t, err)
	require.NotNil(t, controller)

	ops := controller.GetOperations()

	require.Equal(t, 3, len(ops))

	require.Equal(t, "/data-vaults", ops[0].Path())
	require.Equal(t, http.MethodPost, ops[0].Method())
	require.NotNil(t, ops[0].Handle())

	require.Equal(t, "/encrypted-data-vaults/{vaultID}/docs", ops[1].Path())
	require.Equal(t, http.MethodPost, ops[1].Method())
	require.NotNil(t, ops[1].Handle())

	require.Equal(t, "/encrypted-data-vaults/{vaultID}/docs/{docID}", ops[2].Path())
	require.Equal(t, http.MethodGet, ops[2].Method())
	require.NotNil(t, ops[2].Handle())
}
