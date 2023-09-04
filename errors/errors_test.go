package errors_test

import (
	"fmt"
	"github.com/sinhashubham95/go-utils/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	e := errors.New("hello")
	assert.Error(t, e)
	assert.Equal(t, "hello", e.Error())

	er, ok := e.(*errors.Error)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, er.StatusCode)
	assert.Equal(t, "hello", er.Message)
}

func TestError(t *testing.T) {
	e := (&errors.Error{Code: "naruto"}).Value().WithStatusCode(http.StatusOK).WithMessage("naruto").
		WithDetails("naruto")
	assert.Equal(t, "naruto", e.Code)
	assert.Equal(t, "naruto", e.Message)
	assert.Equal(t, "naruto", e.Details)
	assert.Equal(t, http.StatusOK, e.StatusCode)
	assert.Equal(t, "naruto", e.Error())
}

func TestErrorWrapUnwrap(t *testing.T) {
	e := errors.New("naruto")
	assert.Equal(t, e, errors.Wrap(e, errors.New("naruto")))
	assert.Error(t, errors.Unwrap(e))
	e = errors.New("naruto")
	assert.Equal(t, e, errors.Wrap(e, fmt.Errorf("naruto")))
	assert.Nil(t, (&errors.Error{Details: "naruto"}).Unwrap())
	assert.Nil(t, errors.Wrap(nil, nil))
	assert.Nil(t, errors.Unwrap(nil))
}

func TestErrorIs(t *testing.T) {
	assert.True(t, errors.Is(errors.New("naruto"), errors.New("naruto")))
	assert.False(t, errors.Is(errors.New("naruto"), nil))
	assert.True(t, errors.Is(errors.Wrap(errors.New("naruto"), errors.New("boruto")), errors.New("boruto")))
	assert.False(t, errors.Is(errors.New("naruto"), errors.New("boruto")))
	assert.True(t, errors.Is(fmt.Errorf("naruto"), fmt.Errorf("naruto")))
}

func TestErrorAs(t *testing.T) {
	var e *errors.Error
	assert.False(t, errors.As(errors.New("naruto"), nil))
	assert.False(t, errors.As(errors.New("naruto"), e))
	assert.True(t, errors.As(errors.New("naruto"), &e))
	assert.Equal(t, "naruto", e.Message)
	assert.Equal(t, http.StatusInternalServerError, e.StatusCode)
	var x string
	assert.False(t, errors.As(errors.New("naruto"), &x))
	var y struct{}
	assert.False(t, errors.As(errors.New("naruto"), &y))
	assert.False(t, errors.As(errors.Wrap(errors.New("naruto"), errors.New("boruto")), &y))
	e = &errors.Error{StatusCode: http.StatusOK, Message: "naruto"}
	ne := &errors.Error{}
	assert.True(t, errors.As(e, ne))
	assert.Equal(t, http.StatusOK, ne.StatusCode)
	assert.Equal(t, "naruto", ne.Message)
}
