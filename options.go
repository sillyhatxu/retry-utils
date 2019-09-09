package retry

import "time"

// Function signature of retry if function
type RetryIfFunc func(error) bool

type ErrorCallbackFunc func(n uint, err error)

type DelayTypeFunc func(n uint, config *Config) time.Duration

type Config struct {
	attempts      uint
	delay         time.Duration
	errorCallback ErrorCallbackFunc
	delayType     DelayTypeFunc
}

type Option func(*Config)

func Attempts(attempts uint) Option {
	return func(c *Config) {
		c.attempts = attempts
	}
}

func Delay(delay time.Duration) Option {
	return func(c *Config) {
		c.delay = delay
	}
}

func ErrorCallback(errorCallbackFunc ErrorCallbackFunc) Option {
	return func(c *Config) {
		c.errorCallback = errorCallbackFunc
	}
}

func DelayType(delayType DelayTypeFunc) Option {
	return func(c *Config) {
		c.delayType = delayType
	}
}

// BackOffDelay is a DelayType which increases delay between consecutive retries
func BackOffDelay(n uint, config *Config) time.Duration {
	return config.delay * (1 << n)
}
