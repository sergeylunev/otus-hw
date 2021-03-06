package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("add item get item", func(t *testing.T) {
		c := NewCache(2)

		c.Set("a", 100)
		v, ok := c.Get("a")

		require.True(t, ok)
		require.Equal(t, 100, v)
	})

	t.Run("add item change item get item", func(t *testing.T) {
		c := NewCache(2)

		ok := c.Set("a", 100)
		require.False(t, ok)

		v, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 100, v)

		ok = c.Set("a", 200)
		require.True(t, ok)

		v, ok = c.Get("a")
		require.True(t, ok)
		require.Equal(t, 200, v)
	})

	t.Run("Remove last item from cache after add new item", func(t *testing.T) {
		c := NewCache(2)

		c.Set("a", 100)
		c.Set("b", 200)

		v, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 100, v)

		v, ok = c.Get("b")
		require.True(t, ok)
		require.Equal(t, 200, v)

		ok = c.Set("c", 300)
		require.False(t, ok)

		v, ok = c.Get("a")
		require.Nil(t, v)
		require.False(t, ok)

		v, ok = c.Get("c")
		require.Equal(t, 300, v)
		require.True(t, ok)
	})

	t.Run("Add add get add item. Remove second item", func(t *testing.T) {
		c := NewCache(2)

		c.Set("a", 100)
		c.Set("b", 200)
		c.Get("a")

		ok := c.Set("c", 300)
		require.False(t, ok)

		v, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 100, v)

		v, ok = c.Get("b")
		require.Nil(t, v)
		require.False(t, ok)

		v, ok = c.Get("c")
		require.Equal(t, 300, v)
		require.True(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)

		c.Clear()

		_, ok := c.Get("b")
		require.False(t, ok)
		_, ok = c.Get("c")
		require.False(t, ok)

		c.Set("d", 400)
		c.Set("a", 500)

		v, ok := c.Get("d")
		require.True(t, ok)
		require.Equal(t, 400, v)

		v, ok = c.Get("a")
		require.True(t, ok)
		require.Equal(t, 500, v)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
