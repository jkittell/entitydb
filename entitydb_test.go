package entitydb

import (
	"github.com/brianvoe/gofakeit"
	_ "github.com/lib/pq"
	"reflect"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func newEntity() Entity {
	name := gofakeit.Name()
	description := gofakeit.JobDescriptor()
	properties := PropertyMap{
		"email": gofakeit.Email(),
		"phone": gofakeit.Phone(),
		//"address": gofakeit.Address(),
	}

	return Entity{
		Id:          0,
		Name:        name,
		Description: description,
		Properties:  properties,
	}

}

func TestNewEntity(t *testing.T) {
	newEntity()
}

func BenchmarkNewEntity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newEntity()
	}
}

func TestEntityDB_Search(t *testing.T) {
	db := NewEntityDB(host, port, user, password, dbname)
	e := newEntity()
	err := db.Insert(e)
	if err != nil {
		t.Fail()
	}

	results, err := db.Search("email", e.Properties["email"])
	if err != nil {
		t.Fail()
	}

	if len(results) == 0 {
		t.Fail()
	}

}

func BenchmarkEntityDB_Search(b *testing.B) {
	db := NewEntityDB(host, port, user, password, dbname)

	for n := 0; n < b.N; n++ {
		e := newEntity()
		err := db.Insert(e)
		if err != nil {
			b.Fail()
		}

		results, err := db.Search("email", e.Properties["email"])
		if err != nil {
			b.Fail()
		}
		if len(results) == 0 {
			b.Fail()
		}
	}
}

func TestEntityDB_Insert(t *testing.T) {
	db := NewEntityDB(host, port, user, password, dbname)

	err := db.Insert(newEntity())
	if err != nil {
		t.Fail()
	}
}

func BenchmarkEntityDB_Insert(b *testing.B) {
	db := NewEntityDB(host, port, user, password, dbname)

	for n := 0; n < b.N; n++ {
		e := newEntity()
		err := db.Insert(e)
		if err != nil {
			b.Fail()
		}
	}
}

func TestEntityDB_Lookup(t *testing.T) {
	db := NewEntityDB(host, port, user, password, dbname)

	e := newEntity()
	err := db.Insert(e)
	if err != nil {
		t.Fail()
	}

	result, err := db.Lookup(e.Id)
	if err != nil {
		t.Fail()
	}

	if reflect.DeepEqual(result, e) {
		t.Fail()
	}

}

func BenchmarkEntityDB_Lookup(b *testing.B) {
	db := NewEntityDB(host, port, user, password, dbname)

	for n := 0; n < b.N; n++ {
		e := newEntity()
		err := db.Insert(e)
		if err != nil {
			b.Fail()
		}

		result, err := db.Lookup(e.Id)
		if err != nil {
			b.Fail()
		}
		if reflect.DeepEqual(result, e) {
			b.Fail()
		}
	}
}

func TestEntityDB_Delete(t *testing.T) {
	db := NewEntityDB(host, port, user, password, dbname)

	e := newEntity()
	err := db.Insert(e)
	if err != nil {
		t.Fail()
	}

	deletedId, err := db.Delete(e.Id)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(deletedId, e.Id) {
		t.Fail()
	}

	e, err = db.Lookup(e.Id)
	if err != nil {
		t.Fail()
	}

	if e.Name != "" {
		t.Fail()
	}

	if e.Description != "" {
		t.Fail()
	}

	if len(e.Properties) > 0 {
		t.Fail()
	}

}

func BenchmarkEntityDB_Delete(b *testing.B) {
	db := NewEntityDB(host, port, user, password, dbname)

	for n := 0; n < b.N; n++ {
		e := newEntity()
		err := db.Insert(e)
		if err != nil {
			b.Fail()
		}

		deletedId, err := db.Delete(e.Id)
		if err != nil {
			b.Fail()
		}

		if !reflect.DeepEqual(deletedId, e.Id) {
			b.Fail()
		}

		e, err = db.Lookup(e.Id)
		if err != nil {
			b.Fail()
		}

		if e.Name != "" {
			b.Fail()
		}

		if e.Description != "" {
			b.Fail()
		}

		if e.Properties != nil {
			b.Fail()
		}
	}
}
