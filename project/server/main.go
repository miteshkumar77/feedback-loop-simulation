package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "wcl.com/simulation/pxguide"
)

type server struct {
	pb.UnimplementedPxGuideServer
	mu    sync.Mutex
	peers []string
	done  bool
}

type Acceptor struct {
	mu sync.Mutex
	nP int
	nA int
	vA interface{}
}

type Proposer struct {
	mu sync.Mutex
	nP int
	me int
	vP interface{}
}

// func (s *server) Do(c context.Context, request *pb.DecidedArgs) (response *pb.DecidedReply, err error) {

// 	response = &pb.DecidedReply{
// 		Ok: len(request.V)%2 == 0,
// 	}

// 	return response, nil

// }

func (s *server) Start(c context.Context, request *pb.StartArgs) (response *pb.StartReply, err error) {

}

func (s *server) Prepare(c context.Context, request *pb.PrepareArgs) (response *pb.PrepareReply, err error) {

}

func (s *server) Accept(c context.Context, request *pb.AcceptArgs) (response *pb.AcceptReply, err error) {

}

func (s *server) Decided(c context.Context, request *pb.DecidedArgs) (response *pb.DecidedReply, err error) {

}

type RealTimeProofParams struct {
	message string
}

type RealTimeProofEnvelope struct {
	message string
}

type UpdatedSysModel struct {
	message string
}

type RealTimeGuarantees struct {
	message string
}

type RealTimeOperatingConditions struct {
	message string
}

// 1
func proofRefinement(toSentinels chan<- RealTimeProofEnvelope, toModelRefinement chan<- RealTimeProofEnvelope,
	fromSentinels <-chan RealTimeProofParams) {
	defaultM := "from proofRefinement"
	for {
		select {
		case rtp := <-fromSentinels:
			fmt.Println("PROOF REFINEMENT: received real time proof params: ", rtp.message)
		default:
		}
		select {
		case toSentinels <- RealTimeProofEnvelope{message: defaultM}:
		default:
		}
		select {
		case toModelRefinement <- RealTimeProofEnvelope{message: defaultM}:
		default:
		}
		time.Sleep(time.Duration(rand.NormFloat64()*1000+2000) * time.Millisecond)
	}
}

// 2
func modelRefinement(toLiveSystem chan<- UpdatedSysModel, fromProofRefinement <-chan RealTimeProofEnvelope) {
	defaultM := "from modelRefinement"
	for {
		select {
		case rtpe := <-fromProofRefinement:
			fmt.Println("MODEL REFINEMENT: received real time proof envelope: ", rtpe.message)
		default:
		}
		select {
		case toLiveSystem <- UpdatedSysModel{message: defaultM}:
		default:
		}
		time.Sleep(time.Duration(rand.NormFloat64()*1000+2000) * time.Millisecond)
	}
}

// 3
func liveSys(toSentinels chan<- RealTimeOperatingConditions, fromSentinels <-chan RealTimeGuarantees,
	fromModelRefinement <-chan UpdatedSysModel) {
	defaultM := "from liveSys"
	for {
		select {
		case rtg := <-fromSentinels:
			fmt.Println("LIVE SYSTEM: received a real time guarantee: ", rtg.message)
		default:
		}
		select {
		case usm := <-fromModelRefinement:
			fmt.Println("LIVE SYSTEM: received an updated sys model: ", usm.message)
		default:
		}
		select {
		case toSentinels <- RealTimeOperatingConditions{message: defaultM}:
		default:
		}
		time.Sleep(time.Duration(rand.NormFloat64()*1000+2000) * time.Millisecond)
	}
}

// 4
func sentinel(toProofRefinement chan<- RealTimeProofParams, toLiveSys chan<- RealTimeGuarantees,
	fromProofRefinement <-chan RealTimeProofEnvelope, fromLiveSystem <-chan RealTimeOperatingConditions) {
	defaultM := "from sentinel"
	for {
		select {
		case rtpe := <-fromProofRefinement:
			fmt.Println("SENTINEL: received real time proof envelope: ", rtpe.message)
		default:
		}
		select {
		case rtoc := <-fromLiveSystem:
			fmt.Println("SENTINEL: received real time operating conditions: ", rtoc.message)
		default:
		}
		select {
		case toProofRefinement <- RealTimeProofParams{message: defaultM}:
		default:
		}
		select {
		case toLiveSys <- RealTimeGuarantees{message: defaultM}:
		default:
		}
		time.Sleep(time.Duration(rand.NormFloat64()*1000+2000) * time.Millisecond)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPxGuideServer(grpcServer, &server{})

	rtpeToSentinels := make(chan RealTimeProofEnvelope, 3)
	rtpeToModelRefinement := make(chan RealTimeProofEnvelope, 3)
	rtpToProofRefinement := make(chan RealTimeProofParams, 3)
	usmToLiveSys := make(chan UpdatedSysModel, 3)
	rtgToLiveSys := make(chan RealTimeGuarantees, 3)
	rtocToSentinel := make(chan RealTimeOperatingConditions, 3)

	// Start threads
	go func() {
		proofRefinement(rtpeToSentinels, rtpeToModelRefinement, rtpToProofRefinement)
	}()
	go func() {
		modelRefinement(usmToLiveSys, rtpeToModelRefinement)
	}()
	go func() {
		liveSys(rtocToSentinel, rtgToLiveSys, usmToLiveSys)
	}()
	go func() {
		sentinel(rtpToProofRefinement, rtgToLiveSys, rtpeToSentinels, rtocToSentinel)
	}()

	grpcServer.Serve(listener)
}
