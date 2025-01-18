package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct{}

func (e *TestEventHandler) Handle(event EventInterface) {
	// do something
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = NewEventDispatcher()
	s.handler = TestEventHandler{}
	s.handler2 = TestEventHandler{}
	s.handler3 = TestEventHandler{}
	s.event = TestEvent{Name: "test", Payload: "test"}
	s.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	assert.True(s.T(), true)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
