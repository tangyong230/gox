package xerror

import (
	"fmt"
	"github.com/tangyong230/gox/i18n"
	"testing"
)

func TestCustomError(t *testing.T) {
	i18n.InitI18n("en-US", "../locales")
	tests := []struct {
		errNo      int
		errkey string
		expected string
	}{
		{0, "success", "success"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Error code %d", test.errNo), func(t *testing.T) {
			customErr := NewWithLang(test.errNo, test.errkey, "en-US")
			if customErr.Error() != test.expected {
				t.Errorf("Expected error message '%s', but got '%s'", test.expected, customErr.Error())
			}
		})
	}
}


func TestCNCustomError(t *testing.T) {
	i18n.InitI18n("zh-CN", "../locales")
	tests := []struct {
		errNo      int
		errkey string
		expected string
	}{
		{0, "success", "成功"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Error code %d", test.errNo), func(t *testing.T) {
			customErr := NewWithLang(test.errNo, test.errkey, "zh-CN")
			if customErr.Error() != test.expected {
				t.Errorf("Expected error message '%s', but got '%s'", test.expected, customErr.Error())
			}
		})
	}
}
