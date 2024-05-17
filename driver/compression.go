package driver

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/gobkc/fit/conf"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Compression struct {
	d Driver
}

func NewCompression() *Compression {
	return &Compression{d: d}
}

func (c *Compression) AddFiles(password string, files ...string) []byte {
	cachePath := conf.GetCachePath()
	if len(files) == 0 {
		files = c.getAllFiles(cachePath)
	}
	var allBytes bytes.Buffer
	separator := '\u200D'
	timeSeparator := '\u200B'
	for _, filePath := range files {
		fileByte, err := os.ReadFile(filePath)
		if err != nil {
			slog.Default().Warn(`failed to read file`, slog.String(`file path`, filePath), slog.String(`error`, err.Error()))
		} else {
			rawPath := strings.Replace(filePath, cachePath, ``, 1)
			allBytes.WriteString(string(separator))
			allBytes.WriteString(rawPath)
			fileInfo, err := os.Stat(filePath)
			if err == nil {
				allBytes.WriteString(string(timeSeparator))
				allBytes.WriteString(fileInfo.ModTime().Format(time.DateTime))
			}
			allBytes.WriteString("\n")
			allBytes.Write(fileByte)
		}
	}
	return allBytes.Bytes()
	cmzBytes, err := CompressData(allBytes.Bytes())
	if err != nil {
		slog.Default().Warn(`failed to compress file`, slog.String(`error`, err.Error()))
		return nil
	}

	if password != `` {
		key := GetKey(password)
		passed, err := Encrypt(key, cmzBytes)
		if err != nil {
			slog.Default().Warn(`Unable to encrypt file`, slog.String(`error`, err.Error()))
			return nil
		}
		return passed
	}
	return cmzBytes
}

func (c *Compression) DeCompress(password string, file []byte) (files []CompressionFile) {
	var raw = file
	if password != `` {
		key := GetKey(password)
		decryptBytes, err := Decrypt(key, file)
		if err != nil {
			slog.Default().Warn(`Unable to decrypt file`, slog.String(`error`, err.Error()))
			return nil
		}
		raw, err = DecompressData(decryptBytes)
		if err != nil {
			slog.Default().Warn(`Unable to decompress data`, slog.String(`error`, err.Error()))
		}
	}
	separator := '\u200D'
	timeSeparator := '\u200B'
	splitList := bytes.Split(raw, []byte(string(separator)))
	splitListLen := len(splitList)
	cachePath := conf.GetCachePath()
	for i, fileRaw := range splitList {
		if len(fileRaw) == 0 {
			continue
		}
		fileSplit := SplitFirst(fileRaw, []byte("\n"))
		if len(fileSplit) != 2 {
			continue
		}
		fileHeader, fileContent := fileSplit[0], fileSplit[1]
		if fileContentLen := len(fileContent); i < splitListLen-1 && fileContentLen > 0 {
			fileContent = fileContent[:fileContentLen-1]
		}
		fileMetaSplit := SplitFirst(fileHeader, []byte(string(timeSeparator)))
		if len(fileMetaSplit) != 2 {
			continue
		}
		fullPath := cachePath + string(fileMetaSplit[0])
		updatedTime, _ := time.Parse(time.DateTime, string(fileMetaSplit[1]))
		cate := filepath.Dir(fullPath)
		filename := filepath.Base(fullPath)
		files = append(files, CompressionFile{
			Cate:        cate,
			Filename:    filename,
			Content:     string(fileContent),
			UpdatedTime: updatedTime,
		})
	}
	return files
}

func SplitFirst(s, sep []byte) [][]byte {
	index := bytes.Index(s, sep)
	if index == -1 {
		return [][]byte{s, nil}
	}
	return [][]byte{s[:index], s[index+len(sep):]}
}

func (c *Compression) getAllFiles(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		slog.Default().Warn(`failed to walk dir`, slog.String(`error`, err.Error()))
	}
	return files
}

func CompressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecompressData decompresses the given gzip compressed data.
func DecompressData(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

func GetKey(password string) []byte {
	if passLen := 32 - len(password); passLen > 0 {
		password += strings.Repeat(`0`, passLen)
	} else if passLen < 0 {
		password = password[0:32]
	}
	return []byte(password)
}

// Encrypt encrypts data using AES with the given key.
func Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Pad the plaintext to be multiple of block size
	plaintext = PKCS7Padding(plaintext, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	// Use CBC mode with zero IV (not recommended for real-world use)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// Decrypt decrypts data using AES with the given key.
func Decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode decrypts in-place, so need to create a copy
	plaintext := make([]byte, len(ciphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove PKCS7 padding
	plaintext = PKCS7UnPadding(plaintext)

	return plaintext, nil
}

// PKCS7Padding adds padding to the data according to PKCS#7.
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding adds padding to the data according to PKCS#7.
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
