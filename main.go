package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
	Index     int
	Hash      string
	PrevHash  string
	BPM       int
	Proof     string
	Timestamp string
}

func (block *Block) calculateHash() {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	hex.EncodeToString(hashed)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api", create)
	http.ListenAndServe("8080", nil)
}

func newBlock() {

}

func newTransaction() {

}

func hash() {

}
