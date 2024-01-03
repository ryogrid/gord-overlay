package main

import (
	"context"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	sigs          = make(chan os.Signal, 1)
	done          = make(chan bool, 1)
	host          string
	existNodeHost string
)

const (
	externalServerPort = "26041" //TODO: to be configurable
	internalServerPort = "26040" //TODO: to be configurable
)

func main() {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	command := &cobra.Command{
		Use:   "gordctl",
		Short: "Run gord process and gRPC server",
		Long:  "Run gord process and gRPC server",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx, cancel = context.WithCancel(context.Background())
				localNode   = chord.NewLocalNode(host)
				transport   = server.NewChordApiClient(localNode, internalServerPort, time.Second*3)
				process     = chord.NewProcess(localNode, transport)
				opts        = []server.InternalServerOptionFunc{
					server.WithNodeOption(host),
					server.WithTimeoutConnNode(time.Second * 3),
				}
			)
			defer cancel()
			if existNodeHost != "" {
				opts = append(opts, server.WithProcessOptions(chord.WithExistNode(
					chord.NewRemoteNode(existNodeHost, process.Transport),
				)))
			}
			ins := server.NewChordServer(process, internalServerPort, opts...)
			exs := server.NewExternalServer(process, externalServerPort)
			go ins.Run(ctx)
			go exs.Run()

			<-done
			ins.Shutdown()
			exs.Shutdown()
			process.Shutdown()
		},
	}
	command.PersistentFlags().StringVarP(&host, "host", "l", "127.0.0.1", "host name to attach this process.")
	command.PersistentFlags().StringVarP(&existNodeHost, "exist-node", "n", "", "host name of exist node in chord ring.")
	if err := command.Execute(); err != nil {
		log.Fatalf("err(%#v)", err)
	}
}
