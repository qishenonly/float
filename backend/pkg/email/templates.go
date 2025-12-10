package email

// VerificationCodeTemplate 验证码邮件模板
const VerificationCodeTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>验证码</title>
</head>
<body style="margin: 0; padding: 0; background-color: #F2F4F8; font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;">
    <table border="0" cellpadding="0" cellspacing="0" width="100%" style="background: linear-gradient(135deg, #eff6ff 0%, #fff1f2 100%);">
        <tr>
            <td align="center" style="padding: 40px 0;">
                <table border="0" cellpadding="0" cellspacing="0" width="600" style="background-color: #ffffff; border-radius: 24px; box-shadow: 0 10px 30px rgba(0,0,0,0.05); overflow: hidden; border: 4px solid #ffffff;">
                    
                    <tr>
                        <td align="center" style="padding: 40px 0 20px 0;">
                            <div style="background: linear-gradient(135deg, #4F46E5 0%, #7c3aed 100%); width: 80px; height: 80px; border-radius: 20px; display: flex; align-items: center; justify-content: center; font-size: 36px; font-weight: bold; color: #ffffff;">F</div>
                            <h1 style="color: #1f2937; font-size: 24px; margin: 10px 0 0 0; font-weight: 800;">Float Island</h1>
                        </td>
                    </tr>

                    <tr>
                        <td style="padding: 0 40px;">
                            <p style="color: #4b5563; font-size: 16px; line-height: 24px; margin-bottom: 20px; text-align: center;">
                                Hi, 岛民 👋
                            </p>
                            <p style="color: #6b7280; font-size: 14px; line-height: 24px; margin-bottom: 30px; text-align: center;">
                                您正在进行登录或安全验证操作，请使用以下验证码完成验证。<br>验证码 10 分钟内有效。
                            </p>
                            
                            <table border="0" cellpadding="0" cellspacing="0" width="100%">
                                <tr>
                                    <td align="center">
                                        <div style="background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%); border-radius: 16px; padding: 20px 40px; border: 2px dashed #cbd5e1; display: inline-block;">
                                            <span style="font-family: 'Courier New', Courier, monospace; font-size: 32px; font-weight: bold; letter-spacing: 8px; color: #4F46E5;">{{CODE}}</span>
                                        </div>
                                    </td>
                                </tr>
                            </table>

                            <p style="color: #9ca3af; font-size: 12px; line-height: 20px; margin-top: 30px; text-align: center;">
                                如果这不是您本人的操作，请忽略此邮件，您的账户是安全的。
                            </p>
                        </td>
                    </tr>

                    <tr>
                        <td style="padding-top: 40px;">
                            <div style="height: 6px; background: linear-gradient(90deg, #4F46E5, #ec4899); width: 100%;"></div>
                        </td>
                    </tr>
                </table>

                <table border="0" cellpadding="0" cellspacing="0" width="100%">
                    <tr>
                        <td align="center" style="padding-top: 20px; color: #9ca3af; font-size: 12px;">
                            &copy; 2025 Float Island. 构建你的财富岛屿 🏝️
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`

// PasswordResetTemplate 密码重置邮件模板
const PasswordResetTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>重置密码</title>
</head>
<body style="margin: 0; padding: 0; background-color: #F2F4F8; font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;">
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
        <tr>
            <td align="center" style="padding: 40px 0;">
                <table border="0" cellpadding="0" cellspacing="0" width="600" style="background-color: #ffffff; border-radius: 24px; border: 1px solid #e5e7eb; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);">
                    
                    <tr>
                        <td align="center" style="padding: 40px 0 0 0;">
                            <div style="background-color: #eff6ff; width: 64px; height: 64px; border-radius: 50%; text-align: center; line-height: 64px; display: inline-block; font-size: 32px;">
                                🔒
                            </div>
                        </td>
                    </tr>

                    <tr>
                        <td align="center" style="padding: 24px 40px 40px 40px;">
                            <h2 style="color: #1f2937; font-size: 22px; font-weight: 800; margin: 0 0 16px 0;">重置您的密码</h2>
                            
                            <p style="color: #4b5563; font-size: 15px; line-height: 24px; margin-bottom: 30px;">
                                我们收到了重置您浮岛账户密码的请求。点击下方按钮即可设置新密码：
                            </p>

                            <a href="{{RESET_LINK}}" style="background: linear-gradient(90deg, #4F46E5 0%, #7c3aed 100%); color: #ffffff; text-decoration: none; padding: 14px 32px; border-radius: 12px; font-weight: bold; font-size: 14px; display: inline-block; box-shadow: 0 4px 14px rgba(79, 70, 229, 0.4);">
                                重置密码
                            </a>

                            <p style="color: #6b7280; font-size: 13px; margin-top: 30px; margin-bottom: 0;">
                                或者复制以下链接到浏览器打开：
                            </p>
                            <p style="margin-top: 5px; word-break: break-all;">
                                <a href="{{RESET_LINK}}" style="color: #4F46E5; font-size: 13px; text-decoration: underline;">{{RESET_LINK}}</a>
                            </p>

                            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="margin-top: 30px;">
                                <tr>
                                    <td style="background-color: #fff1f2; border-radius: 12px; padding: 16px; border: 1px solid #fecdd3;">
                                        <p style="color: #9f1239; font-size: 12px; margin: 0; line-height: 18px;">
                                            <strong>安全提示：</strong><br>
                                            如果您没有请求重置密码，请忽略此邮件。链接将在 1 小时后失效。
                                        </p>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>
                
                <p style="color: #9ca3af; font-size: 12px; margin-top: 20px;">
                    Float Island Security Team
                </p>
            </td>
        </tr>
    </table>
</body>
</html>`

