package cmd

import (
	"CreateParcelApi/internal/app/producer"
	"CreateParcelApi/internal/app/server"
	"CreateParcelApi/internal/pkg/publisher"
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
		publisher, err := publisher.NewPublisher(os.Getenv("KAFKA_HOST"))
		if err != nil{
			return nil
		}
		publisherSvc := producer.NewProducer(publisher, os.Getenv("KAFKA_PUBLISH_TOPIC"))
		s := server.NewServer(os.Getenv("APP_PORT"), publisherSvc)

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
