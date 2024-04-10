package extractembed

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveFile(t *testing.T) {
	err := saveFile("tmp.txt", []byte("Hello"), "./target")
	assert.NoError(t, err)

	a, err := os.ReadFile("./target/tmp.txt")
	assert.NoError(t, err)

	assert.Equal(t, "Hello", string(a))

	err = os.Remove("./target/tmp.txt")
	assert.NoError(t, err)

	err = os.Remove("./target")
	assert.NoError(t, err)
}
