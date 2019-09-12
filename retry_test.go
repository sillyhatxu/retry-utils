package retry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	err := Do(func() error {
		log.Println("do something")
		return fmt.Errorf("this is error")
	}, Attempts(3), Delay(500*time.Millisecond), ErrorCallback(func(n uint, err error) {
		log.Println(fmt.Sprintf("retry [%d] error : %v", n, err))
	}))
	log.Println(err)
	assert.NotNil(t, err)
}

func TestRetryDefault(t *testing.T) {
	err := Do(func() error {
		log.Println("do something")
		return fmt.Errorf("this is error")
	}, ErrorCallback(func(n uint, err error) {
		log.Println(fmt.Sprintf("retry [%d] error : %v", n, err))
	}))
	log.Println(err)
	assert.NotNil(t, err)
}
