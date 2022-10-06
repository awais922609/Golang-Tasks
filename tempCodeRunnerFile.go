	fmt.Printf("\nChange Block Hash is == %x", b1.CalculateHash(b1.transaction+string(b1.nonce)+b1.previousHash))
