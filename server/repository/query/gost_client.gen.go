// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"server/model"
)

func newGostClient(db *gorm.DB, opts ...gen.DOOption) gostClient {
	_gostClient := gostClient{}

	_gostClient.gostClientDo.UseDB(db, opts...)
	_gostClient.gostClientDo.UseModel(&model.GostClient{})

	tableName := _gostClient.gostClientDo.TableName()
	_gostClient.ALL = field.NewAsterisk(tableName)
	_gostClient.Id = field.NewInt(tableName, "id")
	_gostClient.Code = field.NewString(tableName, "code")
	_gostClient.AllowEdit = field.NewInt(tableName, "allow_edit")
	_gostClient.AllowDel = field.NewInt(tableName, "allow_del")
	_gostClient.Version = field.NewInt64(tableName, "version")
	_gostClient.CreatedAt = field.NewTime(tableName, "created_at")
	_gostClient.UpdatedAt = field.NewTime(tableName, "updated_at")
	_gostClient.Name = field.NewString(tableName, "name")
	_gostClient.UserCode = field.NewString(tableName, "user_code")
	_gostClient.Key = field.NewString(tableName, "key")
	_gostClient.User = gostClientBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.SystemUser"),
		BindEmail: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.BindEmail", "model.SystemUserEmail"),
		},
	}

	_gostClient.fillFieldMap()

	return _gostClient
}

type gostClient struct {
	gostClientDo

	ALL       field.Asterisk
	Id        field.Int
	Code      field.String
	AllowEdit field.Int
	AllowDel  field.Int
	Version   field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	Name      field.String
	UserCode  field.String
	Key       field.String
	User      gostClientBelongsToUser

	fieldMap map[string]field.Expr
}

func (g gostClient) Table(newTableName string) *gostClient {
	g.gostClientDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gostClient) As(alias string) *gostClient {
	g.gostClientDo.DO = *(g.gostClientDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gostClient) updateTableName(table string) *gostClient {
	g.ALL = field.NewAsterisk(table)
	g.Id = field.NewInt(table, "id")
	g.Code = field.NewString(table, "code")
	g.AllowEdit = field.NewInt(table, "allow_edit")
	g.AllowDel = field.NewInt(table, "allow_del")
	g.Version = field.NewInt64(table, "version")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")
	g.Name = field.NewString(table, "name")
	g.UserCode = field.NewString(table, "user_code")
	g.Key = field.NewString(table, "key")

	g.fillFieldMap()

	return g
}

func (g *gostClient) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gostClient) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 11)
	g.fieldMap["id"] = g.Id
	g.fieldMap["code"] = g.Code
	g.fieldMap["allow_edit"] = g.AllowEdit
	g.fieldMap["allow_del"] = g.AllowDel
	g.fieldMap["version"] = g.Version
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["name"] = g.Name
	g.fieldMap["user_code"] = g.UserCode
	g.fieldMap["key"] = g.Key

}

func (g gostClient) clone(db *gorm.DB) gostClient {
	g.gostClientDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gostClient) replaceDB(db *gorm.DB) gostClient {
	g.gostClientDo.ReplaceDB(db)
	return g
}

type gostClientBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	BindEmail struct {
		field.RelationField
	}
}

