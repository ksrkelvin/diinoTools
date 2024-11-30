package tools

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InterfaceFloat64 - Converte interface{} para Float64
func (p Tools) InterfaceFloat64(entrada interface{}) (saida float64) {

	saida = float64(0)
	switch entrada := entrada.(type) {
	case float64:
		saida = entrada
	case float32:
		fmt.Println(entrada)
		saida = decimal.NewFromFloat32(entrada).InexactFloat64()
	case int:
		saida = float64(entrada)
	case int32:
		saida = float64(entrada)
	case int64:
		saida = float64(entrada)
	case string:
		saida, err := strconv.ParseFloat(entrada, 64)
		if err != nil {
			saida = 0
		}
		return saida
	default:
		saida = float64(0)
	}

	return saida
}

// InterfaceString - Converte interface{} para string
func (p Tools) InterfaceString(entrada interface{}) (saida string) {
	saida = ""
	switch entrada := entrada.(type) {
	case int32:
		saida = strconv.Itoa(int(entrada))
	case int64:
		saida = strconv.Itoa(int(entrada))
	case int:
		saida = strconv.Itoa(entrada)
	case float32:
		saida = strconv.FormatFloat(float64(entrada), 'f', -1, 32)
	case float64:
		saida = strconv.FormatFloat(entrada, 'f', -1, 64)
	case string:
		saida = entrada
	case time.Time:
		saida = entrada.Format("2006-01-02 15:04:05")
	default:
		saida = ""
	}
	return saida
}

// InterfaceTime - Converte interface{} para time.Time
func (p Tools) InterfaceTime(entrada interface{}) (saida time.Time) {
	saida, _ = time.Parse("2006-01-02 15:04:05", "1900-01-01 00:00:00")
	switch entrada := entrada.(type) {
	case time.Time:
		saida = entrada
	case primitive.DateTime:
		saida = entrada.Time()
	default:

	}
	return saida
}

// InterfaceInt - Converte interface{} para int
func (p Tools) InterfaceInt(entrada interface{}) (saida int, err error) {
	switch entrada := entrada.(type) {
	case int32:
		saida = int(entrada)
	case int64:
		saida = int(entrada)
	case int:
		saida = entrada
	case float32:
		saida = int(entrada)
	case float64:
		saida = int(entrada)
	case string:
		saida, err = strconv.Atoi(entrada)
	default:
		saida = 0
	}
	return saida, err
}

// InterfaceToIntArray - Converte interface{} para []int
func (p Tools) InterfaceToIntArray(input interface{}) ([]int, error) {
	slice, ok := input.([]interface{})
	if !ok {
		return nil, fmt.Errorf("input is not a []interface{}")
	}

	intArray := make([]int, len(slice))
	for i, v := range slice {
		if num, ok := v.(int); ok {
			intArray[i] = num
		} else {
			return nil, fmt.Errorf("element at index %d is not an int", i)
		}
	}

	return intArray, nil
}
