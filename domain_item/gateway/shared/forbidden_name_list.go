package shared

import (
	"context"
	"strings"
	"theitem/shared/infrastructure/logger"
)

type ForbiddenNameList struct {
	log logger.Logger
}

func NewForbiddenNameList(log logger.Logger) *ForbiddenNameList {
	return &ForbiddenNameList{log: log}
}

func (r *ForbiddenNameList) ExistInForbiddenNameList(ctx context.Context, name string) bool {
	r.log.Info(ctx, "called")

	// hardcoded list
	var forbiddenWords = []string{
		"gay",
		"sex",
		"lesbian",
	}

	for _, fw := range forbiddenWords {
		if fw == strings.ToLower(name) {
			return true
		}
	}

	return false
}
