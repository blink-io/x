// Copyright 2021 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package msgbus

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/google/uuid"
)

type (
	Bus interface {
		Emit(ctx context.Context, topic string, data any) error
		EmitWithOpts(ctx context.Context, topic string, data any, opts ...EventOption) error
		Topics() []string
		RegisterTopics(topics ...string)
		DeregisterTopics(topics ...string)
		TopicHandlerKeys(topic string) []string
		HandlerKeys() []string
		HandlerTopicSubscriptions(handlerKey string) []string
		RegisterHandler(key string, h Handler)
		DeregisterHandler(key string)
	}

	// Bus is a message bus
	bus struct {
		mutex    sync.RWMutex
		idgen    Next
		topics   map[string][]Handler
		handlers map[string]Handler
	}

	// Next is a sequential unique id generator func type
	Next func() string

	// IDGenerator is a sequential unique id generator interface
	IDGenerator interface {
		Generate() string
	}

	// Event is data structure for any logs
	Event struct {
		ID         string    // identifier
		TxID       string    // transaction identifier
		Topic      string    // topic name
		Source     string    // source of the event
		OccurredAt time.Time // creation time in nanoseconds
		Data       any       // actual event data
	}

	// Handler is a receiver for event reference with the given regex pattern
	Handler struct {
		key string

		// handler func to process events
		Handle func(ctx context.Context, e Event)

		// topic matcher as regex pattern
		Matcher string
	}

	// EventOption is a function type to mutate event fields
	EventOption = func(Event) Event

	ctxKey int8
)

const (
	// CtxKeyTxID tx id context key
	CtxKeyTxID = ctxKey(116)

	// CtxKeySource source context key
	CtxKeySource = ctxKey(117)

	// Version syncs with package version
	Version = "3.0.3"

	empty = ""
)

// Generate is an implementation of IDGenerator for bus.Next fn type
func (n Next) Generate() string {
	return n()
}

// New inits a new bus
func New() Bus {
	b, _ := NewWithGen(Next(uuid.NewString))
	return b
}

// NewWithGen inits a new bus with id generator
func NewWithGen(g IDGenerator) (Bus, error) {
	if g == nil {
		return nil, fmt.Errorf("bus: Next() id generator func can't be nil")
	}

	return &bus{
		idgen:    g.Generate,
		topics:   make(map[string][]Handler),
		handlers: make(map[string]Handler),
	}, nil
}

// Emit inits a new event and delivers to the interested in handlers with
// sync safety
func (b *bus) Emit(ctx context.Context, topic string, data any) error {
	b.mutex.RLock()
	handlers, ok := b.topics[topic]
	b.mutex.RUnlock()

	if !ok {
		return fmt.Errorf("bus: topic(%s) not found", topic)
	}

	source, _ := ctx.Value(CtxKeySource).(string)
	txID, _ := ctx.Value(CtxKeyTxID).(string)
	if txID == empty {
		txID = b.idgen()
		ctx = context.WithValue(ctx, CtxKeyTxID, txID)
	}

	e := Event{
		ID:         b.idgen(),
		Topic:      topic,
		Data:       data,
		OccurredAt: time.Now(),
		TxID:       txID,
		Source:     source,
	}

	for _, h := range handlers {
		h.Handle(ctx, e)
	}

	return nil
}

// EmitWithOpts inits a new event and delivers to the interested in handlers
// with sync safety and options
func (b *bus) EmitWithOpts(ctx context.Context, topic string, data any, opts ...EventOption) error {
	b.mutex.RLock()
	handlers, ok := b.topics[topic]
	b.mutex.RUnlock()

	if !ok {
		return fmt.Errorf("bus: topic(%s) not found", topic)
	}

	e := Event{Topic: topic, Data: data}
	for _, o := range opts {
		e = o(e)
	}

	if e.TxID == empty {
		e.TxID = b.idgen()
	}
	if e.ID == empty {
		e.ID = b.idgen()
	}
	if e.OccurredAt.IsZero() {
		e.OccurredAt = time.Now()
	}

	for _, h := range handlers {
		h.Handle(ctx, e)
	}

	return nil
}

