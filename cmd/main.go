package main

import (
	"context"
	"fmt"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	sigs                 = make(chan os.Signal, 1)
	done                 = make(chan bool, 1)
	hostAndPortBase      string
	existNodeHostAndPort string
)

/*
const (
	externalServerPort = "26041" //TODO: to be configurable
	internalServerPort = "26040" //TODO: to be configurable
)
*/

func main() {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	command := &cobra.Command{
		Use:   "gordolctl",
		Short: "Run gord-overlay process and gRPC server",
		Long:  "Run gord-overlay process and gRPC server",
		Run: func(cmd *cobra.Command, args []string) {

			host, basePort, err := net.SplitHostPort(hostAndPortBase)
			if err != nil {
				fmt.Println("invalid hostAndPort. err = %#v", err)
				os.Exit(1)
			}

			var (
				ctx, cancel = context.WithCancel(context.Background())
				localNode   = chord.NewLocalNode(hostAndPortBase)
				transport   = server.NewChordApiClient(localNode, basePort, time.Second*3)
				process     = chord.NewProcess(localNode, transport)
				opts        = []server.InternalServerOptionFunc{
					server.WithNodeOption(host),
					server.WithTimeoutConnNode(time.Second * 3),
				}
			)
			defer cancel()
			if existNodeHostAndPort != "" {
				opts = append(opts, server.WithProcessOptions(chord.WithExistNode(
					chord.NewRemoteNode(existNodeHostAndPort, process.Transport),
				)))
			}
			ins := server.NewChordServer(process, basePort, opts...)
			basePortNum, err := strconv.Atoi(basePort)
			if err != nil {
				fmt.Println("invalid basePort. err = %#v", err)
				os.Exit(1)
			}
			exs := server.NewExternalServer(process, string(basePortNum+1))
			go ins.Run(ctx)
			go exs.Run()

			<-done
			ins.Shutdown()
			exs.Shutdown()
			process.Shutdown()
		},
	}
	command.PersistentFlags().StringVarP(&hostAndPortBase, "hostAndPort", "l", "127.0.0.1", "host name and port to attach this process.")
	command.PersistentFlags().StringVarP(&existNodeHostAndPort, "existNodeHostAndPort", "n", "", "host name of exist node in chord ring.")
	if err := command.Execute(); err != nil {
		log.Fatalf("err(%#v)", err)
	}
}
