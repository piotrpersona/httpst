package cmd

import (
	"fmt"
	"os"

	"github.com/piotrpersona/httpst/data"

	"github.com/spf13/cobra"

	"github.com/piotrpersona/httpst/app"
)

func buildRootCmd() (rootCmd *cobra.Command) {
	var (
		group, long bool
	)

	rootCmd = &cobra.Command{
		Use:   "httpst",
		Short: "Display http status code info",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) (err error) {
			httpInfo := data.GetHttpInfo()
			for _, argument := range args {
				var code data.HTTPCode
				code, err = httpInfo.RetrieveCode(argument)
				if err != nil {
					return
				}
				_, err = code.Group()
				if err != nil {
					return
				}
			}
			return
		},
		Run: func(cmd *cobra.Command, args []string) {
			appConfig := app.Config{
				Group: group,
				Long:  long,
			}
			app.Run(args, appConfig)
		},
	}

	flags := rootCmd.Flags()

	flags.BoolVarP(&group, "group", "g", false, "Display code group")
	flags.BoolVarP(&long, "long", "l", false, "Display long description")

	return
}

// Execute will execute root command.
func Execute() {
	rootCmd := buildRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
