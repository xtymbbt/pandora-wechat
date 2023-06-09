package pandoraGpt

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/zekroTJA/timedmap"
)

// ExpireMap 设置一个会超时的map。当持续有请求时，map中的值不会被删除；当未收到请求5分钟后，map中的值会立即被删除
type ExpireMap struct {
	tm *timedmap.TimedMap
}

func NewExpireMap() *ExpireMap {
	m := &ExpireMap{
		tm: timedmap.New(EXPIRE),
	}
	return m
}

func (m *ExpireMap) LoadOrStore(key string) (value *Conversation, exist bool, err error) {
	err = m.tm.SetExpires(key, EXPIRE)
	if err != nil {
		if err == timedmap.ErrKeyNotFound {
			m.tm.Set(key, value, EXPIRE, expireCallback)
			return nil, false, nil
		}
		return nil, false, err
	}
	val := m.tm.GetValue(key)
	if val == nil {
		m.tm.Set(key, value, EXPIRE, expireCallback)
		return nil, false, nil
	} else {
		return val.(*Conversation), true, nil
	}
}

func (m *ExpireMap) Store(key string, value *Conversation) (err error) {
	err = m.tm.SetExpires(key, EXPIRE)
	if err != nil {
		if err == timedmap.ErrKeyNotFound {
			m.tm.Set(key, value, EXPIRE, expireCallback)
			return nil
		}
		return err
	}
	m.tm.Set(key, value, EXPIRE, expireCallback)
	return nil
}

func expireCallback(value interface{}) {
	conversation := value.(*Conversation)
	request, err := http.NewRequest(http.MethodDelete, PandoraAddr+"/api/conversation/"+conversation.ConversationID, nil)
	if err != nil {
		logrus.Errorf("http.NewRequest to delete conversation failed: %v", err)
		return
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		logrus.Errorf("client.Do to delete conversation failed: %v", err)
		return
	}
	defer resp.Body.Close()
	// 读取 response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("read response body failed: %v", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("delete conversation failed, status code is: %d, body is: %s", resp.StatusCode, string(body))
		return
	}
	logrus.Infof("delete conversation success, conversation id is: %s, body is: %s", conversation.ConversationID, string(body))
}
