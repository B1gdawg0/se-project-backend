package utils

import (
	"errors"

	"github.com/emicklei/pgtalk/convert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToUUID(str string) (*pgtype.UUID, error){
	if !CheckUUID(str){
		return nil, errors.New("can't convert this string to uuid")
	}

	uuid := convert.StringToUUID(str)

	return &uuid, nil
}

func CheckUUID(id string)(bool){
	if id == "" {
		return false
	}

	if _, err := uuid.Parse(id); err!=nil {

		return false
	}

	return true
}


func GenerateUUID() pgtype.UUID{
	return convert.StringToUUID(uuid.NewString())
}