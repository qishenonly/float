package logger

import (
	"regexp"
	"strings"
)

// SanitizeEmail 隐藏邮箱中间部分
// user@example.com -> u***@example.com
func SanitizeEmail(email string) string {
	if len(email) == 0 {
		return email
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	localPart := parts[0]
	domain := parts[1]

	if len(localPart) <= 2 {
		return localPart + "***@" + domain
	}

	// 保留第一个字符和最后一个字符
	sanitized := string(localPart[0]) + strings.Repeat("*", len(localPart)-2) + string(localPart[len(localPart)-1])
	return sanitized + "@" + domain
}

// SanitizePhone 隐藏电话号码中间部分
// 13800138000 -> 138****8000
func SanitizePhone(phone string) string {
	if len(phone) < 4 {
		return phone
	}

	start := phone[:3]
	end := phone[len(phone)-4:]
	return start + "****" + end
}

// SanitizePassword 隐藏密码，显示长度
// "mypassword123" -> "****(13 chars)"
func SanitizePassword(password string) string {
	return "[密码已隐藏:" + string(rune(len(password))) + "位]"
}

// SanitizeCardNumber 隐藏卡号
// 6228888888888888 -> 6228********8888
func SanitizeCardNumber(cardNumber string) string {
	if len(cardNumber) < 8 {
		return cardNumber
	}

	start := cardNumber[:4]
	end := cardNumber[len(cardNumber)-4:]
	return start + "****" + end
}

// SanitizeJSON 隐藏JSON中的敏感字段
// 支持的敏感字段: password, token, verification_code, email, phone, card_number
func SanitizeJSON(data string) string {
	// 隐藏密码
	data = regexp.MustCompile(`"password":\s*"[^"]*"`).ReplaceAllString(data, `"password":"[密码已隐藏]"`)
	data = regexp.MustCompile(`"password_hash":\s*"[^"]*"`).ReplaceAllString(data, `"password_hash":"[已隐藏]"`)

	// 隐藏 token
	data = regexp.MustCompile(`"access_token":\s*"[^"]*"`).ReplaceAllString(data, `"access_token":"[TOKEN已隐藏]"`)
	data = regexp.MustCompile(`"refresh_token":\s*"[^"]*"`).ReplaceAllString(data, `"refresh_token":"[TOKEN已隐藏]"`)
	data = regexp.MustCompile(`"token":\s*"[^"]*"`).ReplaceAllString(data, `"token":"[TOKEN已隐藏]"`)

	// 隐藏验证码
	data = regexp.MustCompile(`"verification_code":\s*"[^"]*"`).ReplaceAllString(data, `"verification_code":"[CODE已隐藏]"`)

	// 隐藏敏感个人信息
	data = regexp.MustCompile(`"email":\s*"([^"]*)"`).ReplaceAllStringFunc(data, func(match string) string {
		re := regexp.MustCompile(`"email":\s*"([^"]*)"`)
		matches := re.FindStringSubmatch(match)
		if len(matches) > 1 {
			return `"email":"` + SanitizeEmail(matches[1]) + `"`
		}
		return match
	})

	data = regexp.MustCompile(`"phone":\s*"([^"]*)"`).ReplaceAllStringFunc(data, func(match string) string {
		re := regexp.MustCompile(`"phone":\s*"([^"]*)"`)
		matches := re.FindStringSubmatch(match)
		if len(matches) > 1 {
			return `"phone":"` + SanitizePhone(matches[1]) + `"`
		}
		return match
	})

	// 隐藏卡号
	data = regexp.MustCompile(`"card_number":\s*"([^"]*)"`).ReplaceAllStringFunc(data, func(match string) string {
		re := regexp.MustCompile(`"card_number":\s*"([^"]*)"`)
		matches := re.FindStringSubmatch(match)
		if len(matches) > 1 {
			return `"card_number":"` + SanitizeCardNumber(matches[1]) + `"`
		}
		return match
	})

	return data
}

// SanitizeValue 通用敏感字段隐藏
func SanitizeValue(fieldName, value string) string {
	lowerFieldName := strings.ToLower(fieldName)

	switch {
	case strings.Contains(lowerFieldName, "password"):
		return SanitizePassword(value)
	case strings.Contains(lowerFieldName, "token"):
		return "[TOKEN已隐藏]"
	case strings.Contains(lowerFieldName, "code"):
		if len(value) <= 6 && strings.Contains(lowerFieldName, "verification") {
			return "[CODE已隐藏]"
		}
	case strings.Contains(lowerFieldName, "email"):
		return SanitizeEmail(value)
	case strings.Contains(lowerFieldName, "phone"):
		return SanitizePhone(value)
	case strings.Contains(lowerFieldName, "card"):
		return SanitizeCardNumber(value)
	}

	return value
}
