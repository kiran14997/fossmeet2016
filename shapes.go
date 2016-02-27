pakage main
import (
"fmt"
"os"
"strconv"
)
type Rectangle struct {
length float64
width float64
}
type Circle struct {
radius float64
}
func (r Rectangle) Area() float64 {
return r.length * r.width
}
func (r Rectangle) Perimeter() float64 {
return 2 * (r.length + r.width)
}
func (c Circle) Area() float64 {
return 3.14 * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
return 2 * c.radius
}
func main() {
shape := os.Args[1]
if shape == "circle" {
r := os.Args[2]
radius, _ := strconv.Atoi(r)
circle := Circle{float64(radius)}
area := circle.Area()
fmt.Println("area", area)
perimeter := circle.perimeter()
fmt.Println("perimeter", perimeter)
} else {
w := os.Args[2]
h := os.Args[2]
width, _ := strconv.Atoi(w)
height, _ := strconv.Atoi(h)
rectangle := Rectangle
