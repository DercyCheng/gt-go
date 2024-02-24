package events

import (
	"context"
	"gitee.com/geekbang/basic-go/webook/feed/domain"
	"gitee.com/geekbang/basic-go/webook/feed/service"
	"gitee.com/geekbang/basic-go/webook/pkg/logger"
	"gitee.com/geekbang/basic-go/webook/pkg/saramax"
	"github.com/IBM/sarama"
	"strconv"
	"time"
)

const topicFeedEvent = "feed_event"

// FeedEvent 为什么你一开始就确定用消息队列
type FeedEvent struct {
	// Type 是我内部定义，我发给不同业务方
	Type string
	// 你业务方的具体的数据
	// 点赞需要三个 key:
	// liker
	// liked
	// biz + bizId
	Metadata map[string]string
	//Metadata string
	//Metadata []byte
}

// LikeFeedEvent 能不能？
// 你就得改代码了
type LikeFeedEvent struct {
	Liker int64
}

type FeedEventConsumer struct {
	client sarama.Client
	l      logger.LoggerV1
	svc    service.FeedService
}

func NewFeedEventConsumer(
	client sarama.Client,
	l logger.LoggerV1,
	svc service.FeedService) *FeedEventConsumer {
	return &FeedEventConsumer{
		svc:    svc,
		client: client,
		l:      l,
	}
}

func (r *FeedEventConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("feed_event",
		r.client)
	if err != nil {
		return err
	}
	go func() {
		err := cg.Consume(context.Background(),
			[]string{topicFeedEvent},
			saramax.NewHandler[FeedEvent](r.l, r.Consume))
		if err != nil {
			r.l.Error("退出了消费循环异常", logger.Error(err))
		}
	}()
	return err
}
func (r *FeedEventConsumer) Consume(msg *sarama.ConsumerMessage,
	evt FeedEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 我要不要在这里提前校验一下 type 是不是合法的值？
	// 我作为一个懒人，我是不校验的
	// 一旦你校验了，业务方接入的时候就要改代码
	return r.svc.CreateFeedEvent(ctx, domain.FeedEvent{
		Type: evt.Type,
		Ext:  evt.Metadata,
	})
}

func (r *FeedEventConsumer) ConsumeV1(msg *sarama.ConsumerMessage,
	evt LikeFeedEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return r.svc.CreateFeedEvent(ctx, domain.FeedEvent{
		Type: "like_event",
		Ext: map[string]string{
			"liker": strconv.FormatInt(evt.Liker, 10),
			//
		},
	})
}
