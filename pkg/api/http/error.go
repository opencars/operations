package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/opencars/httputil"
	"github.com/opencars/seedwork"

	"github.com/opencars/operations/pkg/logger"
)

func handleErr(err error) error {
	if err != nil {
		logger.Errorf("handleErr: %s", err)
	}

	var e seedwork.Error
	if errors.As(err, &e) {
		return httputil.NewError(http.StatusBadRequest, e.Error())
	}

	var vErr seedwork.ValidationError
	if errors.As(err, &vErr) {
		errMessage := make([]string, 0)
		for k, vv := range vErr.Messages {
			for _, v := range vv {
				errMessage = append(errMessage, fmt.Sprintf("%s.%s", k, v))
			}
		}

		return httputil.NewError(http.StatusBadRequest, errMessage...)
	}

	return err
}
