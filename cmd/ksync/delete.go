package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/vapor-ware/ksync/pkg/cli"
	"github.com/vapor-ware/ksync/pkg/ksync"
)

type deleteCmd struct {
	cli.BaseCmd
}

func (d *deleteCmd) new() *cobra.Command {
	long := `Delete an existing spec. This will stop syncing files between your
	local directory and the remote containers.

	The files you've synced are not touched and the remote container is left as is.`
	example := ``

	d.Init("ksync", &cobra.Command{
		Use:     "delete [flags] [name]",
		Short:   "Delete an existing spec",
		Long:    long,
		Example: example,
		Aliases: []string{"d"},
		Args:    cobra.ExactArgs(1),
		Run:     d.run,
	})

	return d.Cmd
}

func (d *deleteCmd) run(cmd *cobra.Command, args []string) {
	name := args[0]

	specs := &ksync.SpecList{}
	if err := specs.Update(); err != nil {
		log.Fatal(err)
	}

	if !specs.Has(name) {
		log.Fatalf("%s does not exist. Did you mean something else?", name)
	}

	if err := specs.Delete(name); err != nil {
		log.Fatalf("Could not delete %s: %v", name, err)
	}

	if err := specs.Save(); err != nil {
		log.Fatal(err)
	}
}
