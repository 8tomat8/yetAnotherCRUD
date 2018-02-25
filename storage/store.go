package storage

import (
	"database/sql"

	"fmt"

	"net"

	"context"

	"time"

	"github.com/8tomat8/yetAnotherCRUD/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const dbName = "yetAnotherCRUD"

type storage struct {
	db     *sql.DB
	logger *logrus.Logger
}

func New(logger *logrus.Logger, host, port, user, password string) (storage, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, net.JoinHostPort(host, port), dbName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return storage{}, errors.Wrap(err, "cannot open connection to "+connString)
	}
	return storage{db, logger}, nil
}

func (s storage) Create(ctx context.Context, user *entity.User) error {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot get connection from pool")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			s.logger.Warn(errors.Wrap(err, "cannot close DB connection"))
		}
	}()

	result, err := conn.ExecContext(ctx, insertUser, user.Username, user.Password, user.Firstname, user.Lastname, user.Sex, user.Birthdate)
	if err != nil {
		return errors.Wrap(err, "cannot execute insert query")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "cannot acquire new users id")
	}

	// Why int32? It does not match to default mysql-go bindings
	user.UserID = int32(id)

	return nil
}

func (s storage) Delete(ctx context.Context, userID int) error {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot get connection from pool")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			s.logger.Warn(errors.Wrap(err, "cannot close DB connection"))
		}
	}()

	result, err := conn.ExecContext(ctx, deleteUser, userID)
	if err != nil {
		return errors.Wrap(err, "cannot execute insert query")
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "cannot get number of affected rows")
	}

	// 1 is valid until we delete users by PK
	if affected != 1 {
		return ErrNotFound
	}

	return nil
}

func (s storage) Update(ctx context.Context, user *entity.User) error {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot get connection from pool")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			s.logger.Warn(errors.Wrap(err, "cannot close DB connection"))
		}
	}()

	result, err := conn.ExecContext(ctx, updateUser, user.Username, user.Password, user.Firstname, user.Lastname, user.Sex, user.Birthdate, user.UserID)
	if err != nil {
		return errors.Wrap(err, "cannot execute update query")
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "cannot get number of affected rows")
	}

	// 1 is valid until we delete users by PK
	if affected != 1 {
		return ErrNotFound
	}

	return nil
}

func (s storage) Search(ctx context.Context, username, sex *string, age *int) ([]entity.User, error) {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get connection from pool")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			s.logger.Warn(errors.Wrap(err, "cannot close DB connection"))
		}
	}()

	rows, err := conn.QueryContext(ctx, buildSearchQuery(username, sex, age))
	if err != nil {
		return nil, errors.Wrap(err, "cannot execute update query")
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Sex, &user.Birthdate)
		if err != nil {
			if err := rows.Close(); err != nil {
				s.logger.Error(errors.Wrap(err, "cannot close sql.Rows object"))
			}
			return nil, errors.Wrap(err, "cannot scan user to struct")
		}
		users = append(users, user)
	}

	return users, nil
}

// In production it should be done using some github.com/Masterminds/squirrel for ex.
// Manual query building should die one day
func buildSearchQuery(username, sex *string, age *int) string {
	conditions := ""
	if username != nil {
		conditions += fmt.Sprintf("Username = '%s'", *username)
	}

	if sex != nil {
		if len(conditions) != 0 {
			conditions += " AND "
		}
		conditions += fmt.Sprintf("Sex = '%s'", *sex)
	}

	if age != nil {
		if len(conditions) != 0 {
			conditions += " AND "
		}
		from := time.Now().AddDate(-*age, 0, 0)
		to := from.AddDate(1, 0, 0)
		conditions += fmt.Sprintf("Birthdate BETWEEN '%s' AND '%s'", from, to)
	}

	if len(conditions) > 0 {
		conditions = " WHERE " + conditions
	}

	return fmt.Sprintf(searchUser, conditions)
}
