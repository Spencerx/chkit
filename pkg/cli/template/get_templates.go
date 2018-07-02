package clitemplate

import (
	"fmt"

	"github.com/containerum/chkit/pkg/context"
	"github.com/containerum/chkit/pkg/util/activekit"
	"github.com/containerum/chkit/pkg/util/coblog"
	"github.com/spf13/cobra"
)

var aliases = []string{"tmpl", "templates", "tmpls", "tmp", "tmps"}

func Get(ctx *context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "template",
		Aliases: aliases,
		Short:   "get solutions templates",
		Long:    "Show list of available solutions templates. To search solution by name add arg.",
		Example: "chkit get template [name]",
		Run: func(cmd *cobra.Command, args []string) {
			logger := coblog.Logger(cmd)
			logger.Debugf("loading solution info")
			var solutions, err = ctx.GetClient().GetSolutionsTemplatesList()
			if err != nil {
				logger.WithError(err).Errorf("unable to get solution list")
				activekit.Attention("Unable to get solution list:\n%v", err)
				ctx.Exit(1)
			}
			if len(args) == 1 {
				solutions = solutions.SearchByName(args[0])
			} else if len(args) > 1 {
				cmd.Help()
				ctx.Exit(1)
			}
			fmt.Println(solutions.RenderTable())
		},
	}
	return command
}