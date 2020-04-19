package main;
import (
"log"
"encoding/json"
"io/ioutil"
)
type student struct{
Name string `json:"name"`
College string `json:"college"`
Branch string `json:"branch"`
Address string `json:"address"`
}
var a map[int]student
func main(){
a=make(map[int]student)
b,_ := ioutil.ReadFile("student.json")
json.Unmarshal(b,&a)
for i:=0;i<len(a);i++{
log.Println("Maps Data:",i,a[i])
}
a[4]=student{"lakshmi","Vijayawada","ECE","Vij"}
log.Println("Inserted at 4:",a)
for i:=0;i<len(a);i++{
if a[i].Name=="Aparna"{
delete(a,3)
a[5]=student{"rajesh","Niet","Civil","Nallapadu"}
}
}
log.Println("Deleted list:",a)
}

