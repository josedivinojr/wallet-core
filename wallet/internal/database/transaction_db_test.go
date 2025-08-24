package database

import (
	"database/sql"
	"testing"

	"josedivinojr/wallet/internal/entity"

	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	clientDB      *ClientDB
	accountDB     *AccountDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance INTEGER, created_at DATE)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount INTEGER, created_at DATE)")
	s.transactionDB = NewTransactionDB(db)
	s.clientDB = NewClientDb(db)
	s.accountDB = NewAccountDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE transactions")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	client1, _ := entity.NewClient("John", "j@j.com")
	s.clientDB.Save(client1)

	account1 := entity.NewAccount(client1)
	account1.Credit(1000)
	s.accountDB.Save(account1)

	client2, _ := entity.NewClient("Jane", "jane@doe.com")
	s.clientDB.Save(client2)

	account2 := entity.NewAccount(client2)
	s.accountDB.Save(account2)

	transaction, err := entity.NewTransaction(account1, account2, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
