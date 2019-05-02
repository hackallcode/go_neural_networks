package main

import (
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
    rnn123 := rnn.CreateRNN([][]int{letter1, letter2, letter3}, 15, 100)
    rnn123.PrintWeights()

    res1 := rnn123.Detect(rnn.BWLetterToVector([]string{
        " * ",
        " * ",
        "** ",
        " * ",
        "***",
    }, 3, 5), true, 0)
    rnn.PrintBWLetterByVector(res1, 3, 5)

    res2 := rnn123.Detect(letter2, true, 0)
    rnn.PrintBWLetterByVector(res2, 3, 5)

    res3 := rnn123.Detect(letter3, true, 0)
    rnn.PrintBWLetterByVector(res3, 3, 5)

    res23 := rnn123.Detect(rnn.BWLetterToVector([]string{
        "***",
        "   ",
        "***",
        "   ",
        "***",
    }, 3, 5), true, 0)
    rnn.PrintBWLetterByVector(res23, 3, 5)

    // rnnSTU := rnn.CreateRNN([][]int{letterS, letterT, letterU}, 35, 100)
    // rnnSTU.PrintWeights()
    //
    // resS := rnnSTU.Detect(letterS, false, 0)
    // rnn.PrintBWLetterByVector(resS, 5, 7)
    //
    // resT := rnnSTU.Detect(letterT, false, 0)
    // rnn.PrintBWLetterByVector(resT, 5, 7)
    //
    // resU := rnnSTU.Detect(letterU, false, 0)
    // rnn.PrintBWLetterByVector(resU, 5, 7)
    //
    // resSTU := rnnSTU.Detect(rnn.BWLetterToVector([]string{
    //     "     ",
    //     "*****",
    //     "  *  ",
    //     "  *  ",
    //     "  *  ",
    //     "     ",
    //     "     ",
    // }, 5, 7), false, 0)
    // rnn.PrintBWLetterByVector(resSTU, 5, 7)
}