package rnn

import (
    "fmt"
)

const (
    BlackChar = '@'
    WhiteChar = ' '
)

type Letter struct {
    data   []int
    width  uint
    height uint
}

func CreateLetterFromText(letter []string, w, h uint) (res Letter) {
    res.width = w
    res.height = h

    res.data = make([]int, w*h)
    for j := uint(0); j < h; j++ {
        for i := uint(0); i < w; i++ {
            if letter[j][i] == BlackChar {
                res.data[h*i+j] = 1
            } else {
                res.data[h*i+j] = -1
            }
        }
    }

    return res
}

func (l *Letter) ToText() []string {
    text := make([]string, l.height)
    for j := uint(0); j < l.height; j++ {
        for i := uint(0); i < l.width; i++ {
            if l.data[l.height*i+j] > 0 {
                text[j] += string(BlackChar)
            } else {
                text[j] += string(WhiteChar)
            }
        }
    }
    return text
}

func (l *Letter) Print() {
    for j := uint(0); j < l.height; j++ {
        for i := uint(0); i < l.width; i++ {
            if l.data[l.height*i+j] > 0 {
                fmt.Printf("%c", BlackChar)
            } else {
                fmt.Printf("%c", WhiteChar)
            }
        }
        fmt.Println()
    }
}

func (l *Letter) Height() uint {
    return l.height
}

func (l *Letter) Width() uint {
    return l.width
}
