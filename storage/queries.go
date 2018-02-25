package storage

const (
	// In production it should be done using some github.com/Masterminds/squirrel for ex.
	insertUser = "INSERT INTO yetAnotherCRUD.users (Username,Password,Firstname,Lastname,Sex,Birthdate) VALUES (?,?,?,?,?,?)"
	updateUser = "UPDATE yetAnotherCRUD.users SET Username = ?, Password = ?,Firstname = ?,Lastname = ?,Sex = ?,Birthdate = ? WHERE ID = ?"
	deleteUser = "DELETE FROM users WHERE ID = ?"
	searchUser = "SELECT users.ID,users.Username,users.Password,users.Firstname,users.Lastname,users.Sex,users.Birthdate FROM yetAnotherCRUD.users %s ORDER BY users.Username ASC"
)
