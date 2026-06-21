package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
    if len(colors) > 2 {
        colors = colors[:2]
    }
    var resistor = map[string]int{
        "black": 0,
        "brown": 1,
        "red": 2,
        "orange": 3,
        "yellow": 4,
        "green": 5,
        "blue": 6,
        "violet": 7,
        "grey": 8,
        "white": 9,
    }
    var resistance string 

    for _, color := range colors {
        resistance += strconv.Itoa(resistor[color])
    }

    res, _ := strconv.Atoi(resistance)

    return res
}
