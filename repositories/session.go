package repositories

import (
	"github.com/ladmakhi81/gobanks/database"
	"github.com/ladmakhi81/gobanks/entities"
)

type SessionRepository struct {
	DatabaseServer *database.DatabaseServer
}

func (repo SessionRepository) CreateSession(session *entities.Session) error {
	if err := repo.DeleteSessionByUserId(session.UserID); err != nil {
		return err
	}
	sql := `
		INSERT INTO "_sessions" ("user_id", "access_token") VALUES ($1, $2);
	`
	statement, pErr := repo.DatabaseServer.DB.Prepare(sql)
	if pErr != nil {

		return pErr
	}
	_, eErr := statement.Exec(session.UserID, session.AccessToken)
	if eErr != nil {

		return eErr
	}

	return nil
}

func (repo SessionRepository) DeleteSessionByUserId(userID int) error {
	sql := `
		DELETE FROM "_sessions" WHERE "user_id"=$1;
	`
	statement, pErr := repo.DatabaseServer.DB.Prepare(sql)
	if pErr != nil {

		return pErr
	}
	_, eErr := statement.Exec(userID)
	if eErr != nil {

		return eErr
	}

	return nil
}

func (repo SessionRepository) GetSessionByToken(accessToken string) (*entities.Session, error) {
	sql := `
		SELECT * FROM "_sessions" WHERE "access_token"=$1 LIMIT 1;
	`
	row := repo.DatabaseServer.DB.QueryRow(sql, accessToken)
	session := new(entities.Session)
	err := row.Scan(
		&session.UserID,
		&session.AccessToken,
	)
	if err != nil {

		return nil, err
	}

	return session, nil
}
