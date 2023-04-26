package entitydb

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type EntityDB struct {
	host     string
	post     int
	user     string
	password string
	dbname   string
	db       *sql.DB
}

// Entity is a struct with same fields as database columns
type Entity struct {
	Id          int         `db:"id"`
	Name        string      `db:"name"`
	Description string      `db:"description""`
	Properties  PropertyMap `db:"properties"`
}

// PropertyMap is a type for our properties field. Note that if you have
// different kinds of entities (orders, customers, books, ...), you can
// simply re-use this type if they have a similar field.
type PropertyMap map[string]interface{}

// Value transforms the map to JSONB data or []byte
func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Scan will take the JSONB data that comes from the database
// and transform it to the entity type.
func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})

	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

func NewEntityDB(host string, port int, user, password, dbname string) EntityDB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return EntityDB{
		host:     host,
		post:     port,
		user:     user,
		password: password,
		dbname:   dbname,
		db:       db,
	}

}

func (e *EntityDB) Search(key string, value interface{}) ([]Entity, error) {
	var results []Entity
	query := "select * from entity;"

	/*stmt, err := e.db.Prepare(query)
	if err != nil {
		return results, err
	}
	defer stmt.Close()


	*/
	rows, err := e.db.Query(query)
	if err != nil {
		return results, err
	}
	defer rows.Close()
	for rows.Next() {
		var entity Entity
		if err = rows.Scan(&entity.Id, &entity.Name, &entity.Description, &entity.Properties); err != nil {
			panic(err)
		} else {
			val, ok := entity.Properties[key]
			// if the key exists
			if ok {
				// if there is a match add it to the results
				if reflect.DeepEqual(val, value) {
					results = append(results, entity)
				}
			}
		}
	}

	return results, nil
}

func (e *EntityDB) Insert(entity Entity) error {
	stmt, err := e.db.Prepare(`INSERT INTO entity(name, description, properties) VALUES($1, $2, $3) RETURNING id;`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	data, err := entity.Properties.Value()
	if err != nil {
		return err
	}

	err = stmt.QueryRow(entity.Name, entity.Description, data).Scan(&entity.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			// handle the case of no row returned.
			return errors.New("no row returned on insert")
		}
	}
	return nil
}

func (e *EntityDB) Lookup(id int) (Entity, error) {
	entity := Entity{Id: id}

	query := "SELECT name, description, properties FROM entity WHERE id = $1"
	/*stmt, err := e.db.Prepare()
	if err != nil {
		return entity, err
	}
	defer stmt.Close()
	*/

	err := e.db.QueryRow(query, entity.Id).Scan(&entity.Name, &entity.Description, &entity.Properties)

	if err != nil {
		if err == sql.ErrNoRows {
			// handle the case of no row returned.
			return entity, nil
		} else {
			return entity, err
		}
	}

	return entity, nil
}

func (e *EntityDB) Delete(id int) (int, error) {
	var n int
	stmt, err := e.db.Prepare(`DELETE FROM entity WHERE id=$1 RETURNING id`)
	if err != nil {
		return n, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&n)
	if err != nil {
		if err == sql.ErrNoRows {
			// handle the case of no row returned.
			return n, nil
		} else {
			return n, err
		}
	}

	return n, nil
}
