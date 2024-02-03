package cronx

import (
	"log/slog"
	"time"
)

// Option represents a modification to the default behavior of a Cron.
type Option func(*Cron)

// WithLocation overrides the timezone of the cron instance.
func WithLocation(loc *time.Location) Option {
	return func(c *Cron) {
		c.location = loc
	}
}

// WithSeconds overrides the parser used for interpreting job schedules to
// include a seconds field as the first one.
func WithSeconds() Option {
	return WithParser(NewParser(
		Second | Minute | Hour | Dom | Month | Dow | Descriptor,
	))
}

// WithParser overrides the parser used for interpreting job schedules.
func WithParser(p ScheduleParser) Option {
	return func(c *Cron) {
		c.parser = p
	}
}

// WithEntityIDProvider sets the entity ID provider for the cron instance.
// The entity ID provider is a function that returns a unique identifier for a job entry.
// The entity ID is used to look up a snapshot or remove a job entry.
//
// Example usage:
//
//	idProvider := func() EntryID {
//	    // generate and return a unique entry ID
//	}
//	option := WithEntityIDProvider(idProvider)
//	cron := NewCron(option)
//
// Parameters:
//   - p: The entity ID provider function.
//
// Returns:
//   - An Option function that sets the entity ID provider.
func WithEntityIDProvider(p EntityIDProvider) Option {
	return func(c *Cron) {
		c.entityIDProv = p
	}
}

// WithChain specifies Job wrappers to apply to all jobs added to this cron.
// Refer to the Chain* functions in this package for provided wrappers.
func WithChain(wrappers ...JobWrapper) Option {
	return func(c *Cron) {
		c.chain = NewChain(wrappers...)
	}
}

// WithLogger uses the provided logger.
func WithLogger(logger *slog.Logger) Option {
	return func(c *Cron) {
		c.logger = logger
	}
}
