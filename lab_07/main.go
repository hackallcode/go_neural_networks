package main

import (
    "fmt"
    "strconv"

    "lab_07/rnn"
)

var (
    letter1 = rnn.CreateLetterFromText([]string{
        " @ ",
        "@@ ",
        " @ ",
        " @ ",
        "@@@",
    }, 3, 5)
    letter2 = rnn.CreateLetterFromText([]string{
        "@@@",
        "  @",
        "@@@",
        "@  ",
        "@@@",
    }, 3, 5)
    letter3 = rnn.CreateLetterFromText([]string{
        "@@@",
        "  @",
        " @@",
        "  @",
        "@@@",
    }, 3, 5)
    letterS = rnn.CreateLetterFromText([]string{
        " @@@ ",
        "@   @",
        "@    ",
        " @@@ ",
        "    @",
        "@   @",
        " @@@ ",
    }, 5, 7)
    letterT = rnn.CreateLetterFromText([]string{
        "@@@@@",
        "  @  ",
        "  @  ",
        "  @  ",
        "  @  ",
        "  @  ",
        "  @  ",
    }, 5, 7)
    letterU = rnn.CreateLetterFromText([]string{
        "@   @",
        "@   @",
        "@   @",
        "@   @",
        "@   @",
        "@   @",
        " @@@ ",
    }, 5, 7)
)

func DelectLetter(lettersRNN *rnn.RNN, letter rnn.Letter, sync bool, printMode uint8) {
    res := lettersRNN.DetectByLetter(letter, sync, printMode)
    before := letter.ToText()
    after := res.ToText()

    contentWidth := int(float64(letter.Width() + 10) / 2)
    strWidth := strconv.FormatInt(int64(contentWidth), 10)
    spaceAfter := ""
    for i := 0; i < 10 - contentWidth; i++ {
        spaceAfter += " "
    }

    fmt.Println()
    fmt.Println("+---------------------+")
    fmt.Println("|  Before  |  After   |")
    fmt.Println("+---------------------+")
    for i := uint(0); i < uint(len(before)); i++ {
        fmt.Printf("|%"+strWidth+"v%v|%"+strWidth+"v%v|\n", before[i], spaceAfter, after[i], spaceAfter)
    }
    fmt.Println("+---------------------+")
}

func main() {
    fmt.Println()
    rnn123 := rnn.CreateRNNByLetters([]rnn.Letter{letter1, letter2, letter3}, 15, 100)
    fmt.Println("123 weights:")
    rnn123.PrintWeights()

    DelectLetter(&rnn123, letter1, true, 0)
    DelectLetter(&rnn123, letter2, true, 0)
    DelectLetter(&rnn123, letter3, true, 0)

    DelectLetter(&rnn123, rnn.CreateLetterFromText([]string{
        "@@ ",
        " @ ",
        " @ ",
        " @ ",
        " @ ",
    }, 3, 5), true, 0)
    DelectLetter(&rnn123, rnn.CreateLetterFromText([]string{
        " @@",
        "   ",
        "@@@",
        "@  ",
        "@@ ",
    }, 3, 5), true, 0)
    DelectLetter(&rnn123, rnn.CreateLetterFromText([]string{
        " @@",
        "  @",
        " @@",
        "  @",
        " @@",
    }, 3, 5), true, 0)

    // fmt.Println()
    // rnnSTU := rnn.CreateRNNByLetters([]rnn.Letter{letterS, letterT, letterU}, 35, 100)
    // fmt.Println("STU weights:")
    // rnnSTU.PrintWeights()
    //
    // DelectLetter(&rnnSTU, letterS, false, 0)
    // DelectLetter(&rnnSTU, letterT, false, 0)
    // DelectLetter(&rnnSTU, letterU, false, 0)
    //
    // DelectLetter(&rnnSTU, rnn.CreateLetterFromText([]string{
    //     "     ",
    //     " @@@ ",
    //     "@    ",
    //     " @@@ ",
    //     "    @",
    //     "    @",
    //     " @@@ ",
    // }, 5, 7), false, 0)
    // DelectLetter(&rnnSTU, rnn.CreateLetterFromText([]string{
    //     "     ",
    //     "     ",
    //     " @@@ ",
    //     "  @  ",
    //     "  @  ",
    //     "     ",
    //     "     ",
    // }, 5, 7), false, 0)
    // DelectLetter(&rnnSTU, rnn.CreateLetterFromText([]string{
    //     "     ",
    //     "     ",
    //     "@   @",
    //     "@   @",
    //     "@   @",
    //     " @@@ ",
    //     "     ",
    // }, 5, 7), false, 0)
}
