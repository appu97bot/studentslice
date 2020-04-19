package main;
import (
"log"
"encoding/json"
"io/ioutil"
"sort"
"time"
)
type student struct{
Name string `json:"name"`
College string `json:"college"`
Branch string `json:"branch"`
Address string `json:"address"`
}
var a []student;
func main(){
b,_ := ioutil.ReadFile("student.json")
json.Unmarshal(b,&a)
for i:=0;i<len(a);i++{
log.Println("Json Data is:",i,a[i])
}
sort.Slice(a, func(i, j int) bool {
  return a[i].Name < a[j].Name
})
for i:=0;i<len(a);i++{
log.Println("sorted by name:",a[i])
}
c:=student{"Kalyani","Niet","IT","Guntur"}
a=append(a,c)
for i:=0;i<len(a);i++{
log.Println("Added data :",a[i])
}
for i:=0;i<len(a);i++{
if a[i].Name=="Anusha"{
a[i]=student{}
log.Println("Dleted Data:",a[i])
}
if a[i].College=="NRT"{
a[i].College="RVR"
log.Println("Updated data:",a[i])
}
if a[i].Address=="Sattenapalli"{
a[i].Name="Vidya poluri"
log.Println("Changed Data:",a[i])
}
}
for i:=0;i<10;i++{
go insert()
}
//go insert(a,student{"Anju","Niet","IT","Guntur"},3)
//go insert(a,student{"Micky","Niet","IT","Nallapadu"},3)
go update(a,student{"Aish","BLR","BSC","Karnataka"},6)
time.Sleep(10*time.Second)
}
func insert(){
a[3]=student{"Kajol","Delhi college","BCOM","Delhi"}
log.Println("Go routine1:",a[3]);
}

//func insert(a []student, y student,z int){
//j:=0
//l:=len(a)+1
//b:=make([]student,l,l)
//for i:=0;i<l;i++{
//if i==z{
//b[i]=y;
//b[i+1]=a[j];
//i=i+1
//}else
//{
//b[i]=a[j]
//}
//j++
//}
//a=b;
//log.Println("Go routine Data:",a)
//}
func update(a []student,y student,z int){
for i:=0;i<len(a);i++{
if i==z{
a[i]=y
}
}
log.Println("Go updated routine:",a)
}
