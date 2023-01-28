package queue_test

import (
	"testing"

	queue "github.com/bimaagung/algoritm-data-structure/queue"
	"github.com/stretchr/testify/assert"
)

func TestAddQueue(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		result := queue.NewQueue(10)
		expected := "success add item in queue"
		assert.Equal(t, result.Add(1), expected)
	})

	t.Run("show throw fail because queue is full", func(t *testing.T) {
		result := queue.NewQueue(2)
		result.Add(10)
		result.Add(20)
		expected := "cannot add item because queue is full"
		assert.Equal(t, result.Add(30), expected)
	})
}

func TestDeleteQueue(t *testing.T) {

	t.Run("show throw fail because queue is empty", func(t *testing.T) {
		result := queue.NewQueue(10)
		expected := "cannot remove item because queue is empty"
		assert.Equal(t, result.Pop(), expected)
	})

	t.Run("success", func(t *testing.T) {
		result := queue.NewQueue(10)
		result.Add(10)
		result.Add(20)

		expected := "success remove item in queue"
		assert.Equal(t, result.Pop(), expected)
	})
}

func TestShowQueue(t *testing.T) {

	t.Run("show empty list", func(t *testing.T) {
		result := queue.NewQueue(10)
		assert.Nil(t, result.Show())
	})

	t.Run("should show list all", func(t *testing.T) {
		result := queue.NewQueue(10)
		result.Add(10)
		result.Add(20)

		assert.Equal(t, result.Show(), []int{10, 20})
	})
}