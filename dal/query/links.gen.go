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

	"blog/dal/model"
)

func newLinks(db *gorm.DB, opts ...gen.DOOption) links {
	_links := links{}

	_links.linksDo.UseDB(db, opts...)
	_links.linksDo.UseModel(&model.Links{})

	tableName := _links.linksDo.TableName()
	_links.ALL = field.NewAsterisk(tableName)
	_links.Id = field.NewInt64(tableName, "mid")
	_links.Title = field.NewString(tableName, "title")
	_links.Image = field.NewString(tableName, "image")
	_links.Description = field.NewString(tableName, "description")
	_links.Link = field.NewString(tableName, "link")
	_links.Sort = field.NewInt(tableName, "sort")

	_links.fillFieldMap()

	return _links
}

type links struct {
	linksDo linksDo

	ALL         field.Asterisk
	Id          field.Int64
	Title       field.String
	Image       field.String
	Description field.String
	Link        field.String
	Sort        field.Int

	fieldMap map[string]field.Expr
}

func (l links) Table(newTableName string) *links {
	l.linksDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l links) As(alias string) *links {
	l.linksDo.DO = *(l.linksDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *links) updateTableName(table string) *links {
	l.ALL = field.NewAsterisk(table)
	l.Id = field.NewInt64(table, "mid")
	l.Title = field.NewString(table, "title")
	l.Image = field.NewString(table, "image")
	l.Description = field.NewString(table, "description")
	l.Link = field.NewString(table, "link")
	l.Sort = field.NewInt(table, "sort")

	l.fillFieldMap()

	return l
}

func (l *links) WithContext(ctx context.Context) *linksDo { return l.linksDo.WithContext(ctx) }

func (l links) TableName() string { return l.linksDo.TableName() }

func (l links) Alias() string { return l.linksDo.Alias() }

func (l links) Columns(cols ...field.Expr) gen.Columns { return l.linksDo.Columns(cols...) }

func (l *links) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *links) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["mid"] = l.Id
	l.fieldMap["title"] = l.Title
	l.fieldMap["image"] = l.Image
	l.fieldMap["description"] = l.Description
	l.fieldMap["link"] = l.Link
	l.fieldMap["sort"] = l.Sort
}

func (l links) clone(db *gorm.DB) links {
	l.linksDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l links) replaceDB(db *gorm.DB) links {
	l.linksDo.ReplaceDB(db)
	return l
}

type linksDo struct{ gen.DO }

func (l linksDo) Debug() *linksDo {
	return l.withDO(l.DO.Debug())
}

func (l linksDo) WithContext(ctx context.Context) *linksDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l linksDo) ReadDB() *linksDo {
	return l.Clauses(dbresolver.Read)
}

func (l linksDo) WriteDB() *linksDo {
	return l.Clauses(dbresolver.Write)
}

func (l linksDo) Session(config *gorm.Session) *linksDo {
	return l.withDO(l.DO.Session(config))
}

func (l linksDo) Clauses(conds ...clause.Expression) *linksDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l linksDo) Returning(value interface{}, columns ...string) *linksDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l linksDo) Not(conds ...gen.Condition) *linksDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l linksDo) Or(conds ...gen.Condition) *linksDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l linksDo) Select(conds ...field.Expr) *linksDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l linksDo) Where(conds ...gen.Condition) *linksDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l linksDo) Order(conds ...field.Expr) *linksDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l linksDo) Distinct(cols ...field.Expr) *linksDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l linksDo) Omit(cols ...field.Expr) *linksDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l linksDo) Join(table schema.Tabler, on ...field.Expr) *linksDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l linksDo) LeftJoin(table schema.Tabler, on ...field.Expr) *linksDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l linksDo) RightJoin(table schema.Tabler, on ...field.Expr) *linksDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l linksDo) Group(cols ...field.Expr) *linksDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l linksDo) Having(conds ...gen.Condition) *linksDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l linksDo) Limit(limit int) *linksDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l linksDo) Offset(offset int) *linksDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l linksDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *linksDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l linksDo) Unscoped() *linksDo {
	return l.withDO(l.DO.Unscoped())
}

func (l linksDo) Create(values ...*model.Links) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l linksDo) CreateInBatches(values []*model.Links, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l linksDo) Save(values ...*model.Links) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l linksDo) First() (*model.Links, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Links), nil
	}
}

func (l linksDo) Take() (*model.Links, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Links), nil
	}
}

func (l linksDo) Last() (*model.Links, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Links), nil
	}
}

func (l linksDo) Find() ([]*model.Links, error) {
	result, err := l.DO.Find()
	return result.([]*model.Links), err
}

func (l linksDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Links, err error) {
	buf := make([]*model.Links, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l linksDo) FindInBatches(result *[]*model.Links, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l linksDo) Attrs(attrs ...field.AssignExpr) *linksDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l linksDo) Assign(attrs ...field.AssignExpr) *linksDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l linksDo) Joins(fields ...field.RelationField) *linksDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l linksDo) Preload(fields ...field.RelationField) *linksDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l linksDo) FirstOrInit() (*model.Links, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Links), nil
	}
}

func (l linksDo) FirstOrCreate() (*model.Links, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Links), nil
	}
}

func (l linksDo) FindByPage(offset int, limit int) (result []*model.Links, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l linksDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l linksDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l linksDo) Delete(models ...*model.Links) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *linksDo) withDO(do gen.Dao) *linksDo {
	l.DO = *do.(*gen.DO)
	return l
}
