package app

import (
	"context"
	"fmt"
	"log"

	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/grpcadp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var GADP *grpcadp.WalletGrpcConn

// start grpc connection
func grpcUp(args *AppArgs) {
	token := "Bearer " + config.Of.Grpc.GrpcTargetToken
	meta := metadata.Pairs(
		"Authorization",
		token,
		"Signature",
		config.Of.Grpc.GrpcTargetSignature,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), meta)
	target := fmt.Sprintf("%v:%v", config.Of.Grpc.GrpcTargetHost, config.Of.Grpc.GrpcTargetPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("FAILED CONNECT GRPC\n")
		panic(err)
	}

	GADP = &grpcadp.WalletGrpcConn{
		Conn: conn,
		Ctx:  ctx,
	}

	printOutUp("New Grpc connection successfully open")
}

// Stop grpc connection
func grpcDown() {
	printOutDown("Closing current grpc Redis connection...")

	if GADP.Conn != nil {
		id := GADP.Conn.GetState()

		if err := GADP.Conn.Close(); err != nil {
			log.Printf("ERROR - failed to close Grpc connection, err: %v \n", err.Error())
		}

		log.Printf("SUCCESS - Grpc connection already closed, %v \n", id)
	}
}
