package main
import (
	"database/sql"   
	"fmt"   
	"github.com/pkg/errors"
)

func GetSql() error {
	return errors.Wrap(sql.ErrNoRows, "GetSql failed")
}
func Biz() error {
	return errors.WithMessage(GetSql(), "Biz failed")
}
func main() {
   	err := Biz()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return   
	}
	if err != nil {
		fmt.Printf("service: %+v\n", err)
	}
}


