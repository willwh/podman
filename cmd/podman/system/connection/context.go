package connection

import (
	"github.com/containers/podman/v4/cmd/podman/registry"
	"github.com/containers/podman/v4/cmd/podman/validate"
	"github.com/spf13/cobra"
)

var (
	// Command: podman context
	// This is a hidden command, which was added to make converting
	// from Docker to Podman easier.
	// For now podman context just calls into podman system connection add
	// If we are adding new features, we will add them by default
	// to podman system connection
	contextCmd = &cobra.Command{
		Use:    "context",
		Short:  "Manage contexts",
		Long:   "Manage contexts",
		RunE:   validate.SubCommandExists,
		Hidden: true,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: contextCmd,
	})
}
