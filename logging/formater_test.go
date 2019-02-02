package logging

import (
	"runtime"
	"testing"
	"time"
)

func TestDefaultFormatter_Format(t *testing.T) {
	type args struct {
		f string
		r *LogRecord
	}

	pc, _, _, _ := runtime.Caller(1)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				f: "%(name)v %(levelno)v %(levelname)v %(pathname)v %(filename)v %(lineno)v %(created)v %(asctime)s %(msecs)v %(relativeCreated)v %(message)s",
				r: &LogRecord{
					"test",
					pc,
					0,
					time.Now(),
					"%d, %s",
					[]interface{}{1, "test"},
				},
			},
			want: "1, test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(tt.args.f)
			if got := f.Format(tt.args.r); got != tt.want {
				t.Errorf("DefaultFormatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
