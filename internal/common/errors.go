package common

import (
	"bytes"
	"errors"
	"fmt"
)

const unhandledErrMsg = "不明なエラーおきました"

var ErrRequired = errors.New("required")
var ErrInvalid = errors.New("invalid")

//go:generate go run golang.org/x/tools/cmd/stringer -type=ErrCode

type ErrCode uint32

const (
	ErrCodeInternal ErrCode = iota // internal error
	ErrCodeInvalid                 // validation failed
	ErrCodeNotFound                // entity does not exist
)

type Error struct {
	//Code エラーコード
	Code ErrCode
	//Message 人間がわかりやすいメッセージ
	Message string
	//Op オペレーション
	Op  string
	Err error
}

func (e *Error) String() string {
	return e.Error()
}

func ErrorCode(err error) ErrCode {
	if err == nil {
		return ErrCodeInternal
	}
	var er *Error
	if !errors.As(err, &er) {
		return ErrCodeInternal
	}
	// Unknownじゃないエラーコードまで深ぼって取得
	if er.Code != ErrCodeInternal {
		return er.Code
	}
	return ErrorCode(er.Err)
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}
	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		fmt.Fprintf(&buf, "<%s> ", e.Code.String())
		buf.WriteString(e.Message)
	}
	return buf.String()
}

// ErrorMessage メッセージ取得
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	var er *Error
	if !errors.As(err, &er) {
		return unhandledErrMsg
	}
	if er.Message != "" {
		return er.Message
	}
	if er.Err == nil {
		return unhandledErrMsg
	}
	return ErrorMessage(er.Err)
}
