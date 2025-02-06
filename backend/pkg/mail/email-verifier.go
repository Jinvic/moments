package mail

import (
	"errors"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/rs/zerolog"
)

var (
	verifier = emailverifier.NewVerifier()
)

func VerifyEmail(logger zerolog.Logger, email string) error {

	ret, err := verifier.Verify(email)
	if err != nil {
		logger.Debug().Msgf("email-verifier: 验证邮箱失败 %v", err.Error())
		return errors.New("验证邮箱失败" + err.Error())
	}

	if !ret.Syntax.Valid {
		logger.Debug().Msgf("email-verifier: 邮箱格式不正确")
		return errors.New("邮箱格式不正确")
	}

	// if ret.Reachable == "unknown" {
	// 	return errors.New("邮箱无法验证")
	// }

	// if ret.RoleAccount || ret.Disposable || ret.Free {
	// 	return errors.New("邮箱是临时邮箱")
	// }
	return nil

}

func GetDomain(email string) string {
	return verifier.ParseAddress(email).Domain
}
