package question

import (
	"context"
	"github.com/mymmrac/telego"
	"sync"
	"time"
)

type Manager struct {
	mx               sync.RWMutex
	questions        map[int64]map[int64]chan *telego.Message
	questionsCreated map[int64]map[int64]time.Time
	middlewareCh     chan *telego.Message
}

func NewManager(ctx context.Context) *Manager {
	manager := &Manager{
		questions:        make(map[int64]map[int64]chan *telego.Message),
		questionsCreated: make(map[int64]map[int64]time.Time),
		middlewareCh:     make(chan *telego.Message),
	}

	go manager.clear(ctx)

	return manager
}

func (m *Manager) NewQuestion(message telego.Update) (chan *telego.Message, func()) {
	m.mx.Lock()
	defer m.mx.Unlock()

	var msg *telego.Message

	if message.CallbackQuery != nil {
		msg = message.CallbackQuery.Message
	} else {
		msg = message.Message
	}

	if _, ok := m.questions[msg.Chat.ID]; !ok {
		m.questions[msg.Chat.ID] = make(map[int64]chan *telego.Message)
		m.questionsCreated[msg.Chat.ID] = make(map[int64]time.Time)
	}

	if _, ok := m.questions[msg.Chat.ID][msg.Chat.ID]; !ok {
		res := make(chan *telego.Message)
		m.questions[msg.Chat.ID][msg.Chat.ID] = res
		m.questionsCreated[msg.Chat.ID][msg.Chat.ID] = time.Now()
		return res, func() {
			m.close(msg.Chat.ID, msg.Chat.ID)
		}
	}
	return nil, nil
}

func (m *Manager) clear(ctx context.Context) {
	interval := time.Second * 30

	ticker := time.Tick(interval)

	for {
		select {
		case <-ctx.Done():
			break
		case <-ticker:
			for chatID, chatQuestions := range m.questionsCreated {
				for userID, question := range chatQuestions {
					if time.Since(question).Seconds()-interval.Seconds() < 0 {
						continue
					}

					m.close(chatID, userID)
				}
			}
		}
	}
}

func (m *Manager) close(chatID, userID int64) {
	_, ok := m.questions[chatID][userID]
	if ok {
		close(m.questions[chatID][userID])

		delete(m.questions[chatID], userID)

		if len(m.questions[chatID]) != 0 {
			return
		}

		delete(m.questions, chatID)

		if len(m.questionsCreated[chatID]) != 0 {
			return
		}

		delete(m.questionsCreated, chatID)
	}
}

func (m *Manager) SendMsgToQuestion(message *telego.Message) bool {
	m.mx.RLock()
	defer m.mx.RUnlock()
	questions, ok := m.questions[message.Chat.ID]
	if !ok {
		return true
	}
	question, ok := questions[message.Chat.ID]
	if !ok {
		return true
	}

	// Если канал вопроса открыт, отправляем сообщение в канал
	question <- message

	return false
}

func (m *Manager) Middleware(ctx context.Context, message *telego.Message) bool {
	return m.SendMsgToQuestion(message)
}
