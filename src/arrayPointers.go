package main

import "fmt"
 
const MAX int = 3
 
func main() {
   a := []int{10,100,200}
   var i int
   //var ptr [MAX]*int;

   // for  i = 0; i < MAX; i++ {
   //    ptr[i] = &a[i] /* assign the address of integer. */
   // }
   //fmt.Println(ptr)
   for  i = 0; i < MAX; i++ {
      fmt.Printf("Value of a[%d] = %d\n", i,a[i] )
      fmt.Println(&a[i])
   }
}