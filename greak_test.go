package greak

import (
	"sync"
	"testing"
)

func TestNoDiff(t *testing.T) {
	ctx := New()
	expect(t, 0, ctx.Check())
}

func TestOne(t *testing.T) {
	ctx := New()
	var wg sync.WaitGroup
	wg.Add(1)
	die := make(chan bool)
	go func() {
		<-die
		wg.Done()
	}()
	before := ctx.Check()
	close(die)
	wg.Wait()
	after := ctx.Check()

	expect(t, 1, before)
	expect(t, 0, after)
}

func expect(t *testing.T, i int, es Entries) {
	if len(es) != i {
		t.Fatalf("Expected %d entries, but had\n%s", i, es)
	}
}
