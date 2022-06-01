package global

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultDB(t *testing.T) {
	db, err := DefaultDB()
	assert.Equal(t, nil, err, "err should be nil")
	assert.NotEmptyf(t, db, "DefaultDB should not be empty")
}
