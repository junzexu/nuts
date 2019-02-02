package nuts

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestNewNutConf(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *NutConf
	}{
		{
			name: "nuts_configure.xml",
			args: args{
				name: "nuts_configure.xml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNutConf(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				dt, _ := xml.MarshalIndent(got, "", "\t")
				t.Errorf("NewNutConf() = %+v, want %+v", string(dt), tt.want)
			}
		})
	}
}
