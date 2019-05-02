package main

import (
    "fmt"

    "labs/rnn"
)

var (
    letter1 = rnn.BWLetterToVector([]string{
        " * ",
        "** ",
        " * ",
        " * ",
        "***",
    }, 3, 5)
    letter2 = rnn.BWLetterToVector([]string{
        "***",
        "  *",
        "***",
        "*  ",
        "***",
    }, 3, 5)
    letter3 = rnn.BWLetterToVector([]string{
        "***",
        "  *",
        " **",
        "  *",
        "***",
    }, 3, 5)
    letterS = rnn.BWLetterToVector([]string{
        " *** ",
        "*   *",
        "*    ",
        " *** ",
        "    *",
        "*   *",
        " *** ",
    }, 5, 7)
    letterT = rnn.BWLetterToVector([]string{
        "*****",
        "  *  ",
        "  *  ",
        "  *  ",
        "  *  ",
        "  *  ",
        "  *  ",
    }, 5, 7)
    letterU = rnn.BWLetterToVector([]string{
        "*   *",
        "*   *",
        "*   *",
        "*   *",
        "*   *",
        "*   *",
        " *** ",
    }, 5, 7)
)

func main() {
    fmt.Println()
    rnn123 := rnn.CreateRNN([][]int{letter1, letter2, letter3}, 15, 100)
    fmt.Println("123 weights:")
    rnn123.PrintWeights()

    fmt.Println()
    res1 := rnn123.Detect(letter1, true, 1)
    fmt.Println("Letter 1:")
    rnn.PrintBWLetterByVector(res1, 3, 5)

    fmt.Println()
    res2 := rnn123.Detect(letter2, true, 1)
    fmt.Println("Letter 2:")
    rnn.PrintBWLetterByVector(res2, 3, 5)

    fmt.Println()
    res3 := rnn123.Detect(letter3, true, 1)
    fmt.Println("Letter 3:")
    rnn.PrintBWLetterByVector(res3, 3, 5)

    fmt.Println()
    resBad1 := rnn123.Detect(rnn.BWLetterToVector([]string{
        "** ",
        " * ",
        " * ",
        " * ",
        " * ",
    }, 3, 5), true, 1)
    fmt.Println("Bad letter 1:")
    rnn.PrintBWLetterByVector(resBad1, 3, 5)

    fmt.Println()
    resBad2 := rnn123.Detect(rnn.BWLetterToVector([]string{
        " **",
        "   ",
        "***",
        "*  ",
        "** ",
    }, 3, 5), true, 1)
    fmt.Println("Bad letter 2:")
    rnn.PrintBWLetterByVector(resBad2, 3, 5)

    fmt.Println()
    resBad3 := rnn123.Detect(rnn.BWLetterToVector([]string{
        " * ",
        "  *",
        " **",
        "  *",
        " **",
    }, 3, 5), true, 1)
    fmt.Println("Bad letter 3:")
    rnn.PrintBWLetterByVector(resBad3, 3, 5)

    // fmt.Println()
    // rnnSTU := rnn.CreateRNN([][]int{letterS, letterT, letterU}, 35, 100)
    // fmt.Println("STU weights:")
    // rnnSTU.PrintWeights()
    //
    // fmt.Println()
    // resS := rnnSTU.Detect(letterS, false, 0)
    // fmt.Println("Letter S:")
    // rnn.PrintBWLetterByVector(resS, 5, 7)
    //
    // fmt.Println()
    // resT := rnnSTU.Detect(letterT, false, 0)
    // fmt.Println("Letter T:")
    // rnn.PrintBWLetterByVector(resT, 5, 7)
    //
    // fmt.Println()
    // resU := rnnSTU.Detect(letterU, false, 0)
    // fmt.Println("Letter U:")
    // rnn.PrintBWLetterByVector(resU, 5, 7)
    //
    // fmt.Println()
    // resBadS := rnnSTU.Detect(rnn.BWLetterToVector([]string{
    //     "     ",
    //     " *** ",
    //     "*    ",
    //     " *** ",
    //     "    *",
    //     "    *",
    //     " *** ",
    // }, 5, 7), false, 0)
    // fmt.Println("Bad letter S:")
    // rnn.PrintBWLetterByVector(resBadS, 5, 7)
    //
    // fmt.Println()
    // resBadT := rnnSTU.Detect(rnn.BWLetterToVector([]string{
    //     "     ",
    //     "     ",
    //     " *** ",
    //     "  *  ",
    //     "  *  ",
    //     "     ",
    //     "     ",
    // }, 5, 7), false, 0)
    // fmt.Println("Bad letter T:")
    // rnn.PrintBWLetterByVector(resBadT, 5, 7)
    //
    // fmt.Println()
    // resBadU := rnnSTU.Detect(rnn.BWLetterToVector([]string{
    //     "     ",
    //     "     ",
    //     "*   *",
    //     "*   *",
    //     "*   *",
    //     " *** ",
    //     "     ",
    // }, 5, 7), false, 0)
    // fmt.Println("Bad letter U:")
    // rnn.PrintBWLetterByVector(resBadU, 5, 7)
}
