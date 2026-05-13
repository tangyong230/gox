package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type I18nManager struct {
	mu          sync.RWMutex
	messages    map[string]map[string]string // lang -> key -> message
	defaultLang string
}

var defaultManager *I18nManager
var once sync.Once

func InitI18n(defaultLang string) error {
	var err error
	once.Do(func() {
		defaultManager = &I18nManager{
			defaultLang: defaultLang,
			messages:    make(map[string]map[string]string),
		}
		err = defaultManager.loadAllLocales()
	})
	return err
}

func (m *I18nManager) loadAllLocales() error {
	localeDir := "locales"
	files, err := os.ReadDir(localeDir)
	if err != nil {
		return fmt.Errorf("failed to read locale directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		lang := file.Name()[:len(file.Name())-5] // remove .json
		data, err := os.ReadFile(filepath.Join(localeDir, file.Name()))
		if err != nil {
			return fmt.Errorf("failed to read locale file %s: %w", file.Name(), err)
		}

		var messages map[string]string
		if err := json.Unmarshal(data, &messages); err != nil {
			return fmt.Errorf("failed to parse locale file %s: %w", file.Name(), err)
		}

		m.messages[lang] = messages
	}

	return nil
}

func GetMessage(key string, lang string) string {
	if defaultManager == nil {
		return key
	}

	defaultManager.mu.RLock()
	defer defaultManager.mu.RUnlock()

	// 尝试获取指定语言
	if msgs, ok := defaultManager.messages[lang]; ok {
		if msg, ok := msgs[key]; ok {
			return msg
		}
	}

	// 回退到默认语言
	if msgs, ok := defaultManager.messages[defaultManager.defaultLang]; ok {
		if msg, ok := msgs[key]; ok {
			return msg
		}
	}

	// 最后返回key本身
	return key
}

func GetMessageWithArgs(key string, lang string, args ...interface{}) string {
	msg := GetMessage(key, lang)
	if len(args) > 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}
