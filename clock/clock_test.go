package clock_test

import (
	"fmt"
	"testing"

	"github.com/hsmtkk/cuddly-guacamole/clock"
	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	t0 := clock.New()
	t1 := clock.New()
	t0.Add(100)
	t1.Add(30)
	assert.False(t, t0.Equal(t1))
	t1.Add(70)
	assert.True(t, t0.Equal(t1))
}

func TestString(t *testing.T) {
	t0 := clock.New()
	t0.Add(1234)
	assert.Equal(t, "20:34", t0.String())
}

func TestSub(t *testing.T) {
	t0 := clock.New()
	t0.Sub(1000)
	fmt.Println(t0.String())
	t1 := clock.New()
	t1.Add(440)
	assert.True(t, t0.Equal(t1))
}
