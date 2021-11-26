package cmd

import (
	"CreateParcelApi/internal/app/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Start server",
	Long: `Start server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := server.NewServer(os.Getenv("APP_PORT"))

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sig
			if err := s.Shutdown(); err != nil {
				log.Error().Err(err).Msg("error during server shutdown")
			}
		}()

		return s.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