// Topics lists the all registered topics
func (b *bus) Topics() []string {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	topics, index := make([]string, len(b.topics)), 0

	for topic := range b.topics {
		topics[index] = topic
		index++
	}
	return topics
}

// RegisterTopics registers topics and fulfills handlers
func (b *bus) RegisterTopics(topics ...string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, n := range topics {
		b.registerTopic(n)
	}
}

// DeregisterTopics deletes topic
func (b *bus) DeregisterTopics(topics ...string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, n := range topics {
		b.deregisterTopic(n)
	}
}

// TopicHandlerKeys returns all handlers for the topic
func (b *bus) TopicHandlerKeys(topic string) []string {
	b.mutex.RLock()
	handlers := b.topics[topic]
	b.mutex.RUnlock()

	keys := make([]string, len(handlers))

	for i, h := range handlers {
		keys[i] = h.key
	}

	return keys
}

// HandlerKeys returns list of registered handler keys
func (b *bus) HandlerKeys() []string {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	keys, index := make([]string, len(b.handlers)), 0

	for k := range b.handlers {
		keys[index] = k
		index++
	}
	return keys
}

// HandlerTopicSubscriptions returns all topic subscriptions of the handler
func (b *bus) HandlerTopicSubscriptions(handlerKey string) []string {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return b.handlerTopicSubscriptions(handlerKey)
}

// RegisterHandler re/register the handler to the registry
func (b *bus) RegisterHandler(key string, h Handler) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	h.key = key
	b.registerHandler(h)
}

// DeregisterHandler deletes handler from the registry
func (b *bus) DeregisterHandler(key string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.deregisterHandler(key)
}

func (b *bus) registerHandler(h Handler) {
	b.deregisterHandler(h.key)
	b.handlers[h.key] = h
	for _, t := range b.handlerTopicSubscriptions(h.key) {
		b.registerTopicHandler(t, h)
	}
}

func (b *bus) deregisterHandler(handlerKey string) {
	if _, ok := b.handlers[handlerKey]; ok {
		for _, t := range b.handlerTopicSubscriptions(handlerKey) {
			b.deregisterTopicHandler(t, handlerKey)
		}
		delete(b.handlers, handlerKey)
	}
}

func (b *bus) registerTopicHandler(topic string, h Handler) {
	b.topics[topic] = append(b.topics[topic], h)
}

func (b *bus) deregisterTopicHandler(topic, handlerKey string) {
	l := len(b.topics[topic])
	for i, h := range b.topics[topic] {
		if h.key == handlerKey {
			b.topics[topic][i] = b.topics[topic][l-1]
			b.topics[topic] = b.topics[topic][:l-1]
			break
		}
	}
}

func (b *bus) registerTopic(topic string) {
	if _, ok := b.topics[topic]; ok {
		return
	}

	b.topics[topic] = b.buildHandlers(topic)
}

func (b *bus) deregisterTopic(topic string) {
	delete(b.topics, topic)
}

func (b *bus) buildHandlers(topic string) []Handler {
	handlers := make([]Handler, 0)
	for _, h := range b.handlers {
		if matched, _ := regexp.MatchString(h.Matcher, topic); matched {
			handlers = append(handlers, h)
		}
	}
	return handlers
}

func (b *bus) handlerTopicSubscriptions(handlerKey string) []string {
	var subscriptions []string
	h, ok := b.handlers[handlerKey]
	if !ok {
		return subscriptions
	}

	for topic := range b.topics {
		if matched, _ := regexp.MatchString(h.Matcher, topic); matched {
			subscriptions = append(subscriptions, topic)
		}
	}
	return subscriptions
}
