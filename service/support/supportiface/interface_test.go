// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package supportiface_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*supportiface.SupportAPI)(nil), support.New(nil))
}
