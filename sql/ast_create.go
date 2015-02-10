package sql

type Create interface {
	IStatement
}

type CreateTable struct {
	Table *TableInfo
}

func (*CreateTable) Statement() {}

type CreateIndex struct{}

func (*CreateIndex) Statement() {}

type CreateDatabase struct{}

func (*CreateDatabase) Statement() {}

type CreateView struct{}

func (*CreateView) Statement() {}

type CreateUser struct{}

func (*CreateUser) Statement() {}

type CreateLog struct{}

func (*CreateLog) Statement() {}

type CreateTablespace struct{}

func (*CreateTablespace) Statement() {}

type CreateServer struct{}

func (*CreateServer) Statement() {}
