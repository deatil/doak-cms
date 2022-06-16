package validate

import (
    "github.com/gookit/validate"
)

// 验证
func Validate(data map[string]any, rules map[string]string, m map[string]string) validate.Errors {
    v := validate.Map(data)

    if len(rules) > 0 {
        for key, rule := range rules {
            v.StringRule(key, rule)
        }
    }

    // 返回信息
    v.WithMessages(m)

    if v.Validate() {
        // safeData := v.SafeData()

        return nil
    }

    return v.Errors
}
