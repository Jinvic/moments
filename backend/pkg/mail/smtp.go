package mail

import (
	"crypto/tls"
	"fmt"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/rs/zerolog"
)

func GetSMTPClient(logger zerolog.Logger, host, port, username, password string) (*smtp.Client, error) {
	// TLS配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// 连接到SMTP服务器
	logger.Debug().Msgf("smtp: 尝试DialTLS连接...")
	client, err := smtp.DialTLS(host+":"+port, tlsConfig)
	if err != nil {
		logger.Debug().Msgf("smtp: DialTLS连接失败: %v", err)
		logger.Debug().Msgf("smtp: 尝试DialStartTLS连接...")
		client, err = smtp.DialStartTLS(host+":"+port, tlsConfig)
		if err != nil {
			logger.Debug().Msgf("smtp: DialStartTLS连接失败: %v", err)
			return nil, fmt.Errorf("无法连接到SMTP服务器: %v", err)
		}
	}

	// 测试HELO/EHLO命令
	if err := client.Hello("localhost"); err != nil {
		logger.Debug().Msgf("smtp: SMTP服务器不可用: %v", err)
		return nil, fmt.Errorf("SMTP服务器不可用: %v", err)
	}

	// 测试认证
	auth := sasl.NewPlainClient("", username, password)
	if err := client.Auth(auth); err != nil {
		logger.Debug().Msgf("smtp: SMTP认证失败: %v", err)
		return nil, fmt.Errorf("SMTP认证失败: %v", err)
	}

	return client, nil
}
