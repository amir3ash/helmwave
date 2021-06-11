package action

import (
	"github.com/helmwave/helmwave/pkg/plan"
	"github.com/urfave/cli/v2"
)

type Status struct {
	plandir string
}

func (l *Status) Run() error {
	p := plan.New(l.plandir)
	if err := p.Import(); err != nil {
		return err
	}
	return p.List()
}

func (l *Status) Cmd() *cli.Command {
	return &cli.Command{
		Name:  "status",
		Usage: "Status of deployed releases",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "plandir",
				Value:       ".helmwave/",
				Usage:       "Path to plandir",
				EnvVars:     []string{"HELMWAVE_PLANDIR"},
				Destination: &l.plandir,
			},
		},
		Action: toCtx(l.Run),
	}
}
