package queryMongo

import (
	"encoding/json"
	"io/ioutil"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func getAllOrderdetails() []*Orderdetail {
	orderdetailsFile, err := os.Open("orderwithdetails.json")
	if err != nil {
		panic(err)
	}
	defer orderdetailsFile.Close()

	byteValue, _ := ioutil.ReadAll(orderdetailsFile)

	orderdetails := []*Orderdetail{}

	err = json.Unmarshal(byteValue, &orderdetails)
	if err != nil {
		panic(err)
	}

	return orderdetails
}

func insertOrderdetailsRowByRow() {
	orderdetails := getAllOrderdetails()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("DBMS-Final").C("Orderdetails")
	//c.RemoveAll(nil)

	for _, orderdetail := range orderdetails {
		//fmt.Println(*car, bson.NewObjectId())
		//orderdetail.ID = bson.NewObjectId()
		err := c.Insert(orderdetail)
		if err != nil {
			panic(err)
		}
	}
}

func insertOrderdetailsBulk() {
	orderdetails := getAllOrderdetails()

	c := session.DB("DBMS-Final").C("Orderdetails")
	//c.RemoveAll(nil)
	bulk := c.Bulk()

	for _, orderdetail := range orderdetails {
		//fmt.Println(*car, bson.NewObjectId())
		//	orderdetail.ID = bson.NewObjectId()
		bulk.Insert(orderdetail)
	}
	_, err := bulk.Run()

	if err != nil {
		panic(err)
	}
}
