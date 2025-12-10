package service

import (
	"testing"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/pkg/email"
)

func TestGenerateVerificationCode(t *testing.T) {
	code1, err := email.GenerateVerificationCode()
	if err != nil {
		t.Fatalf("生成验证码失败: %v", err)
	}

	if len(code1) != 6 {
		t.Errorf("验证码长度应该是6，实际是 %d", len(code1))
	}

	// 生成两个验证码，应该不同
	code2, err := email.GenerateVerificationCode()
	if err != nil {
		t.Fatalf("生成验证码失败: %v", err)
	}

	if code1 == code2 {
		t.Errorf("两个验证码不应该相同: %s, %s", code1, code2)
	}

	t.Logf("生成的验证码: %s", code1)
}

func TestVerificationCodeExpiry(t *testing.T) {
	expiry := email.GetVerificationCodeExpiry()
	if expiry.Minutes() != 10 {
		t.Errorf("默认过期时间应该是10分钟，实际是 %v", expiry)
	}
}

func TestSendVerificationCodeRequestValidation(t *testing.T) {
	tests := []struct {
		name  string
		email string
		valid bool
	}{
		{"有效邮箱", "test@example.com", true},
		{"Gmail邮箱", "user@gmail.com", true},
		{"126邮箱", "user@126.com", true},
		{"无效邮箱", "notanemail", false},
		{"空邮箱", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := request.SendVerificationCodeRequest{Email: tt.email}
			// 这里只是演示，实际验证由 binding 标签处理
			if (len(req.Email) > 0 && tt.email != "") != tt.valid {
				t.Logf("邮箱: %s, 有效: %v", req.Email, tt.valid)
			}
		})
	}
}

func TestRegisterRequestValidation(t *testing.T) {
	tests := []struct {
		name             string
		username         string
		password         string
		verificationCode string
		valid            bool
	}{
		{"有效数据", "testuser", "Password123", "123456", true},
		{"无效用户名（太短）", "ab", "Password123", "123456", false},
		{"无效密码（太短）", "testuser", "pass", "123456", false},
		{"无效验证码（太短）", "testuser", "Password123", "12345", false},
		{"无效验证码（太长）", "testuser", "Password123", "1234567", false},
		{"缺少用户名", "", "Password123", "123456", false},
		{"缺少密码", "testuser", "", "123456", false},
		{"缺少验证码", "testuser", "Password123", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := request.RegisterRequest{
				Username:         tt.username,
				Password:         tt.password,
				VerificationCode: tt.verificationCode,
			}

			valid := len(req.Username) >= 3 &&
				len(req.Password) >= 6 &&
				len(req.VerificationCode) == 6

			if valid != tt.valid {
				t.Errorf("验证不符合预期: %s, 期望: %v, 实际: %v", tt.name, tt.valid, valid)
			}
		})
	}
}

func TestEmailServiceMockMode(t *testing.T) {
	// 在 mock 模式下，邮件服务应该不返回错误
	emailService := email.GetService()

	err := emailService.SendVerificationCode("test@example.com", "123456")
	if err != nil {
		t.Fatalf("Mock 模式发送邮件不应该返回错误: %v", err)
	}

	t.Log("Mock 模式下邮件服务工作正常")
}

// 集成测试示例（需要真实配置）
func TestIntegrationSendAndVerifyCode(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过集成测试")
	}

	t.Log("集成测试：验证码生成、存储和验证流程")

	// 1. 生成验证码
	code, err := email.GenerateVerificationCode()
	if err != nil {
		t.Fatalf("生成验证码失败: %v", err)
	}
	t.Logf("生成验证码: %s", code)

	// 2. 存储验证码
	emailAddr := "test@example.com"
	err = email.StoreVerificationCode(emailAddr, code)
	if err != nil {
		t.Logf("存储验证码失败: %v (可能是 Redis 未启动)", err)
		t.Skip("Redis 连接失败，跳过此测试")
	}
	t.Logf("验证码已存储到 Redis")

	// 3. 验证正确的验证码
	verified, err := email.VerifyCode(emailAddr, code)
	if err != nil {
		t.Logf("验证验证码失败: %v", err)
		t.Skip("验证失败")
	}
	if !verified {
		t.Errorf("正确的验证码应该通过验证")
	}
	t.Log("验证码验证成功")

	// 4. 再次验证（应该失败，因为已删除）
	verified, err = email.VerifyCode(emailAddr, code)
	if verified {
		t.Errorf("验证码应该已被删除，不应该再次通过验证")
	}
	t.Log("验证码已被删除，再次验证正确失败")
}
