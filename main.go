package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strings"

	//"github.com/e-dard/netbug"
	"github.com/ryogrid/gord-overlay/chord"
	"github.com/ryogrid/gord-overlay/core"
	"github.com/ryogrid/gossip-overlay/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net"
	"os"
	"os/signal"
	"runtime"
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
	runtime.GOMAXPROCS(10)

	//go func() {
	//	r := http.NewServeMux()
	//	netbug.RegisterHandler("/debug/pprof/", r)
	//	if err := http.ListenAndServe(":8080", r); err != nil {
	//		log.Fatal(err)
	//	}
	//}()

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

			//_, basePort, err := net.SplitHostPort(hostAndPortBase)
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

			// execute proxy process
			exe, err := os.Executable()
			if err != nil {
				panic(err)
			}
			dirStr := filepath.Dir(exe)
			launchDir := dirStr
			var proxyBinaryName = launchDir + "/gossip-port-forward"
			if runtime.GOOS == "windows" {
				proxyBinaryName = launchDir + "/gossip-port-forward.exe"
			}
			startCmd := proxyBinaryName
			proxyPortNumStr := strconv.Itoa(basePortNum + 3)
			argsStr := "both -a " + host + " -f " + strconv.Itoa(basePortNum) + " -l " + proxyPortNumStr
			fmt.Println("lauch proxy: ", startCmd+" "+argsStr)
			startArgs := strings.Fields(argsStr)
			proxyProc := exec.Command(startCmd, startArgs...)
			proxyProc.Dir = launchDir
			stdout, _ := proxyProc.StdoutPipe()
			stderr, _ := proxyProc.StderrPipe()
			stdoutStderr := io.MultiReader(stdout, stderr)

			// print proxy process output in background
			go func(childOut io.Reader) {
				scanner := bufio.NewScanner(childOut)
				for scanner.Scan() {
					fmt.Println(scanner.Text())
				}
			}(stdoutStderr)

			err = proxyProc.Start()
			if err != nil {
				panic(err)
			}
			time.Sleep(5 * time.Second)

			proxyAddress := host + ":" + proxyPortNumStr
			var (
				ctx, cancel = context.WithCancel(context.Background())
				localNode   = chord.NewLocalNode(hostAndPortBase)
				//transport   = core.NewChordApiClient(localNode, olPeer, time.Second*3)
				transport = core.NewChordApiClient(localNode, nil, &proxyAddress, time.Second*3*60)
				process   = chord.NewProcess(localNode, transport)
				opts      = []core.InternalServerOptionFunc{
					core.WithNodeOption(hostAndPortBase),
					//core.WithTimeoutConnNode(time.Second * 3),
					core.WithTimeoutConnNode(time.Second * 3 * 60),
				}
			)
			defer cancel()
			if existNodeHostAndPort != "" {
				opts = append(opts, core.WithProcessOptions(chord.WithExistNode(
					chord.NewRemoteNode(existNodeHostAndPort, process.Transport),
				)))
			}
			ins := core.NewChordServer(process, nil, basePort, opts...)
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
