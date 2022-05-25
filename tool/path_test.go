package tool

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetCurrentAbPath(t *testing.T) {
	path := GetCurrentAbPath()
	folders := strings.Split(path, "/")
	assert.Equal(t, "just-img-go-server", folders[len(folders)-2], "path should be just-img-go-server")
	assert.Equal(t, "tool", folders[len(folders)-1], "path should be tool")
}

func TestGetProjectAbsPath(t *testing.T) {
	path := GetProjectAbsPath()
	folders := strings.Split(path, "/")
	assert.Equal(t, "just-img-go-server", folders[len(folders)-1], "path should be just-img-go-server")
}
