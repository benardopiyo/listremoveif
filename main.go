package main

import (
    "github.com/01-edu/z01"
)

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

// ListPushBack adds a new element at the end of the list.
func ListPushBack(l *List, data interface{}) {
    n := &NodeL{Data: data}
    if l.Head == nil {
        l.Head = n
        l.Tail = n
    } else {
        l.Tail.Next = n
        l.Tail = n
    }
}

// PrintList prints the list using z01.
func PrintList(l *List) {
    for n := l.Head; n != nil; n = n.Next {
        switch v := n.Data.(type) {
        case int:
            printInt(v)
        case string:
            for _, r := range v {
                z01.PrintRune(r)
            }
        }
        z01.PrintRune(' ')
        z01.PrintRune('-')
        z01.PrintRune('>')
        z01.PrintRune(' ')
    }
    z01.PrintRune('n')
    z01.PrintRune('i')
    z01.PrintRune('l')
    z01.PrintRune('\n')
}

// printInt prints an integer using z01.
func printInt(n int) {
    if n == 0 {
        z01.PrintRune('0')
        return
    }
    if n < 0 {
        z01.PrintRune('-')
        n = -n
    }
    digits := []rune{}
    for n > 0 {
        digits = append([]rune{rune('0' + n%10)}, digits...)
        n /= 10
    }
    for _, r := range digits {
        z01.PrintRune(r)
    }
}

// ListRemoveIf removes all elements from the list that are equal to data_ref.
func ListRemoveIf(l *List, data_ref interface{}) {
    for l.Head != nil && l.Head.Data == data_ref {
        l.Head = l.Head.Next
    }
    if l.Head == nil {
        l.Tail = nil
        return
    }
    prev := l.Head
    curr := l.Head.Next
    for curr != nil {
        if curr.Data == data_ref {
            prev.Next = curr.Next
            if curr.Next == nil {
                l.Tail = prev
            }
        } else {
            prev = curr
        }
        curr = curr.Next
    }
}

func main() {
    link := &List{}
    link2 := &List{}

    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('n')
    z01.PrintRune('o')
    z01.PrintRune('r')
    z01.PrintRune('m')
    z01.PrintRune('a')
    z01.PrintRune('l')
    z01.PrintRune(' ')
    z01.PrintRune('s')
    z01.PrintRune('t')
    z01.PrintRune('a')
    z01.PrintRune('t')
    z01.PrintRune('e')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('\n')

    ListPushBack(link2, 1)
    PrintList(link2)
    ListRemoveIf(link2, 1)
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('a')
    z01.PrintRune('n')
    z01.PrintRune('s')
    z01.PrintRune('w')
    z01.PrintRune('e')
    z01.PrintRune('r')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('\n')
    PrintList(link2)
    z01.PrintRune('\n')

    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('n')
    z01.PrintRune('o')
    z01.PrintRune('r')
    z01.PrintRune('m')
    z01.PrintRune('a')
    z01.PrintRune('l')
    z01.PrintRune(' ')
    z01.PrintRune('s')
    z01.PrintRune('t')
    z01.PrintRune('a')
    z01.PrintRune('t')
    z01.PrintRune('e')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('\n')

    ListPushBack(link, 1)
    ListPushBack(link, "Hello")
    ListPushBack(link, 1)
    ListPushBack(link, "There")
    ListPushBack(link, 1)
    ListPushBack(link, 1)
    ListPushBack(link, "How")
    ListPushBack(link, 1)
    ListPushBack(link, "are")
    ListPushBack(link, "you")
    ListPushBack(link, 1)
    PrintList(link)

    ListRemoveIf(link, 1)
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('a')
    z01.PrintRune('n')
    z01.PrintRune('s')
    z01.PrintRune('w')
    z01.PrintRune('e')
    z01.PrintRune('r')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('—')
    z01.PrintRune('\n')
    PrintList(link)
}
