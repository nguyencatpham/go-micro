// Package sql is the micro data model implementation
package sql

import (
	"github.com/nguyencatpham/go-micro/codec/json"
	"github.com/nguyencatpham/go-micro/model"
	"github.com/nguyencatpham/go-micro/store"
	"github.com/nguyencatpham/go-micro/store/memory"
	memsync "github.com/nguyencatpham/go-micro/sync/memory"
)

type sqlModel struct {
	options model.Options
}

func (m *sqlModel) Init(opts ...model.Option) error {
	for _, o := range opts {
		o(&m.options)
	}
	return nil
}

func (m *sqlModel) NewEntity(name string, value interface{}) model.Entity {
	// TODO: potentially pluralise name for tables
	return newEntity(name, value, m.options.Codec)
}

func (m *sqlModel) Create(e model.Entity) error {
	// lock on the name of entity
	if err := m.options.Sync.Lock(e.Name()); err != nil {
		return err
	}
	// TODO: deal with the error
	defer m.options.Sync.Unlock(e.Name())

	// TODO: potentially add encode to entity?
	v, err := m.options.Codec.Marshal(e.Value())
	if err != nil {
		return err
	}

	// TODO: include metadata and set database
	return m.options.Store.Write(&store.Record{
		Key:   e.Id(),
		Value: v,
	}, store.WriteTo(m.options.Database, e.Name()))
}

func (m *sqlModel) Read(opts ...model.ReadOption) ([]model.Entity, error) {
	var options model.ReadOptions
	for _, o := range opts {
		o(&options)
	}
	// TODO: implement the options that allow querying
	return nil, nil
}

func (m *sqlModel) Update(e model.Entity) error {
	// TODO: read out the record first, update the fields and store

	// lock on the name of entity
	if err := m.options.Sync.Lock(e.Name()); err != nil {
		return err
	}
	// TODO: deal with the error
	defer m.options.Sync.Unlock(e.Name())

	// TODO: potentially add encode to entity?
	v, err := m.options.Codec.Marshal(e.Value())
	if err != nil {
		return err
	}

	// TODO: include metadata and set database
	return m.options.Store.Write(&store.Record{
		Key:   e.Id(),
		Value: v,
	}, store.WriteTo(m.options.Database, e.Name()))
}

func (m *sqlModel) Delete(opts ...model.DeleteOption) error {
	var options model.DeleteOptions
	for _, o := range opts {
		o(&options)
	}
	// TODO: implement the options that allow deleting
	return nil
}

func (m *sqlModel) String() string {
	return "sql"
}

func NewModel(opts ...model.Option) model.Model {
	options := model.Options{
		Codec: new(json.Marshaler),
		Sync:  memsync.NewSync(),
		Store: memory.NewStore(),
	}

	return &sqlModel{
		options: options,
	}
}
