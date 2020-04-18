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
func main(){
var a []student;
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
go insert(a,student{"Anju","Niet","IT","Guntur"},3)
go insert(a,student{"Micky","Niet","IT","Nallapadu"},3)
time.Sleep(10*time.Second)
}
func insert(x []student, y student,z int){
l:=len(x)+1;j:=0;
b:=make([]student,l,l);
for i:=0;i<len(b);i++{
if i==z{
b[i]=y;
b[i+1]=x[j];
i=i+1
}else
{
b[i]=x[j]
}
j++
}
log.Println("Go routine Data:",b)
}

