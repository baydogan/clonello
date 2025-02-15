package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"github.com/baydogan/clonello/board-service/internal/database"
	"github.com/baydogan/clonello/board-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	listpb "github.com/baydogan/clonello/proto/pb"
	pb "github.com/baydogan/clonello/proto/pb"
)

type BoardServer struct {
	pb.UnimplementedBoardServiceServer
	listClient listpb.ListServiceClient
}

func NewBoardServer(listServiceAddr string) *BoardServer {
	conn, err := grpc.Dial(listServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to List Service: %v", err)
	}

	return &BoardServer{
		listClient: listpb.NewListServiceClient(conn),
	}

}

func (s *BoardServer) CreateBoard(ctx context.Context, req *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {
	collection := database.GetCollection("boards")

	board := models.Board{
		ID:      primitive.NewObjectID(),
		Title:   req.Title,
		OwnerID: req.OwnerId,
	}

	_, err := collection.InsertOne(ctx, board)
	if err != nil {
		log.Printf("Error during create board %v", err)
		return nil, err
	}

	return &pb.CreateBoardResponse{BoardId: board.ID.Hex()}, nil
}

func (s *BoardServer) GetBoards(ctx context.Context, req *pb.GetBoardsRequest) (*pb.GetBoardsResponse, error) {
	collection := database.GetCollection("boards")

	cursor, err := collection.Find(ctx, bson.M{"owner_id": req.OwnerId})
	if err != nil {
		log.Printf("Error getting boards: %v", err)
		return nil, err
	}

	var boards []*pb.Board
	for cursor.Next(ctx) {
		var board models.Board
		if err := cursor.Decode(&board); err != nil {
			log.Printf("Board decode error: %v", err)
			continue
		}

		log.Printf("Calling GetLists for Board ID: %s", board.ID.Hex())

		listResp, err := s.listClient.GetLists(ctx, &listpb.GetListsRequest{BoardId: board.ID.Hex()})
		if err != nil {
			log.Printf("Error getting lists for board %s: %v", board.ID.Hex(), err)
			listResp = &listpb.GetListsResponse{Lists: []*listpb.List{}}
		} else {
			log.Printf("Received %d lists for board %s", len(listResp.Lists), board.ID.Hex())

		}

		var lists []*pb.List
		log.Printf("Processing %d lists for board", len(listResp.Lists))
		for _, l := range listResp.Lists {
			log.Printf("Raw List Data -> ID: %s, Title: %s, BoardID: %s", l.Id, l.Title, l.BoardId)
			lists = append(lists, &pb.List{
				Id:      l.Id,
				Title:   l.Title,
				BoardId: l.BoardId,
			})
			log.Printf("Added List -> ID: %s, Title: %s", l.Id, l.Title)
		}

		log.Printf("Final List Count for Board: %d", len(lists))

		boards = append(boards, &pb.Board{
			Id:      board.ID.Hex(),
			Title:   board.Title,
			OwnerId: board.OwnerID,
			Lists:   lists,
		})
	}

	return &pb.GetBoardsResponse{Boards: boards}, nil
}
