package tracing

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleErrorForGin(ctx *gin.Context, err error) {
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}
	log.WithFields(log.Fields{
		"path":      ctx.Request.URL.Path,
		"requestId": ctx.Request.Header.Get("X-Request-ID"),
	}).Error("Error " + err.Error())
}

func HandleError(err error) {
	sentry.CaptureException(err)
	log.Error("Error " + err.Error())
}
