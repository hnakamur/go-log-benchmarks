package ltsvlog_test

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/hnakamur/ltsvlog"
)

func BenchmarkLTSVLogErr(b *testing.B) {
	benchmarks := []struct {
		name    string
		errFunc func() error
	}{
		{
			name: "StackAndTime",
			errFunc: func() error {
				return ltsvlog.Err(errors.New("some error")).Stack("").UTCTime("time2", time.Now()).String("key1", "value1")
			},
		},
		{
			name: "Stack",
			errFunc: func() error {
				return ltsvlog.Err(errors.New("some error")).Stack("").String("key1", "value1")
			},
		},
		{
			name: "Time",
			errFunc: func() error {
				return ltsvlog.Err(errors.New("some error")).UTCTime("time2", time.Now()).String("key1", "value1")
			},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			tmpfile, err := ioutil.TempFile("", "benchmark")
			if err != nil {
				b.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			logger := ltsvlog.NewLTSVLogger(tmpfile, false)
			for i := 0; i < b.N; i++ {
				err = bm.errFunc()
				logger.Err(err)
			}
		})
	}
}
