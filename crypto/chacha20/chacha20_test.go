package chacha20

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChacha20_1(t *testing.T) {
	keyb := []byte("adsfdsaenc,.dsafadafd14454646dasfasdf90789689dasfdas113232")
	sskey := [32]byte(keyb)

	plan := []byte("hello")
	edata, err := Encrypt(sskey, plan)
	require.NoError(t, err)

	ddata, err := Decrypt(sskey, edata)
	require.NoError(t, err)

	fmt.Println("Decrypted:  ", string(ddata))
}

// GenerateRandomBytes 生成指定长度的随机字节（用于密钥/Nonce）
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("生成随机字节失败: %w", err)
	}
	return b, nil
}
func TestChaCha20Poly1305(t *testing.T) {
	// 1. 原始明文
	plaintext := []byte("Hello, ChaCha20! 这是测试数据")
	fmt.Printf("原始明文: %s\n", plaintext)

	// 2. 生成随机密钥（32字节）和Nonce（12字节）
	key, err := GenerateRandomBytes(32)
	require.NoError(t, err)

	nonce, err := GenerateRandomBytes(12)
	require.NoError(t, err)

	adata, err := GenerateRandomBytes(64)
	require.NoError(t, err)

	fmt.Printf("随机密钥 (hex): %s\n", hex.EncodeToString(key))
	fmt.Printf("随机密钥 (base64): %s\n", base64.StdEncoding.EncodeToString(key))
	fmt.Printf("随机Nonce (hex): %s\n", hex.EncodeToString(nonce))
	fmt.Printf("随机Nonce (base64): %s\n", base64.StdEncoding.EncodeToString(nonce))
	fmt.Printf("随机AData (hex): %s\n", hex.EncodeToString(adata))
	fmt.Printf("随机AData (base64): %s\n", base64.StdEncoding.EncodeToString(adata))

	// 3. 加密
	ciphertext, err := ChaCha20Poly1305Encrypt(key, nonce, plaintext, adata)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	fmt.Printf("加密后密文 (hex): %s\n", hex.EncodeToString(ciphertext))
	fmt.Printf("加密后密文 (base64): %s\n", base64.StdEncoding.EncodeToString(ciphertext))

	// 4. 解密
	decryptedText, err := ChaCha20Poly1305Decrypt(key, nonce, ciphertext, adata)
	if err != nil {
		log.Fatalf("解密失败: %v", err)
	}
	fmt.Printf("解密后明文: %s\n", decryptedText)

	// 5. 验证解密结果
	if string(decryptedText) != string(plaintext) {
		log.Fatal("解密结果与原始明文不一致！")
	}
	fmt.Println("✅ 加解密验证成功")
}

func TestXChaCha20Poly1305(t *testing.T) {
	// 1. 原始明文
	plaintext := []byte("TestXChaCha20Poly1305")
	fmt.Printf("原始明文: %s\n", plaintext)

	// 2. 生成随机密钥（32字节）和Nonce（12字节）
	key, err := GenerateRandomBytes(32)
	require.NoError(t, err)

	nonce, err := GenerateRandomBytes(24)
	require.NoError(t, err)

	adata, err := GenerateRandomBytes(64)
	require.NoError(t, err)

	fmt.Printf("随机密钥 (hex): %s\n", hex.EncodeToString(key))
	fmt.Printf("随机密钥 (base64): %s\n", base64.StdEncoding.EncodeToString(key))
	fmt.Printf("随机Nonce (hex): %s\n", hex.EncodeToString(nonce))
	fmt.Printf("随机Nonce (base64): %s\n", base64.StdEncoding.EncodeToString(nonce))
	fmt.Printf("随机AData (hex): %s\n", hex.EncodeToString(adata))
	fmt.Printf("随机AData (base64): %s\n", base64.StdEncoding.EncodeToString(adata))

	// 3. 加密
	ciphertext, err := XChaCha20Poly1305Encrypt(key, nonce, plaintext, adata)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	fmt.Printf("加密后密文 (hex): %s\n", hex.EncodeToString(ciphertext))
	fmt.Printf("加密后密文 (base64): %s\n", base64.StdEncoding.EncodeToString(ciphertext))

	// 4. 解密
	decryptedText, err := XChaCha20Poly1305Decrypt(key, nonce, ciphertext, adata)
	if err != nil {
		log.Fatalf("解密失败: %v", err)
	}
	fmt.Printf("解密后明文: %s\n", decryptedText)

	// 5. 验证解密结果
	if string(decryptedText) != string(plaintext) {
		log.Fatal("解密结果与原始明文不一致！")
	}
	fmt.Println("✅ 加解密验证成功")
}
