package filerw

import (
	"bufio"
	"errors"
	"os"
)

type Handler struct {
	Path string

	file *os.File
}

//func Write

func (h *Handler) File() error {
	var err error
	if _, err := os.Stat(h.Path); !os.IsNotExist(err) {
		h.file, err = os.OpenFile(h.Path, os.O_APPEND|os.O_RDWR, 0600)
		return err
	}
	_, err = os.Create(h.Path)
	h.File()

	return err
}

func (h *Handler) Close() error {
	if h.file == nil {
		return errors.New("nil file")
	}
	return h.file.Close()
}

func (h *Handler) WriteString(str string) error {
	if h.file == nil {
		return errors.New("nil file")
	}
	_, err := h.file.WriteString(str)
	return err
}

func (h *Handler) Read() error {
	sc := bufio.NewScanner(h.file)
	for sc.Scan() {
		_ = sc.Text() // GET the line string
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return nil
}
