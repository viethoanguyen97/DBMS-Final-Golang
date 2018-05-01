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

var session *mgo.Session

func init() {
	var err error
	session, err = mgo.Dial("mongodb://127.0.0.1:27017/DBMS-Final")

	if err != nil {
		panic(err)
	}
	//defer session.Close()

	fmt.Println("Connect database successfully")

}

func closeDB() {
	session.Close()
}

