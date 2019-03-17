package log

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	b := bytes.Buffer{}

	error.SetOutput(&b)

	t.Run("Error", func(t *testing.T) {
		b.Reset()
		Error("one", 2, "three")
		assert.Equal(t, "<3>log_test.go:18: one2three\n", b.String())
	})

	t.Run("Errorf", func(t *testing.T) {
		b.Reset()
		Errorf("one: %d %s", 2, "three")
		assert.Equal(t, "<3>log_test.go:24: one: 2 three\n", b.String())
	})

	warn.SetOutput(&b)

	t.Run("Warn", func(t *testing.T) {
		b.Reset()
		Warn("one", 2, "three")
		assert.Equal(t, "<4>log_test.go:32: one2three\n", b.String())
	})

	t.Run("Warnf", func(t *testing.T) {
		b.Reset()
		Warnf("one: %d %s", 2, "three")
		assert.Equal(t, "<4>log_test.go:38: one: 2 three\n", b.String())
	})

	info.SetOutput(&b)

	t.Run("Info", func(t *testing.T) {
		b.Reset()
		Info("one", 2, "three")
		assert.Equal(t, "<5>log_test.go:46: one2three\n", b.String())
	})

	t.Run("Info", func(t *testing.T) {
		b.Reset()
		Infof("one: %d %s", 2, "three")
		assert.Equal(t, "<5>log_test.go:52: one: 2 three\n", b.String())
	})

	debug.SetOutput(&b)

	t.Run("Debug", func(t *testing.T) {
		b.Reset()
		Debug("one", 2, "three")
		assert.Equal(t, "<7>log_test.go:60: one2three\n", b.String())
	})

	t.Run("Debugf", func(t *testing.T) {
		b.Reset()
		Debugf("one: %d %s", 2, "three")
		assert.Equal(t, "<7>log_test.go:66: one: 2 three\n", b.String())
	})
}

func BenchmarkLog(b *testing.B) {
	error.SetOutput(ioutil.Discard)

	b.Run("log", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Errorf("one: %s %d", "two", 3)
		}
	})

	log.SetFlags(flags)
	log.SetOutput(ioutil.Discard)

	b.Run("std", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			log.Printf("one: %s %d", "two", 3)
		}
	})
}
