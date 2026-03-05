package handler

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/arif14377/koda-b6-backend1/internal/entity"
	"github.com/gin-gonic/gin"
)

var listProducts []entity.Products

// get all products
//
// @Summary      Show an all products available
// @Description  get list all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /products [get]
func GetProducts(c *gin.Context) {
	c.JSON(200, entity.Response{
		Success: true,
		Message: "List products",
		Results: listProducts,
	})
}

// get product details
//
// @Summary      Show details product
// @Description  Get details product with param id
// @Tags         details-product
// @Accept       json
// @Produce      json
// @Param		 id path		int	true		"Product ID"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /products/{id} [get]
func ProductDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := entity.Products{}
	isExist := false

	for _, p := range listProducts {
		if p.Id == id {
			product = p
			isExist = true
			break
		}
	}

	if !isExist {
		c.JSON(404, entity.Response{
			Success: false,
			Message: "Product tidak ditemukan.",
		})
		return
	}

	c.JSON(200, entity.Response{
		Success: true,
		Message: fmt.Sprintf("Product detail id: %d", id),
		Results: product,
	})
}

// add product
//
// @Summary      Add new product
// @Description  Add new product into slice of struct products
// @Tags         add-products
// @Accept       json
// @Produce      json
// @Param		 name 			body		string	true		"Name"
// @Param		 description	body		string	true		"Description"
// @Param		 qty	 		body		int		true		"Quantity"
// @Param		 price	 		body		int		true		"Price"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /products/ [post]
func AddProduct(c *gin.Context) {
	data := entity.Products{}
	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(401, entity.Response{
			Success: false,
			Message: "JSON tidak valid",
		})
		return
	}

	for _, p := range listProducts {
		if strings.Contains(p.Name, data.Name) {
			c.JSON(400, entity.Response{
				Success: false,
				Message: "Nama product sudah ada.",
			})
			return
		}
	}

	if data.Name == "" || data.Description == "" || data.Qty == 0 || data.Price == 0 {
		c.JSON(400, entity.Response{
			Success: false,
			Message: "Data tidak boleh kosong.",
		})
		return
	}

	data.Id = len(listProducts) + 1
	listProducts = append(listProducts, data)
	c.JSON(200, entity.Response{
		Success: true,
		Message: "Produk berhasil ditambahkan.",
	})
}

// delete product
//
// @Summary      delete a product
// @Description  delete data product with id param
// @Tags         delete-products
// @Accept       json
// @Produce      json
// @Param		 id path		int	true "Product ID"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /products/{id} [put]
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	isExist := false

	for i, p := range listProducts {
		if p.Id == id {
			listProducts = slices.Delete(listProducts, i, i+1)
			isExist = true
			break
		}
	}

	if !isExist {
		c.JSON(404, entity.Response{
			Success: false,
			Message: "Id produk tidak ditemukan",
		})
		return
	}

	c.JSON(200, entity.Response{
		Success: true,
		Message: fmt.Sprintf("Data id ke-%d berhasil dihapus.", id),
	})
}

// update product
//
// @Summary      Edit a product
// @Description  edit data product with id param
// @Tags         edit-products
// @Accept       json
// @Produce      json
// @Param 		 id path		int	true "Product ID"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := entity.Products{}
	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(401, entity.Response{
			Success: false,
			Message: "JSON tidak valid.",
		})
		return
	}

	for _, p := range listProducts {
		if strings.Contains(p.Name, data.Name) {
			c.JSON(400, entity.Response{
				Success: false,
				Message: "Nama product sudah ada.",
			})
			return
		}
	}

	if data.Name == "" || data.Description == "" || data.Price == 0 {
		c.JSON(400, entity.Response{
			Success: false,
			Message: "Data tidak boleh kosong.",
		})
		return
	}

	for i, p := range listProducts {
		if p.Id == id {
			listProducts[i] = data
			listProducts[i].Id = id
			break
		}
	}

	c.JSON(200, entity.Response{
		Success: true,
		Message: "Data berhasil diupdate.",
	})
}