func (a gostClientBelongsToUser) Where(conds ...field.Expr) *gostClientBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a gostClientBelongsToUser) WithContext(ctx context.Context) *gostClientBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a gostClientBelongsToUser) Session(session *gorm.Session) *gostClientBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a gostClientBelongsToUser) Model(m *model.GostClient) *gostClientBelongsToUserTx {
	return &gostClientBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type gostClientBelongsToUserTx struct{ tx *gorm.Association }

func (a gostClientBelongsToUserTx) Find() (result *model.SystemUser, err error) {
	return result, a.tx.Find(&result)
}

func (a gostClientBelongsToUserTx) Append(values ...*model.SystemUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a gostClientBelongsToUserTx) Replace(values ...*model.SystemUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a gostClientBelongsToUserTx) Delete(values ...*model.SystemUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a gostClientBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a gostClientBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type gostClientDo struct{ gen.DO }

type IGostClientDo interface {
	gen.SubQuery
	Debug() IGostClientDo
	WithContext(ctx context.Context) IGostClientDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGostClientDo
	WriteDB() IGostClientDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGostClientDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGostClientDo
	Not(conds ...gen.Condition) IGostClientDo
	Or(conds ...gen.Condition) IGostClientDo
	Select(conds ...field.Expr) IGostClientDo
	Where(conds ...gen.Condition) IGostClientDo
	Order(conds ...field.Expr) IGostClientDo
	Distinct(cols ...field.Expr) IGostClientDo
	Omit(cols ...field.Expr) IGostClientDo
	Join(table schema.Tabler, on ...field.Expr) IGostClientDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGostClientDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGostClientDo
	Group(cols ...field.Expr) IGostClientDo
	Having(conds ...gen.Condition) IGostClientDo
	Limit(limit int) IGostClientDo
	Offset(offset int) IGostClientDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGostClientDo
	Unscoped() IGostClientDo
	Create(values ...*model.GostClient) error
	CreateInBatches(values []*model.GostClient, batchSize int) error
	Save(values ...*model.GostClient) error
	First() (*model.GostClient, error)
	Take() (*model.GostClient, error)
	Last() (*model.GostClient, error)
	Find() ([]*model.GostClient, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GostClient, err error)
	FindInBatches(result *[]*model.GostClient, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GostClient) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGostClientDo
	Assign(attrs ...field.AssignExpr) IGostClientDo
	Joins(fields ...field.RelationField) IGostClientDo
	Preload(fields ...field.RelationField) IGostClientDo
	FirstOrInit() (*model.GostClient, error)
	FirstOrCreate() (*model.GostClient, error)
	FindByPage(offset int, limit int) (result []*model.GostClient, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGostClientDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g gostClientDo) Debug() IGostClientDo {
	return g.withDO(g.DO.Debug())
}

func (g gostClientDo) WithContext(ctx context.Context) IGostClientDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gostClientDo) ReadDB() IGostClientDo {
	return g.Clauses(dbresolver.Read)
}

func (g gostClientDo) WriteDB() IGostClientDo {
	return g.Clauses(dbresolver.Write)
}

func (g gostClientDo) Session(config *gorm.Session) IGostClientDo {
	return g.withDO(g.DO.Session(config))
}

func (g gostClientDo) Clauses(conds ...clause.Expression) IGostClientDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gostClientDo) Returning(value interface{}, columns ...string) IGostClientDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gostClientDo) Not(conds ...gen.Condition) IGostClientDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gostClientDo) Or(conds ...gen.Condition) IGostClientDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gostClientDo) Select(conds ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gostClientDo) Where(conds ...gen.Condition) IGostClientDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gostClientDo) Order(conds ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gostClientDo) Distinct(cols ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gostClientDo) Omit(cols ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gostClientDo) Join(table schema.Tabler, on ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gostClientDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gostClientDo) RightJoin(table schema.Tabler, on ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gostClientDo) Group(cols ...field.Expr) IGostClientDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gostClientDo) Having(conds ...gen.Condition) IGostClientDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gostClientDo) Limit(limit int) IGostClientDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gostClientDo) Offset(offset int) IGostClientDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gostClientDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGostClientDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gostClientDo) Unscoped() IGostClientDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gostClientDo) Create(values ...*model.GostClient) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gostClientDo) CreateInBatches(values []*model.GostClient, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gostClientDo) Save(values ...*model.GostClient) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gostClientDo) First() (*model.GostClient, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GostClient), nil
	}
}

func (g gostClientDo) Take() (*model.GostClient, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GostClient), nil
	}
}

func (g gostClientDo) Last() (*model.GostClient, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GostClient), nil
	}
}

func (g gostClientDo) Find() ([]*model.GostClient, error) {
	result, err := g.DO.Find()
	return result.([]*model.GostClient), err
}

func (g gostClientDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GostClient, err error) {
	buf := make([]*model.GostClient, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gostClientDo) FindInBatches(result *[]*model.GostClient, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gostClientDo) Attrs(attrs ...field.AssignExpr) IGostClientDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gostClientDo) Assign(attrs ...field.AssignExpr) IGostClientDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gostClientDo) Joins(fields ...field.RelationField) IGostClientDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gostClientDo) Preload(fields ...field.RelationField) IGostClientDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gostClientDo) FirstOrInit() (*model.GostClient, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GostClient), nil
	}
}

func (g gostClientDo) FirstOrCreate() (*model.GostClient, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GostClient), nil
	}
}

func (g gostClientDo) FindByPage(offset int, limit int) (result []*model.GostClient, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gostClientDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gostClientDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gostClientDo) Delete(models ...*model.GostClient) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gostClientDo) withDO(do gen.Dao) *gostClientDo {
	g.DO = *do.(*gen.DO)
	return g
}
