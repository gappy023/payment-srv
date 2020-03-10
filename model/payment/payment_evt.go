package payment

import (
	"context"
	"time"

	proto "microservice/part1/payment-srv/proto/payment"

	"github.com/google/uuid"
	log "github.com/micro/go-micro/v2/logger"
)

// sendPayDoneEvt 发送支付事件
func (s *service) sendPayDoneEvt(orderId int64, state int32) {
	// 构建事件
	ev := &proto.PayEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Infof("[sendPayDoneEvt] 发送支付事件，%+v\n", ev)

	// 广播
	if err := payPublisher.Publish(context.Background(), ev); err != nil {
		log.Errorf("[sendPayDoneEvt] 异常: %v", err)
	}
}
