package milihttp

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/0mili/mili"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func miliConf(t *testing.T) *mili.Config {
	miliConf := new(mili.Config)
	require.NoError(t, mili.WithLogger(zaptest.NewLogger(t)).Apply(miliConf))
	return miliConf
}

func TestWithLogger(t *testing.T) {
	logger := zaptest.NewLogger(t)
	conf, err := newConf("localhost:0", miliConf(t), []Option{
		WithLogger(logger),
	})

	require.NoError(t, err)
	assert.Equal(t, logger, conf.logger)
}

func TestWithTLS(t *testing.T) {
	conf, err := newConf("localhost:0", miliConf(t), []Option{
		WithTLS("foo.cert", "foo.key"),
	})

	require.NoError(t, err)
	assert.Equal(t, "foo.cert", conf.certFile)
	assert.Equal(t, "foo.key", conf.keyFile)
}

func TestWithTLSConfig(t *testing.T) {
	tlsConf := &tls.Config{ServerName: "foo"}
	conf, err := newConf("localhost:0", miliConf(t), []Option{
		WithTLSConfig(tlsConf),
	})

	require.NoError(t, err)
	assert.Equal(t, tlsConf, conf.tlsConf)
}

func TestWithTimeouts(t *testing.T) {
	conf, err := newConf("localhost:0", miliConf(t), []Option{
		WithTimeouts(42 * time.Second),
	})

	require.NoError(t, err)
	assert.EqualValues(t, 42*time.Second, conf.readTimeout)
	assert.EqualValues(t, 42*time.Second, conf.writeTimeout)
}

func TestWithReadTimeout(t *testing.T) {
	conf, err := newConf("localhost:0", miliConf(t), []Option{
		WithReadTimeout(42 * time.Second),
	})

	require.NoError(t, err)
	assert.EqualValues(t, 42*time.Second, conf.readTimeout)
	assert.EqualValues(t, 0, conf.writeTimeout)
}

func TestWithWriteTimeout(t *testing.T) {
	conf, err := newConf("localhost:0", miliConf(t), []Option{
		WithWriteTimeout(42 * time.Second),
	})

	require.NoError(t, err)
	assert.EqualValues(t, 0, conf.readTimeout)
	assert.EqualValues(t, 42*time.Second, conf.writeTimeout)
}

func TestWithTrustedHeader(t *testing.T) {
	conf, err := newConf("localhost:0", miliConf(t), []Option{
		WithTrustedHeader("x-real-ip"),
	})
	require.NoError(t, err)
	assert.EqualValues(t, "x-real-ip", conf.trustedHeader)
}
