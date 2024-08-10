package msgbus

import "time"

// WithID returns an option to set event's id field
func WithID(id string) EventOption {
	return func(e Event) Event {
		e.ID = id
		return e
	}
}

// WithTxID returns an option to set event's txID field
func WithTxID(txID string) EventOption {
	return func(e Event) Event {
		e.TxID = txID
		return e
	}
}

// WithSource returns an option to set event's source field
func WithSource(source string) EventOption {
	return func(e Event) Event {
		e.Source = source
		return e
	}
}

// WithOccurredAt returns an option to set event's occurredAt field
func WithOccurredAt(time time.Time) EventOption {
	return func(e Event) Event {
		e.OccurredAt = time
		return e
	}
}
