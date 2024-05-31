package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"os"
	"testing"
)

var goodInterfaceAddrsFunc = func() ([]net.Addr, error) {
	goodAddr := &net.IPNet{
		IP: net.ParseIP("128.168.0.44"),
	}
	return []net.Addr{goodAddr}, nil
}

var nonBlockingListenAndServeFunc = func(addr string, handler http.Handler) error {
	return nil
}

var nonBlockingListenAndServeFuncOnError = func(addr string, handler http.Handler) error {
	return fmt.Errorf("fake server error")
}

func TestFlags(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []struct {
		Name           string
		Args           []string
		ExpectedExit   int
		ExpectedOutput string
	}{
		{"no flag", []string{}, 0, "env: PRD"},
		{"flag development environment", []string{"-env", "DEV"}, 0, "env: DEV"},
		{"flag production environment", []string{"-env", "PRD"}, 0, "env: PRD"},
		{"flag production environment", []string{"-env", "AAA"}, 0, "env: PRD"},
	}
	for _, tc := range cases {
		prepareFlag(tc.Name, tc.Args)

		var outputBuffer bytes.Buffer

		actualExit := realMain(goodInterfaceAddrsFunc, nonBlockingListenAndServeFunc, &outputBuffer)
		actualOutput := outputBuffer.String()

		assert.Equal(t, tc.ExpectedExit, actualExit)
		assert.Contains(t, actualOutput, tc.ExpectedOutput)
	}
}

func TestRetrieveLocalIP(t *testing.T) {
	t.Run("on success", func(t *testing.T) {
		prepareFlag("Retrieve Local IP on success", []string{})

		var outputBuffer bytes.Buffer

		actualExit := realMain(goodInterfaceAddrsFunc, nonBlockingListenAndServeFunc, &outputBuffer)
		actualOutput := outputBuffer.String()

		assert.Equal(t, 0, actualExit)
		assert.Contains(t, actualOutput, "local IP: 128.168.0.44")
	})

	t.Run("on error", func(t *testing.T) {
		prepareFlag("Retrieve Local IP on error", []string{})

		var badInterfaceAddrsFunc = func() ([]net.Addr, error) {
			return nil, fmt.Errorf("fake network error")
		}
		var outputBuffer bytes.Buffer

		actualExit := realMain(badInterfaceAddrsFunc, nonBlockingListenAndServeFunc, &outputBuffer)
		actualOutput := outputBuffer.String()

		assert.Equal(t, 1, actualExit)
		assert.Contains(t, actualOutput, "no local ip found")
	})

}

func TestLaunchServer(t *testing.T) {
	t.Run("on success in dev environment", func(t *testing.T) {
		prepareFlag("Launch server in dev environment on success", []string{"-env", "DEV"})

		var outputBuffer bytes.Buffer

		var spyServerAddress string
		var spyNonBlockingListenAndServeFunc = func(addr string, handler http.Handler) error {
			spyServerAddress = addr
			return nil
		}
		actualExit := realMain(goodInterfaceAddrsFunc, spyNonBlockingListenAndServeFunc, &outputBuffer)
		actualOutput := outputBuffer.String()

		assert.Equal(t, 0, actualExit)
		assert.Equal(t, "localhost:3000", spyServerAddress)
		assert.Contains(t, actualOutput, "visit http://localhost:3000")
	})

	t.Run("on error in dev environment", func(t *testing.T) {
		prepareFlag("Launch server in dev environment on error", []string{"-env", "DEV"})

		var outputBuffer bytes.Buffer

		actualExit := realMain(goodInterfaceAddrsFunc, nonBlockingListenAndServeFuncOnError, &outputBuffer)
		actualOutput := outputBuffer.String()

		assert.Equal(t, 1, actualExit)
		assert.Contains(t, actualOutput, "error during launch of the server: fake server error")
	})

	t.Run("on success in prod environment", func(t *testing.T) {
		prepareFlag("Launch server in prod environment on success", []string{"-env", "PRD"})

		var outputBuffer bytes.Buffer

		var spyServerAddress string
		var spyNonBlockingListenAndServeFunc = func(addr string, handler http.Handler) error {
			spyServerAddress = addr
			return nil
		}
		actualExit := realMain(goodInterfaceAddrsFunc, spyNonBlockingListenAndServeFunc, &outputBuffer)
		actualOutput := outputBuffer.String()

		assert.Equal(t, 0, actualExit)
		assert.Equal(t, ":3000", spyServerAddress)
		assert.Contains(t, actualOutput, "visit http://localhost:3000")
	})

	t.Run("on error in prod environment", func(t *testing.T) {
		prepareFlag("Launch server in dev environment on error", []string{"-env", "PRD"})

		var outputBuffer bytes.Buffer

		actualExit := realMain(goodInterfaceAddrsFunc, nonBlockingListenAndServeFuncOnError, &outputBuffer)
		actualOutput := outputBuffer.String()

		assert.Equal(t, 1, actualExit)
		assert.Contains(t, actualOutput, "error during launch of the server: fake server error")
	})

}

func prepareFlag(Name string, Args []string) {
	// this call is required because otherwise flags panics, if args are set between flag.Parse calls
	flag.CommandLine = flag.NewFlagSet(Name, flag.ExitOnError)
	// we need a value to set Args[0] to, cause flag begins parsing at Args[1]
	os.Args = append([]string{Name}, Args...)
}
