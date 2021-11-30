package main

import (
	"flag"
	"net"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"harmonycloud.cn/multi-cluster-manager/config"
	clientset "harmonycloud.cn/multi-cluster-manager/pkg/client/clientset/versioned"
	corecfg "harmonycloud.cn/multi-cluster-manager/pkg/core/config"
	"harmonycloud.cn/multi-cluster-manager/pkg/core/handler"
	"harmonycloud.cn/multi-cluster-manager/pkg/core/utils"
)

var (
	lisPort               int
	heartbeatExpirePeriod time.Duration
	kubeconfig            string
	masterURL             string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.IntVar(&lisPort, "listen-port", 8080, "Bind port used to provider grpc serve")
	flag.DurationVar(&heartbeatExpirePeriod, "heartbeat-expire-period", 30, "The period of maximum heartbeat interval")
}

func main() {
	flag.Parse()

	addr := ":" + strconv.Itoa(lisPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("listen port %d error: %s", lisPort, err)
	}

	// construct client
	kubeCfg, err := utils.GetKubeConfig(kubeconfig, masterURL)
	if err != nil {
		logrus.Fatalf("failed connect kube-apiserver: %s", err)
	}
	mClient, err := clientset.NewForConfig(kubeCfg)
	if err != nil {
		logrus.Fatalf("failed get multicluster client set: %s", err)
	}

	cfg := corecfg.DefaultConfiguration()
	cfg.HeartbeatExpirePeriod = heartbeatExpirePeriod

	s := grpc.NewServer()
	config.RegisterChannelServer(s, &handler.Channel{
		Server: handler.NewCoreServer(cfg, mClient),
	})
	if err := s.Serve(l); err != nil {
		logrus.Fatalf("grpc server running error: %s", err)
	}

	logrus.Infof("listening port %d", lisPort)
}
