package command

import (
	"github.com/pyroscope-io/pyroscope/pkg/cli"
	"github.com/pyroscope-io/pyroscope/pkg/config"
	"github.com/pyroscope-io/pyroscope/pkg/exec"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newConnectCmd(cfg *config.Exec) *cobra.Command {
	connectCmd := &cobra.Command{
		Use:   "connect [flags]",
		Short: "Connect to an existing process and profile it",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.NoLogging {
				logrus.SetLevel(logrus.PanicLevel)
			} else if l, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
				logrus.SetLevel(l)
			}
			if len(args) > 0 && args[0] == "help" {
				_ = cmd.Help()
				return nil
			}

			err := exec.Cli(cfg, args)
			if err != nil {
				// do not print usage in case of an error while running the subcommand
				cmd.SilenceUsage = true
			}

			return err
		},
	}

	loadFlags(cfg, connectCmd, cli.WithSkip("group-name", "user-name", "no-root-drop"))
	return connectCmd
}
