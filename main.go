package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	index               int
	timestamp           time.Time
	data, previous_hash string
	next_block          *Block
}

func hashing(b Block) []byte {
	k := sha256.New()
	k.Write([]byte(strconv.Itoa(b.index)))
	k.Write([]byte(b.timestamp.String()))
	k.Write([]byte(b.data))
	k.Write([]byte(b.previous_hash))
	s := k.Sum(nil)
	return s

}

type Chain struct {
	first_block, current_block *Block
}

func add_to_chain(first_block *Block, d string) {
	b := first_block
	fmt.Println("adding to chain")
	fmt.Println("first: %+v", b)
	for b.next_block != nil {
		b = b.next_block
	}
	fmt.Println("current b: %+v", b)
	new_block := make_block_to_add(b, d)
	fmt.Println("new block to add: %+v", b)
	b.next_block = new_block
	fmt.Println("b with b.next_block: %+v", b)
	b = b.next_block
	fmt.Println("	b = b.next_block : %+v", b)
	fmt.Println("done adding to chain")

}

func make_block_to_add(b *Block, d string) *Block {
	i := b.index + 1
	previous_hash := string(hashing(*b))
	b_2 := Block{i, time.Now(), d, previous_hash, nil}
	return &b_2
}

func print_chain(b *Block) {
	i := 0
	for b.next_block != nil {
		i += 1
		fmt.Println("%+v", b)
		b = b.next_block
	}
	i += 1
	fmt.Println("%+v", b)
	fmt.Println("There were %v blocks", i)
}

func main() {
	b := &Block{0, time.Now(), "Genesis", "FAKE", nil}
	//c := Chain{first_block, first_block}
	print_chain(b)
	add_to_chain(b, "second_block")
	print_chain(b)
	add_to_chain(b, "take my money")
	print_chain(b)
}
