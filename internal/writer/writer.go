package writer

import (
	"drago/internal/event"
	"os"
	"google.golang.org/protobuf/proto"
	"drago/internal/proto"
)

type Writer struct {
	file *os.File
}

func NewWriter(path string) (*Writer, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &Writer{file: f}, nil
}

func (w *Writer) Write(batch []event.Event) error {
	pbBatch := convertToProto(batch)
	data, err := proto.Marshal(pbBatch)
	if err != nil {
		return err
	}

	_, err = w.file.Write(data)
	return err
}

func (w *Writer) Close() error {
	return w.file.Close()
}
