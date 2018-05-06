package queryMongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

/*type Device struct {
    Id           bson.ObjectId `bson:"_id" json:"_id,omitempty"`
    UserId       string        `bson:"userId" json:"userId"`
    CategorySlug string        `bson:"categorySlug" json:"categorySlug"`
    CreatedAt    time.Time     `bson:"createdAt" json:"createdAt"`
    ModifiedAt   time.Time     `bson:"modifiedAt" json:"modifiedAt"`
    BrandId      int           `bson:"brandId" json:"brandId"`
    Category     string        `bson:"category" json:"category"`
}
*/

var Session *mgo.Session

func init() {
	var err error
	Session, err = mgo.Dial("mongodb://127.0.0.1:27017/DBMSFinal")

	if err != nil {
		panic(err)
	}
	//defer session.Close()

	fmt.Println("Connect database MongoDB successfully")

}

func EnsureIndex() {
	session := Session.Copy()
	defer session.Close()

	c := session.DB("DBMSFinal").C("Cars")

	index := mgo.Index{
		Key:        []string{"car_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	c = session.DB("DBMSFinal").C("Customers")

	index = mgo.Index{
		Key:        []string{"customer_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	c = session.DB("DBMSFinal").C("Orders")

	index = mgo.Index{
		Key:        []string{"order_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	c = session.DB("DBMSFinal").C("Orderdetails")

	/*	index = mgo.Index{
			Key:        []string{"order_id", "car_id"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}

		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}
	*/
}

func CloseDB() {
	Session.Close()
}
