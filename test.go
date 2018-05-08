package main

import (
	"log"
	"net"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/mingkaic/go_tenncor/graphmgr"
	ser "github.com/mingkaic/go_tenncor/serial"
)

const (
	port    = ":50051"
	address = "localhost:50051"
)

type Server struct{}

func (s *Server) PostGraph(ctx context.Context, in *pb.GraphCreated) (*pb.Nothing, error) {
	return &pb.Nothing{}, nil
}

func (s *Server) PostData(ctx context.Context, in *pb.DataCreated) (*pb.Nothing, error) {
	return &pb.Nothing{}, nil
}

func (s *Server) GetGraphList(ctx context.Context, in *pb.GraphListReq) (*pb.GraphList, error) {
	return &pb.GraphList{
		Gids: []string{"sample"},
	}, nil
}

func (s *Server) GetGraphPb(ctx context.Context, in *pb.GraphReq) (*ser.GraphPb, error) {
	createOrder := []string{"A", "B", "C"}
	nodeMap := make(map[string]*ser.NodePb)
	nodeMap["A"] = &ser.NodePb{}
	nodeMap["B"] = &ser.NodePb{}
	nodeMap["C"] = &ser.NodePb{}
	return &ser.GraphPb{
		Gid:         "sample",
		CreateOrder: createOrder,
		NodeMap:     nodeMap,
	}, nil
}

func (s *Server) GetData(in *pb.DataReq, stream pb.Graphmgr_GetDataServer) error {
	return nil
}

func (s *Server) GetTestData(in *pb.DataReq, stream pb.Graphmgr_GetTestDataServer) error {
	return nil
}

func startServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGraphmgrServer(s, &Server{})
	s.Serve(lis)
}

func TestClient(t *testing.T) {
	go startServer()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGraphmgrClient(conn)
	_, err = c.PostGraph(context.Background(), &pb.GraphCreated{
		Graph: &ser.GraphPb{
			Gid: "sample_graphid",
			CreateOrder: []string{
				"nodeId1",
				"nodeId2",
				"nodeId3",
			},
			NodeMap: map[string]*ser.NodePb{
				"nodeId1": &ser.NodePb{
					Type:  ser.NodePb_VARIABLE,
					Label: "A",
				},
				"nodeId2": &ser.NodePb{
					Type:  ser.NodePb_VARIABLE,
					Label: "B",
				},
				"nodeId3": &ser.NodePb{
					Type:  ser.NodePb_FUNCTOR,
					Label: "C",
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("could notify graph creation: %v", err)
	}
}
