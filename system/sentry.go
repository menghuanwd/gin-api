package system

import (
	"github.com/getsentry/sentry-go"
	"gin-api/exceptions"
)

func Sentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	})

	exceptions.Error500(err)
}