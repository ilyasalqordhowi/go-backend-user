package models

import (
	"backend/lib"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" form:"email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
	Username string `json:"username" form:"username" db:"username"`
}

func FindAllUsers()[]Users {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
	context.Background(),
	`SELECT * FROM "users"`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	}

	return users
	 
}
func FindOneUsers(id int)Users{
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
	context.Background(),
	`SELECT * FROM "users" where `,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	} 
	
	user := Users{}
	for _, v := range users{
		if v.Id == id {
			user = v
		}
	}
	return user
}
func CreateUser(user Users) error {
    db := lib.DB()
    defer db.Close(context.Background())

    _, err := db.Exec(
        context.Background(),
        `INSERT INTO "users" (email, password, username) VALUES ($1, $2, $3)`,
        user.Email, user.Password, user.Username,
    )

    if err != nil {
        return fmt.Errorf("failed to execute insert")
    }

    return nil
}
func DeleteUsers(id int) error {
    db := lib.DB()
    defer db.Close(context.Background())

    commandTag, err := db.Exec(
        context.Background(),
        `DELETE FROM "users" WHERE id = $1`,
        id,
    )

    if err != nil {
        return fmt.Errorf("failed to execute delete")
    }

    if commandTag.RowsAffected() == 0 {
        return fmt.Errorf("no user found")
    }

    return nil
}
func EditUser(email string, username string, password string, id string) {
    db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4`

    db.Exec(context.Background(), dataSql, email, username, password, id)

}
	