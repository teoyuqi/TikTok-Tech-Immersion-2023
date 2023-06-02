package main

import (
	"context"
	"testing"
	"time"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

func TestIMServiceImpl_Send(t *testing.T) {
	type args struct {
		ctx context.Context
		req *rpc.SendRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.SendRequest{
					Message: &rpc.Message{
						Chat:     "a:b",
						Text:     "test message",
						Sender:   "a",
						SendTime: time.Now().Unix(),
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := &IMServiceImpl{}
			//got, err := s.Send(tt.args.ctx, tt.args.req)
			//assert.True(t, errors.Is(err, tt.wantErr))
			//assert.NotNil(t, got)
		})
	}
}
