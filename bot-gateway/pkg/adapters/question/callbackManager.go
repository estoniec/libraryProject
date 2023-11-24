package question

import (
	"context"
	"github.com/mymmrac/telego"
	"sync"
	"time"
)

type CallbackQuestion struct {
	Size int64
	C    chan telego.Update
}

type CallbackManager struct {
	mx               sync.RWMutex
	questions        map[int64]CallbackQuestion
	questionsCreated map[int64]map[int64]time.Time
	middlewareCh     chan *telego.CallbackQuery
}

func NewCallbackManager(ctx context.Context) *CallbackManager {
	manager := &CallbackManager{
		questions:        make(map[int64]CallbackQuestion),
		questionsCreated: make(map[int64]map[int64]time.Time),
		middlewareCh:     make(chan *telego.CallbackQuery),
	}

	go manager.clear(ctx)

	return manager
}

func (m *CallbackManager) NewQuestion(message telego.Update, size int64) (chan telego.Update, func()) {
	m.mx.Lock()
	defer m.mx.Unlock()

	var msg *telego.Message

	if message.CallbackQuery != nil {
		msg = message.CallbackQuery.Message
	} else {
		msg = message.Message
	}

	if _, ok := m.questions[msg.Chat.ID]; !ok {
		m.questionsCreated[msg.Chat.ID] = make(map[int64]time.Time)
		res := make(chan telego.Update)
		m.questions[msg.Chat.ID] = CallbackQuestion{
			Size: size,
			C:    res,
		}
		m.questionsCreated[msg.Chat.ID][msg.Chat.ID] = time.Now()
		return res, func() {
			m.close(msg.Chat.ID)
		}
	}

	return nil, nil
}

func (m *CallbackManager) clear(ctx context.Context) {
	interval := time.Second * 120

	ticker := time.Tick(interval)

	for {
		select {
		case <-ctx.Done():
			break
		case <-ticker:
			for chatID, chatQuestions := range m.questionsCreated {
				for _, question := range chatQuestions {
					if time.Since(question).Seconds()-interval.Seconds() < 0 {
						continue
					}

					m.close(chatID)
				}
			}
		}
	}
}

func (m *CallbackManager) close(chatID int64) {
	_, ok := m.questions[chatID]
	if ok {
		close(m.questions[chatID].C)

		delete(m.questions, chatID)

		if len(m.questionsCreated[chatID]) != 0 {
			return
		}

		delete(m.questionsCreated, chatID)
	}
}

func (m *CallbackManager) SendMsgToQuestion(message telego.Update) bool {
	m.mx.RLock()
	defer m.mx.RUnlock()

	questions, ok := m.questions[message.CallbackQuery.Message.Chat.ID]
	if !ok {
		return true
	}

	if questions.Size == 0 {
		return true
	}

	questions.Size -= 1
	// Если канал вопроса открыт, отправляем сообщение в канал
	questions.C <- message

	return false
}

func (m *CallbackManager) Middleware(ctx context.Context, message telego.Update) bool {
	return m.SendMsgToQuestion(message)
}
