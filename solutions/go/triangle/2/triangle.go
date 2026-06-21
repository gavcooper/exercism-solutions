// Package triangle determines what type of triangle three side lengths must form.
package triangle

type Kind string 

const (
    NaT = "not"
    Equ = "equilateral"
    Iso = "isosceles"
    Sca = "scalene"
)

// KindFromSides checks if the three lengths passed in can form a triangle
// If they can, it determines what type of triangle it is. 
func KindFromSides(a, b, c float64) Kind {  
	var k Kind

    switch {
        case a == 0 || b == 0 || c == 0,
        a + b < c || b + c < a || a + c < b:
        k = NaT
        case a == b && b == c:
        k = Equ
        case a == b || a == c || b == c:
        k = Iso
        default:
        k = Sca
    }
    
	return k
}
