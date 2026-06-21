package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
    if len(colors) > 2 {
        colors = colors[:2]
    }
    var resistor = make(map[string]int)
    var resistance string 

    resistor["black"] = 0
    resistor["brown"] = 1
    resistor["red"] = 2
    resistor["orange"] = 3
    resistor["yellow"] = 4
    resistor["green"] = 5
    resistor["blue"] = 6
    resistor["violet"] = 7
    resistor["grey"] = 8
    resistor["white"] = 9

    for _, color := range colors {
        resistance += strconv.Itoa(resistor[color])
    }

    res, _ := strconv.Atoi(resistance)

    return res
}
