package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	//println("-----------BLOCK HEADER JSON-----------------")
	//json, _ := block.Header().MarshalJSON()
	//println(string(json))
	//println("------------BLOCK HEADER---------------------")
	//payload, err := rlp.EncodeToBytes(block.Header())
	//println(hexutil.Encode(payload)[2:])
	//println("----------EXTRA DATA SHOULD BE---------------")
	//println(hexutil.Encode(block.Header().Extra[:len(block.Header().Extra)-65])[2:])
	//println("----------SIGNING DATA-----------------------")
	//signingData := parlia.ParliaRLP(block.Header(), big.NewInt(56))
	//println(hexutil.Encode(signingData)[2:])
}
