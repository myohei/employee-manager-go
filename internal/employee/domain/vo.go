package domain

import (
	"github.com/google/uuid"
	"github.com/myohei/employee-manager-go/internal/common"
	"strconv"
	"time"
)

type EmployeeID uuid.UUID

func (e EmployeeID) String() string {
	return uuid.UUID(e).String()
}

func EmployeeIDString(s string) (EmployeeID, error) {
	op := "EmployeeIDString"
	if s == "" {
		return EmployeeID{}, &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "IDは必須です",
			Op:      op,
			Err:     common.ErrRequired,
		}
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return EmployeeID{}, &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "不正なフォーマットです",
			Op:      op,
			Err:     err,
		}
	}
	return EmployeeID(id), nil
}

type Name string

func (n Name) String() string {
	return string(n)
}

func NameString(s string) (Name, error) {
	if s == "" {
		return "", &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "名前は必須です",
			Op:      "NameString",
			Err:     common.ErrRequired,
		}
	}
	return Name(s), nil
}

const birthdayFmt = "2006-01-02"

type Birthday time.Time

func (b Birthday) Time() time.Time {
	return time.Time(b)
}

func (b Birthday) Age() uint {
	duration := time.Since(b.Time())
	return uint(duration.Seconds() / 31207680)
}

func BirthdayString(s string) (Birthday, error) {
	op := "BirthdayString"
	if s == "" {
		return Birthday{}, &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "誕生日は必須です",
			Op:      op,
			Err:     common.ErrRequired,
		}
	}
	d, err := time.Parse(birthdayFmt, s)
	if err != nil {
		return Birthday{}, &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "不正なフォーマットです",
			Op:      op,
			Err:     err,
		}
	}
	twentyYearsAgo := time.Now().AddDate(-20, 0, 0)
	if d.After(twentyYearsAgo) {
		return Birthday{}, &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "20歳以上じゃないと登録できません",
			Op:      op,
			Err:     common.ErrInvalid,
		}
	}
	return Birthday(d), nil
}

type Salary uint

func (s Salary) Uint() uint {
	return uint(s)
}

func SalaryString(s string) (Salary, error) {
	op := "SalaryString"
	if s == "" {
		return 0, &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "給与は必須です",
			Op:      op,
			Err:     common.ErrRequired,
		}
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, &common.Error{
			Code:    common.ErrCodeInvalid,
			Message: "不正な数字です",
			Op:      op,
			Err:     common.ErrInvalid,
		}
	}
	return Salary(i), nil
}

//go:generate go run github.com/alvaroloes/enumer -type=Gender -linecomment

type Gender int

const (
	GenderMale   Gender = iota // 男性
	GenderFemale               //女性
)
