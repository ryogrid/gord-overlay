package main

import (
	"context"
	"fmt"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/core"
	"github.com/ryogrid/gossip-overlay/overlay"
	"github.com/ryogrid/gossip-overlay/util"
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

func main() {
	//overlay_setting.OVERLAY_DEBUG = true

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
			basePortNum, err := strconv.Atoi(basePort)
			if err != nil {
				fmt.Println("invalid basePort. err = %#v", err)
				os.Exit(1)
			}

			peers := &util.Stringset{}
			if existNodeHostAndPort != "" {
				peers.Set(existNodeHostAndPort)
			}

			olPeer, err := overlay.NewOverlayPeer(&host, uint16(basePortNum), peers)
			if err != nil {
				fmt.Println("failed to create overlay peer. err = %#v", err)
				panic(err)
			}

			var (
				ctx, cancel = context.WithCancel(context.Background())
				localNode   = chord.NewLocalNode(hostAndPortBase)
				//transport   = core.NewChordApiClient(localNode, olPeer, time.Second*3)
				transport = core.NewChordApiClient(localNode, olPeer, time.Second*60)
				process   = chord.NewProcess(localNode, transport)
				opts      = []core.InternalServerOptionFunc{
					core.WithNodeOption(hostAndPortBase),
					//core.WithTimeoutConnNode(time.Second * 3),
					core.WithTimeoutConnNode(time.Second * 60),
				}
			)
			defer cancel()
			if existNodeHostAndPort != "" {
				opts = append(opts, core.WithProcessOptions(chord.WithExistNode(
					chord.NewRemoteNode(existNodeHostAndPort, process.Transport),
				)))
			}
			ins := core.NewChordServer(process, olPeer, basePort, opts...)
			exs := core.NewExternalServer(process, strconv.Itoa(basePortNum+1))
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
