package command

import (
	"fmt"

	"github.com/cli/cli/internal/ghrepo"
	"github.com/cli/cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func getTitle(cmd *cobra.Command, cmdType string, matchCount int, totalMatchCount int, baseRepo ghrepo.Interface) string {
	userSetFlagCounter := 0
	limitSet := false

	cmd.Flags().Visit(func(f *pflag.Flag) {
		userSetFlagCounter += 1
		if f.Name == "limit" {
			limitSet = true
		}
	})

	title := "\n%s in %s\n\n"
	if totalMatchCount == 0 {
		msg := fmt.Sprintf("There are no open %ss", cmdType)

		if userSetFlagCounter > 0 {
			msg = fmt.Sprintf("No %ss match your search", cmdType)
		}
		return fmt.Sprintf(title, msg, ghrepo.FullName(baseRepo))
	}

	title = "\nShowing %d of %s in %s"
	if (!limitSet && userSetFlagCounter > 0) || (userSetFlagCounter > 1) {
		title += " that match your search\n\n"
	} else {
		title += "\n\n"
	}

	out := fmt.Sprintf(title, matchCount, utils.Pluralize(totalMatchCount, cmdType), ghrepo.FullName(baseRepo))

	return out
}
