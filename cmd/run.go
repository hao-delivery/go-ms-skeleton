package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func runCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run microservice",
		Long:  `...`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runner(); err != nil {
				log.Fatal(err)
			}
		},
	}

	// Пример обработки параметров командной строки
	cmd.Flags().StringP("env", "e", "local", "Launch environment")

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}

	return cmd
}

func runner() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Тело команды
	fmt.Println("Hao!")

	<-ctx.Done()
	stop()

	// Код который нужно выполнить перед завершением работы
	fmt.Println("Finish!")

	return nil
}
