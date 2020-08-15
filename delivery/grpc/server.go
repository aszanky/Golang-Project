package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golangproject/delivery/contentpb"
	"github.com/golangproject/entity/content"
	"github.com/golangproject/repository/memcache"
	"google.golang.org/grpc"
)

//Server struct
type Server struct {
	gcache memcache.Memcache
}

//New Instantiation
func New(gc memcache.Memcache) *Server {
	//Create Listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	contentpb.RegisterContentServiceServer(s, &Server{gcache: gc})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)

	}

	return &Server{
		gcache: gc,
	}
}

//CreateContent is for Creating Content
func (s *Server) CreateContent(ctx context.Context, req *contentpb.CreateContentReq) (res *contentpb.CreateContentRes, err error) {

	c := req.GetContent()
	data := content.Content{
		ID:      c.GetId(),
		Content: c.GetContent(),
	}
	fmt.Println(data)
	err = s.gcache.Set(data.ID, data.Content)
	if err != nil {
		log.Println(err)
		return
	}

	res = &contentpb.CreateContentRes{
		Content: &contentpb.Content{
			Id:      data.ID,
			Content: data.Content,
		},
	}

	return
}

//UpdateContent is for updating content
func (s *Server) UpdateContent(ctx context.Context, req *contentpb.UpdateContentReq) (res *contentpb.UpdateContentRes, err error) {

	c := req.GetContent()
	data := content.Content{
		ID:      c.GetId(),
		Content: c.GetContent(),
	}
	err = s.gcache.Update(data.ID, data.Content)
	if err != nil {
		log.Println(err)
		return
	}

	res = &contentpb.UpdateContentRes{
		Content: &contentpb.Content{
			Id:      data.ID,
			Content: data.Content,
		},
	}

	return
}

//DeleteContent is for deleting content
func (s *Server) DeleteContent(ctx context.Context, req *contentpb.DeleteContentReq) (res *contentpb.DeleteContentRes, err error) {

	c := req.GetContent()
	data := content.Content{
		ID:      c.GetId(),
		Content: c.GetContent(),
	}

	var resp string
	success := s.gcache.Delete(data.ID)
	if !success {
		resp = "Failed"
	} else {
		resp = "Success"
	}

	res = &contentpb.DeleteContentRes{
		Success: resp,
	}

	return
}

//ReadContent for reading specific content
func (s *Server) ReadContent(ctx context.Context, req *contentpb.ReadContentReq) (res *contentpb.ReadContentRes, err error) {

	c := req.GetContent()
	data := content.Content{
		ID:      c.GetId(),
		Content: c.GetContent(),
	}

	getData, err := s.gcache.Get(data.ID)
	if err != nil {
		log.Println(err)
		return
	}

	res = &contentpb.ReadContentRes{
		Content: &contentpb.Content{
			Id:      data.ID,
			Content: getData.(string),
		},
	}

	return
}

//GetAllContent for getting all content
func (s *Server) GetAllContent(ctx context.Context, req *contentpb.GetAllContentReq) (res *contentpb.GetAllContentRes, err error) {

	getData := s.gcache.GetALL()
	if err != nil {
		log.Println(err)
		return
	}

	var result []*contentpb.Content
	for key, value := range getData {
		result = append(result, &contentpb.Content{
			Id:      key.(int32),
			Content: value.(string),
		})
	}

	res = &contentpb.GetAllContentRes{
		Content: result,
	}

	return
}
