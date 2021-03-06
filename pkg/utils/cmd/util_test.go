/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

const (
	hostURLFlagName      = "host-url"
	hostURLFlagShorthand = "u"
	hostURLFlagUsage     = "URL to run the edv instance on. Format: HostName:Port."
	hostURLEnvKey        = "TEST_HOST_URL"
)

func TestGetUserSetVar(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "short usage",
		Long:  "long usage",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// test missing both command line argument and environment vars
	env, err := GetUserSetVar(cmd, hostURLFlagName, hostURLEnvKey, false)
	require.Error(t, err)
	require.Empty(t, env)
	require.Equal(t, "neither host-url (command line flag) nor TEST_HOST_URL (environment variable) have been set",
		err.Error())

	hostURL := "localhost:8080"
	err = os.Setenv(hostURLEnvKey, hostURL)
	require.NoError(t, err)

	// test resolution via environment variable
	env, err = GetUserSetVar(cmd, hostURLFlagName, hostURLEnvKey, false)
	require.NoError(t, err)
	require.Equal(t, hostURL, env)

	// set command line arguments
	cmd.Flags().StringP(hostURLFlagName, hostURLFlagShorthand, "initial", hostURLFlagUsage)
	args := []string{"--" + hostURLFlagName, "other"}
	cmd.SetArgs(args)
	err = cmd.Execute()
	require.NoError(t, err)

	// test resolution via command line argument - no environment variable set
	env, err = GetUserSetVar(cmd, hostURLFlagName, "", false)
	require.NoError(t, err)
	require.Equal(t, "other", env)
}
