package filerw

import (
	"testing"
	"time"

	"github.com/rs/xid"
)

func TestWrite(t *testing.T) {
	h := Handler{
		Path: xid.New().String() + ".csv",
	}

	err := h.File()
	if err != nil {
		t.Error(err)
		return
	}

	wStart := time.Now()
	var load = 100000
	for i := 0; i < load; i++ {
		err = h.WriteString("data;data;data;data;data;data\n")
		if err != nil {
			t.Error(err)
		}
	}
	t.Log("Write: ", time.Since(wStart))

	err = h.Close()
	if err != nil {
		t.Error(err)
	}

	err = h.File()
	if err != nil {
		t.Error(err)
		return
	}

	rStart := time.Now()
	err = h.Read()
	if err != nil {
		t.Error(err)
	}
	t.Log("Read: ", time.Since(rStart))

	err = h.Close()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkWrite(b *testing.B) {
	h := Handler{
		Path: xid.New().String() + ".csv",
	}

	err := h.File()
	if err != nil {
		b.Error(err)
		return
	}

	for i := 0; i < b.N; i++ {
		err = h.WriteString("data;data;data;data;data;data\n")
		if err != nil {
			b.Error(err)
		}
	}

	err = h.Close()
	if err != nil {
		b.Error(err)
	}
}
