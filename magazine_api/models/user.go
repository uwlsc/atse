package models

import (
	"errors"
	"fmt"

	"github.com/jackc/pgtype"
)

type UserRole []string

type UserBase struct {
	Name          *string  `json:"name"`
	Email         *string  `json:"email"`
	Role          UserRole `json:"role"`
	Password      *string  `json:"password"`
	ContactNumber *string  `json:"contact_number"`
}

// The User Model
type User struct {
	Base
	BaseDate
	UserBase
}

// DecodeText is Decoder for User role, to be used for database
func (u *UserRole) DecodeText(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		return errors.New("NULL values can't be decoded")
	}
	var text []pgtype.GenericText
	for _, val := range *u {
		fmt.Println(u, "gggg<<>>>")
		text = append(text, pgtype.GenericText{String: val, Status: pgtype.Present})
	}

	enum := pgtype.EnumArray{Elements: text}
	fmt.Println(enum, "fff>>>>", string(src), text)
	return enum.DecodeText(ci, src)
}

// EncodeText is Encoder for User role, to be used for database
func (u UserRole) EncodeText(ci *pgtype.ConnInfo, buf []byte) (newBuf []byte, err error) {
	var text []pgtype.GenericText
	for _, val := range u {
		text = append(text, pgtype.GenericText{String: val, Status: pgtype.Present})
	}

	enum := pgtype.EnumArray{
		Elements:   text,
		Status:     pgtype.Present,
		Dimensions: []pgtype.ArrayDimension{{Length: int32(len(text)), LowerBound: 1}},
	}
	return enum.EncodeText(ci, buf)
}
