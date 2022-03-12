package request
import(
	"fmt"

	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)
func Goodbye() {
	fmt.Println("Goodbye Aleks")
}

//func Insert() {
//	result, err := db.Exec("insert into testbd.user (name, lastName,login, password, deleted) values ('Aleks', 'Piatrou', 'PetAL', '12345', 'No')")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(result.RowsAffected())
