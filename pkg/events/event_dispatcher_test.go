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

type TestEventHandler struct {
	ID int
}

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
	s.handler = TestEventHandler{
		ID: 1,
	}
	s.handler2 = TestEventHandler{
		ID: 2,
	}
	s.handler3 = TestEventHandler{
		ID: 3,
	}
	s.event = TestEvent{Name: "test", Payload: "test"}
	s.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	assert.Equal(s.T(), &s.handler, s.eventDispatcher.handlers[s.event.GetName()][0])
	assert.Equal(s.T(), &s.handler2, s.eventDispatcher.handlers[s.event.GetName()][1])
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
