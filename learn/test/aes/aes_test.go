package aes_test

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	// 示例密文、密钥和IV，请替换为实际值
	base64Bytes, _ := base64.StdEncoding.DecodeString("JQntYQIQTWUb6cvqgKLOs8SbUcJMB+uWM9ZJOcbuXj4tw1J1N9QBb3rzPSFHpmiRQvMHhoeTfHxXMo1ZMMvX+gLRsM8FXB4IbtZ8cntcelPLtAcF+SAdQESNlAIUbJYAI3/XIWNAsy27/02ir2T92OEv6XhD4MBLccx14J1Q8mtKyxJqr95SPn4KOekXOhs2uDRetV0b1jqqVoEswNyFrcXmEczqmEKmnHWVgjHtnj7IPNaOIFoMx/ALGYa7nGsJV1wHTOcN31b8d9gKBJL5JgDEAbIp7tCuUT4nx0NtTAHSYlwA1UvSK4mSyDhm1JbOQdts18bK69EvVMbWZlu/NGg/hphRlcUDXpPo5JZcMyxclOmnVAgX4NnmpQRNCfhUcAywW81N5u8SFMfODnolON/uFR5pOxHbVEzsqd/Rfqg/3dLQQHi2AIzQV7DrUur8fyv+watEUatw1wVgqAd1Jz+gZfHqIML/d2am9v7JuAff7vBZbm7twcRhBUYQqhy9XkNGCxUwuj1xv9y/8BH4uhy+urAQPMuUrH+K15H1evGLDD8Re4j2MH4G3vOB5o+XvEzyjLCCJdKpfCN1HEFKq4GMUyoofj/RJpFdk+YudUXR9ubkNOTuIHjCkitDl7AkKwHeG/MptyGkVqsx4bxwc7UIKvnKbWZwyey/hVWkzoSS4oPY884ceKN+nmDzJ8Omf4igP2GuwOt+Thc3/sTW1DII7esO84Ahb0Bt4lXLbC8S2DImtHUVnI+p7Nnr3OD65+ThOoC70cnaO9JvtfQfG9h+AjCRH4ADRPO8STEzrgIwGFjva8l0N054cJdJMWwNCaotjISru1T4F2J7JgDc/MQbCJhWR7/9vALxYkWjTeI32uVHefBlc8EQreD1E2U9Eu4AqwYSebkE4Z5M54Yng6NVIlLTz+DczwaGQuXSCUB8bNT47C5UIO+aDQXS5qxxKxiGUDvMxw0mCqFHeNcUx3W+6dVIRG2eE0QuF9dKzYh/X3KAVgu4v4PPyc/BHRD2aTxfeMrNr5UPSC/IstFSH+NG50GDgHZSHlC8ffbVy4nbllxU6j+0XE2rteNe5/w0uONWnepJgcD0oe4OLkuBEzHAnEGPJ+4kG6GdQcaLeILBfE45H70XyrVn5MCGWGNcc3KD3W1dlfpgs08hljtFxvBJstd1Fl8PhyfOEVlPKrR8W0PyI7Pl8iQ70Dr1t9/2KYFR2N5PlBZiSXR3jp/ghJEtwkC4RIBT6nvgHZ/55cSGyGf/UkSS1LjtJC95H3sI2mN4fRkdapxMqLzsdq+Q1MnevOlPWsOikJxhqbRwNH8hRlDoTmIlgvWZFj/00Y/LVmDTxU/eAQaowdsjErHR/pAfH37+yl81sM327Ruvfv3wbQrrhZ1JPWitXROBlgyaYPhwQTBzZC7C1wx+dj91IkmtR3UGcoAVQ0ywrCWxXDaAgxq8tcpjXP03RNJRPys1kmdd1z/7+NRW8OWD19hQRpAnZpv5FNm+164JZXuwwpo=")
	iv := base64Bytes[:16]
	ciphertext := base64Bytes[16:]
	key := make([]byte, 32)
	copy(key, []byte("64f968f04843877e4525fb24"))

	plaintext, err := decryptAES(ciphertext, key, iv)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Decrypted Text:", string(plaintext))
}

func decryptAES(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}

	// CBC mode decryption
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	// PKCS7 unpadding
	return unpadPKCS7(ciphertext), nil
}

func unpadPKCS7(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
