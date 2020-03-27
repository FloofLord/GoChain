package blockchain

import "math/big"

// Proof of Work Algorithm

// Take the data from the block

//Create a counter [nonce]  which starts at 0

// create a hash of the data plus the counter

//check the hash to see if its meet a set of requirements

//Requirements:
// The First few bytes must contains 0s

const Difficulty = 12 // Real Block chains dont have this static

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uintn(256-Difficulty))
}
