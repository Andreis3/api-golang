package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NEW_PROduct(t *testing.T) {
	product, err := NewProduct("product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "product 1", product.Name)
	assert.Equal(t, 100, product.Price)
}

func Test_PRODUCT_WHEN_NAME_IS_REQUIRED(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func Test_PRODUCT_WHEN_PRICE_IS_REQUIRED(t *testing.T) {
	product, err := NewProduct("product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func Test_PRODUCT_WHEN_PRICE_IS_INVALID(t *testing.T) {
	product, err := NewProduct("product 1", -1)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrorInvalidPrice, err)
}

func Test_PRODUCT_VALIDATE(t *testing.T) {
	product, err := NewProduct("product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	err = product.Validate()
	assert.Nil(t, err)
}
