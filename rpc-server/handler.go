package main

import (
	"context"
	"log"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"gorm.io/gorm"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{
	db *gorm.DB
}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	newMessage := req.GetMessage()
	s.db.Create(&newMessage)
	log.Println(newMessage.Text)
	resp := rpc.NewSendResponse()
	resp.Code = 0
	resp.Msg = "Received: " + newMessage.Text + "\n"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	resp := rpc.NewPullResponse()
	resp.Code = 0
	var order string
	if *req.Reverse {
		order = "send_time desc"
	} else {
		order = "send_time"
	}
	s.db.Where("chat = ?", req.Chat).Where("send_time > ?",
		req.Cursor).Order(order).Limit(int(req.Limit)).Find(&resp.Messages)
	log.Println(resp.Messages)
	return resp, nil
}
