package main;
import (
"log"
"encoding/json"
"io/ioutil"
"sort"
)
type student struct{
Name string `json:"name"`
College string `json:"college"`
Branch string `json:"branch"`
Address string `json:"address"`
}
func main(){
var a []student;
b,_ := ioutil.ReadFile("student.json")
json.Unmarshal(b,&a)
log.Println("Json Data is:",a)
sort.Slice(a, func(i, j int) bool {
  return a[i].Name < a[j].Name
})
log.Println("sorted by name:",a)
c:=student{"Kalyani","Niet","IT","Guntur"}
a=append(a,c)
log.Println("Added data :",a)
for i:=0;i<len(a);i++{
if a[i].Name=="Anusha"{
a[i].Name=" "
}
if a[i].College=="NRT"{
a[i].College="RVR"
}
if a[i].Address=="Sattenapalli"{
a[i].Name="Vidya poluri"
}
}
log.Println("Deleted,updated, insertat  Data is :",a);
}
