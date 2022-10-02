package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Alias for any slice
type Messages = []any

type Error struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

// error makes it compatible with the error interface
func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

// A struct to return normal response
type Response struct {
	Code     int      `json:"code"`
	Messages Messages `json:"messages"`
	Data     any      `json:"data,omitempty"`
	Meta	 any      `json:"meta,omitempty"`
}

// nothiing to describe this fucking variable
var IsProduction bool

// Default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Response{
		Code: fiber.StatusInternalServerError,
	}

	// handle errors
	if c, ok := err.(validator.ValidationErrors); ok {
		resp.Code = fiber.StatusUnprocessableEntity
		resp.Messages = Messages{removeTopStruct(c.Translate(trans))}
	} else if c, ok := err.(*fiber.Error); ok {
		resp.Code = c.Code
		resp.Messages = Messages{c.Message}
	} else if c, ok := err.(*Error); ok {
		resp.Code = c.Code
		resp.Messages = Messages{c.Message}

		// for ent and other errors
		if resp.Messages == nil {
			resp.Messages = Messages{err}
		}
	} else {
		resp.Messages = Messages{err.Error()}
	}

	if !IsProduction {
		log.Error().Err(err).Msg("From: Fiber's error handler")
	}

	return Resp(c, resp)
}

// function to return pretty json response
func Resp(c *fiber.Ctx, resp Response) error {
	// set status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(resp.Code)
	// return response
	return c.JSON(resp)
}

// remove unecessary fields from validator message
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}

	for field, msg := range fields {
		stripStruct := field[strings.Index(field, ".")+1:]
		res[stripStruct] = msg
	}

	return res
}
