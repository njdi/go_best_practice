package go_best_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionalOptions(t *testing.T) {
	// 使用默认选项创建 Server 实例
	s1 := NewServer()
	assert.Equal(t, "127.0.0.1", s1.Host)
	assert.Equal(t, 80, s1.Port)

	// 使用 Functional Options 来设置 Host 和 Port，创建 Server 实例，
	s2 := NewServer(WithHost("localhost"), WithPort(8080))
	assert.Equal(t, "localhost", s2.Host)
	assert.Equal(t, 8080, s2.Port)
}
