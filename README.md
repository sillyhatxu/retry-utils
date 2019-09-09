# retry-utils

### examples

http get with retry:

```
func TestRetry(t *testing.T) {
	err := Do(func() error {
		log.Println("do something")
		return fmt.Errorf("this is error")
	}, Attempts(5), Delay(2*time.Second), ErrorCallback(func(n uint, err error) {
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
```

[examples](https://https://github.com/sillyhatxu/retry-utils/retry_test.go)