package BLC 
import(
	"fmt"
	"math/big"
	"bytes"
	"math"
	"crypto/sha256"
)
const targetBits = 12
var Maxnonce = math.MaxInt64
type ProofOfWork struct {
	Block *Block
	target *big.Int
}


func NewProofOfWork(block *Block) *ProofOfWork{
	target := big.NewInt(1)
	// fmt.Println(target)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Isvalid() bool{
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)
	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0
	var hashInt big.Int
	var hash [32]byte
	for nonce < Maxnonce{
		dataBytes := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)     // \r的作用是 把之前的打印抹掉，即实现只有一行的那种输出
		hashInt.SetBytes(hash[:])
		if proofOfWork.target.Cmp(&hashInt) == 1{
			break
		}
		nonce += 1
	}
	return hash[:], int64(nonce)
}

