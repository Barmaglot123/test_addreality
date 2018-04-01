package repositories

import (
    "github.com/jinzhu/gorm"
)


type TransactionFactory interface {
    BeginNewTransaction() Transaction
}

type transactionFactory struct {
    db *gorm.DB
}

func NewTransactionFactory(db *gorm.DB) TransactionFactory {
    return &transactionFactory{db: db}
}

func (t transactionFactory)BeginNewTransaction() Transaction {
    tx := new(transaction)
    tx.db = t.db
    tx.Begin()
    return tx
}

type Transaction interface {
    Begin()
    Commit()
    Rollback()
    DataSource() interface{}
}

type transaction struct {
    Transaction
    db *gorm.DB
    tx *gorm.DB
}

func (t *transaction)Begin() {
    t.tx = t.db.Begin()
}

func (t *transaction)Commit() {
    t.tx.Commit()
}

func (t *transaction)Rollback() {
    t.tx.Rollback()
}

func (t *transaction)DataSource() interface{} {
    return t.tx
}