package main

import (
	"log"
	"net"
	"strconv"

	cr "github.com/brkelkar/common_utils/configreader"
	"github.com/brkelkar/common_utils/logger"

	"github.com/brkelkar/test_grpc_server/auth"
	"google.golang.org/grpc"
)

var envVerables = []string{"SERVER_PORT", "SERVER_HOST", "AWACS_DB", "AWACS_SMART_DB", "AWACS_SMART_STOCKIST_DB",
	"DB_PORT", "DB_HOST", "DB_USERNAME", "DB_PASSWORD", "GRPC_SERVER", "GRPC_CLIENT"}

func main() {

	logger.Info("Starting Sync upload service")

	var cfg cr.Config
	//Read configuration from config.yml file
	cfg.ReadFile("config.yml")

	//Read enviroment variables
	m := cfg.ReadEnv(envVerables)

	//Over write config file variables if enviroment variable is set
	cfg.MapEnv(m)
	port := strconv.Itoa(cfg.GrpcConfig.Port)
	lis, err := net.Listen("tcp", cfg.GrpcConfig.Host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := auth.Server{}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthValidationServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
