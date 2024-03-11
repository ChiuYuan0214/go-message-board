package services

import (
	"fmt"
	"net/http"
	"security/utils"
)

func UpdateColumnsById(data interface{}, id *uint64) (string, int) {
	cols, args := utils.ConstructParamsFromStruct(data)
	if cols == "" {
		return "column name invalid.", http.StatusBadRequest
	}

	stmt := fmt.Sprintf("update users set %s where user_id = ?", cols)
	args = append(args, *id)
	sqlRes, err := connPool.Exec(stmt, args...)
	if err != nil {
		return "failed to update columns.", http.StatusInternalServerError
	}
	count, _ := sqlRes.RowsAffected()
	if count < 1 {
		return "entity not exist or nothing to update.", http.StatusBadRequest
	}
	return "", 0
}
