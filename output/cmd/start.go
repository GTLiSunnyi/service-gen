package main

import (
	"github.com/spf13/cobra"
	"gitlab.bianjie.ai/tianle/service-gen/app"
	"gitlab.bianjie.ai/tianle/service-gen/common"
	"gitlab.bianjie.ai/tianle/service-gen/service"
	"gitlab.bianjie.ai/tianle/service-gen/{{service_name}}"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `{{service_name}}-sp start [config-file]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase := readPassphrase(1)

			configPath := ""

			if len(args) == 0 {
				configPath = common.ConfigPath
			} else {
				configPath = args[0]
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), passphrase)

			logger := common.Logger

			appInstance := app.NewApp(serviceClient, {{service_name}}.ServiceCallback, logger)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
