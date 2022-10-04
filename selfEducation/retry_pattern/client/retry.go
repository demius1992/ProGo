package main

import (
	"context"
	"math"
	"math/rand"
	"time"
)

type RetrySettings struct {
	BackoffCoef                                float64
	IntervalBetweenAttempts                    time.Duration
	MinIntervalAfterFail, MaxIntervalAfterFail time.Duration
	delay                                      time.Duration
	attemptNum                                 int32
	MaxAttempts                                int32
}

func (r *RetrySettings) NextAttempt() time.Duration {
	if r.attemptNum >= r.MaxAttempts {
		return -1
	}
	r.attemptNum++

	r.delay = time.Duration(math.Pow(r.BackoffCoef, float64(r.attemptNum)) * float64(r.MinIntervalAfterFail))
	randCoef := time.Duration(rand.Float64() * float64(r.MinIntervalAfterFail) * float64(r.attemptNum))
	r.delay += randCoef
	if r.delay > r.MaxIntervalAfterFail {
		r.delay = r.MaxIntervalAfterFail
	}
	return r.delay
}

func (r *RetrySettings) ResetAttempts() {
	r.attemptNum = 0
}

type RetryWorker func(ctx context.Context) error

type RetryHelper struct {
	settings *RetrySettings
	policy   RetryPolicy
}

func New(settings RetrySettings, policy RetryPolicy) *RetryHelper {
	return &RetryHelper{
		settings: &settings,
		policy:   policy,
	}
}

func (rh *RetryHelper) Run(ctx context.Context, job RetryWorker) error {
	defer rh.settings.ResetAttempts()

	for {
		err := job(ctx)
		validate := rh.policy.Validate(err)

		switch validate {
		case Succeed:
			if rh.settings.IntervalBetweenAttempts > 0 {
				timeout := time.After(rh.settings.IntervalBetweenAttempts)
				if err := rh.sleep(ctx, timeout); err != nil {
					return err
				}
			}
			return nil
		case Fail:
			return err
		case Retry:
			delay := rh.settings.NextAttempt()
			if delay == -1 {
				return err
			}
			timeout := time.After(delay)
			if err := rh.sleep(ctx, timeout); err != nil {
				return err
			}
		}
	}
}

func (rh *RetryHelper) sleep(ctx context.Context, t <-chan time.Time) error {
	select {
	case <-t:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
