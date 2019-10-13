package testutils

import (
	"github.com/elliotcourant/melogale/pkg/engine"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func NewTestStore(t *testing.T) (engine.Store, func()) {
	tmpDir, err := ioutil.TempDir("", "")
	if !assert.NoError(t, err, "could not create temp directory") {
		panic(err)
	}

	store, err := engine.NewStore(engine.Options{Directory: tmpDir})
	if !assert.NoError(t, err, "could not open store") {
		panic(err)
	}
	return store, func() {
		os.RemoveAll(tmpDir)
	}
}