// WelcomeTemplate 欢迎邮件模板
const WelcomeTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>欢迎登岛</title>
</head>
<body style="margin: 0; padding: 0; background-color: #F2F4F8; font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;">
    <table border="0" cellpadding="0" cellspacing="0" width="100%" style="background-color: #F2F4F8;">
        <tr>
            <td align="center" style="padding: 40px 0;">
                <table border="0" cellpadding="0" cellspacing="0" width="600" style="background-color: #ffffff; border-radius: 24px; overflow: hidden; border: 1px solid #e5e7eb;">
                    
                    <tr>
                        <td align="center" style="background: linear-gradient(135deg, #e0f2fe 0%, #dbeafe 100%); padding: 40px 0;">
                            <div style="font-size: 48px;">🎉</div>
                            <h1 style="color: #1f2937; font-size: 28px; font-weight: 800; margin: 10px 0 0 0;">欢迎来到浮岛！</h1>
                        </td>
                    </tr>

                    <tr>
                        <td style="padding: 40px;">
                            <p style="color: #4b5563; font-size: 16px; line-height: 26px; margin-bottom: 24px;">
                                亲爱的 {{DISPLAY_NAME}}，<br><br>
                                很高兴你能加入我们！浮岛不仅仅是一个记账工具，更是你构建财富大厦的基石。从今天开始，我们将陪伴你记录每一笔收支，见证每一次积累。
                            </p>
                            
                            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="margin-bottom: 30px;">
                                <tr>
                                    <td width="30" valign="top" style="padding-bottom: 10px;">✨</td>
                                    <td style="color: #6b7280; font-size: 14px; padding-bottom: 10px;"><strong>极速记账：</strong>3秒完成一笔记录，自动分类。</td>
                                </tr>
                                <tr>
                                    <td width="30" valign="top" style="padding-bottom: 10px;">📊</td>
                                    <td style="color: #6b7280; font-size: 14px; padding-bottom: 10px;"><strong>资产全景：</strong>资金与负债一目了然。</td>
                                </tr>
                                <tr>
                                    <td width="30" valign="top" style="padding-bottom: 10px;">☁️</td>
                                    <td style="color: #6b7280; font-size: 14px; padding-bottom: 10px;"><strong>云端同步：</strong>数据永不丢失，多端实时同步。</td>
                                </tr>
                            </table>

                            <table border="0" cellpadding="0" cellspacing="0" width="100%">
                                <tr>
                                    <td align="center">
                                        <a href="https://floatisland.app" style="background: linear-gradient(90deg, #4F46E5 0%, #7c3aed 100%); color: #ffffff; text-decoration: none; padding: 14px 32px; border-radius: 50px; font-weight: bold; font-size: 16px; display: inline-block; box-shadow: 0 4px 14px rgba(79, 70, 229, 0.4);">
                                            记下第一笔账
                                        </a>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>

                    <tr>
                        <td style="background-color: #f9fafb; padding: 24px; text-align: center;">
                            <p style="color: #9ca3af; font-size: 12px; margin: 0;">
                                遇到问题？<a href="#" style="color: #4F46E5; text-decoration: none;">查看帮助中心</a> 或直接回复此邮件。
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`
