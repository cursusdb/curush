package main_test

import (
	main "curush"
	"testing"
)

func TestCurush_RunShell(t *testing.T) {
	type fields struct {
		TLS         bool
		ClusterHost string
		ClusterPort int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// more tests if you want
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			curush := &main.Curush{
				TLS:         tt.fields.TLS,
				ClusterHost: tt.fields.ClusterHost,
				ClusterPort: tt.fields.ClusterPort,
			}
			if err := curush.RunShell(); (err != nil) != tt.wantErr {
				t.Errorf("RunShell() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
