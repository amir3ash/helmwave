package clictx

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type (
	flagName string
)

var cliFlag = struct{}{}

//nolint:fatcontext
func CLIContextToContext(c *cli.Context) context.Context {
	ctx := c.Context

	for _, f := range c.FlagNames() {
		g := c.Value(f)
		log.WithField("name", f).WithField("value", g).Trace("adding flag to action context.Context")
		ctx = addFlagToContext(ctx, f, g)
	}

	ctx = addCLIToContext(ctx, c)

	return ctx
}

func addFlagToContext(ctx context.Context, name string, value any) context.Context {
	return context.WithValue(ctx, flagName(name), value)
}

func GetFlagFromContext(ctx context.Context, name string) any {
	return ctx.Value(flagName(name))
}

func addCLIToContext(ctx context.Context, c *cli.Context) context.Context {
	return context.WithValue(ctx, cliFlag, c)
}

func GetCLIFromContext(ctx context.Context) *cli.Context {
	c, ok := ctx.Value(cliFlag).(*cli.Context)
	if !ok {
		return nil
	}

	return c
}
