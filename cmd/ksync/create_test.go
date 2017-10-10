package main

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestCreateCmd_New(t *testing.T) {
	type fields struct {
		viper *viper.Viper
	}
	tests := []struct {
		name   string
		fields fields
		want   *cobra.Command
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		this := &CreateCmd{
			viper: tt.fields.viper,
		}
		if got := this.New(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. CreateCmd.New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCreateCmd_run(t *testing.T) {
	type fields struct {
		viper *viper.Viper
	}
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		this := &CreateCmd{
			viper: tt.fields.viper,
		}
		this.run(tt.args.cmd, tt.args.args)
	}
}
