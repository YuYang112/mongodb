package mongodb

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)
type Persion struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}
type Men struct {
	Persons []Persion
}
func main() {
	//1、建立连接
	Session,err:= mgo.Dial("mongodb://root:123456@188.131.192.242:27017")
	if err != nil{
		fmt.Println("mgo.Dial Error:",err)
		return
	}
	defer Session.Close()
	//2、选择数据库
	db := Session.DB("test")
	//3、选择表
	collection :=db.C("person")
	temp := &Persion{
		Phone:"13301178909",
		Name:"Ale",
	}
	//4、插入数据，一次可以插入多个对象
	err = collection.Insert(&Persion{"Ale","13301178909"},temp)
	if err != nil{
		fmt.Println("collection.Insert Error:",err)
		return
	}

	//5.查询数据
	People := make(map[string]interface{})
	err = collection.Find(Persion{"Ale","13301178909"}).One(&People)
	if err != nil{
		fmt.Println("collection.Find Error:",err)
		return
	}
	fmt.Println("People===>",People)

	//6、删除所有name等于xiaomin的数据
	_,err=collection.RemoveAll(bson.M{"name":"xiaomin"})

	if err != nil{
		fmt.Println("collection.RemoveAll Error:",err)
		return
	}

	//7、更新所有name为Ale的对象name为bbb
	err = collection.Update(bson.M{"name":"Ale"},bson.M{"name":"ddd","phone":"123456789"})
	if err != nil{
		fmt.Println("collection.Update Error:",err)
		return
	}
	//7、查找全部
	//每次最多输出15条数据
	result:=Persion{}
	var persionAll Men
	iter := collection.Find(nil).Sort("_id").Skip(1).Limit(15).Iter()
	for iter.Next(&result){
		fmt.Printf("Result:%v\n",result)
		persionAll.Persons = append(persionAll.Persons,result)
	}
	fmt.Println("persionAll===>",persionAll)
	//8、集合中元素数目
	countNum,err := collection.Count()
	if err != nil{
		fmt.Println("collection.Count Error:",err)
		return
	}

	fmt.Println("countNum :",countNum)



}
