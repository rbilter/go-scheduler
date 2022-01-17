package queue_test

import (
	"testing"

	"bitbucket.org/rbilter/timer.loop/scheduler/queue"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPush_PushMultipleItems_ExpectMultipleItems(t *testing.T) {
	// Arrange
	q := queue.NewQueue(queue.Infinite)

	// Act
	uuid1 := uuid.New()
	q.Push(&uuid1)
	uuid2 := uuid.New()
	q.Push(&uuid2)

	// Assert
	assert.Equal(t, 2, q.Length())
}

func TestPop_PopMultipleItemsPushed_ExpectItemsPushed(t *testing.T) {
	// Arrange
	q := queue.NewQueue(queue.Infinite)
	uuid1 := uuid.New()
	q.Push(&uuid1)
	uuid2 := uuid.New()
	q.Push(&uuid2)

	// Act && Assert
	assert.Equal(t, &uuid1, q.Pop())
	assert.Equal(t, &uuid2, q.Pop())
	assert.Nil(t, q.Pop())
}
