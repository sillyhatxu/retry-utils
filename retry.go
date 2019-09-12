package retry

import "time"

type CustomRetryFunc func() error

func Do(customRetryFunc CustomRetryFunc, opts ...Option) error {
	//default
	config := &Config{
		attempts:      3,
		delay:         500 * time.Millisecond,
		errorCallback: func(n uint, err error) {},
		delayType:     BackOffDelay,
	}

	//apply opts
	for _, opt := range opts {
		opt(config)
	}
	var n uint
	for n < config.attempts {
		err := customRetryFunc()
		if err != nil {
			config.errorCallback(n, err)
			// if this is last attempt - don't wait
			if n == config.attempts-1 {
				return err
			}
			time.Sleep(config.delayType(n, config))
			n++
			continue
		}
		break
	}
	return nil
}
