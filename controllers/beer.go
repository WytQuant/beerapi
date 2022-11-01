package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"komgrip-api/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type BeerController struct {
	DB *gorm.DB
}

type beersPaging struct {
	Items  []beerResponse
	Paging *pagingResult
}

type beerResponse struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Detail   string `json:"detail"`
	Image    string `json:"image"`
}

type createBeer struct {
	Name     string                `form:"name" binding:"required"`
	Category string                `form:"category" binding:"required"`
	Detail   string                `form:"detail" binding:"required"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
}

type updateBeer struct {
	Name     string                `form:"name"`
	Category string                `form:"category"`
	Detail   string                `form:"detail"`
	Image    *multipart.FileHeader `form:"image"`
}

func (b *BeerController) GetAll(c *gin.Context) {
	var beers []models.Beer

	query := b.DB.Order("id desc")

	beerName := c.Query("name")
	if beerName != "" {
		query = query.Where("name LIKE ?", "%"+beerName+"%")
	}

	pagination := pagination{ctx: c, query: query, records: &beers}
	paging := pagination.paginate()

	serializedBeers := []beerResponse{}
	copier.Copy(&serializedBeers, &beers)

	c.JSON(http.StatusOK, gin.H{
		"beers information": beersPaging{
			Items:  serializedBeers,
			Paging: paging,
		},
	})
}

func (b *BeerController) Create(c *gin.Context) {
	var beerInputData createBeer
	if err := c.ShouldBind(&beerInputData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error message": err.Error(),
		})
		return
	}

	var beer models.Beer
	copier.Copy(&beer, beerInputData)

	if err := b.DB.Create(&beer).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error message": err.Error()})
		return
	}

	if err := b.setBeerImage(c, &beer); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error message": err.Error()})
		return
	}

	serializedBeer := beerResponse{}
	copier.Copy(&serializedBeer, &beer)

	c.JSON(http.StatusCreated, gin.H{
		"beer information": serializedBeer,
	})
}

func (b *BeerController) Update(c *gin.Context) {
	var beerInputData updateBeer
	if err := c.ShouldBind(&beerInputData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error message": err.Error()})
		return
	}

	beer, err := b.findBeerById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	var updateBeerData models.Beer
	copier.Copy(&updateBeerData, &beerInputData)

	if err := b.DB.Model(&beer).Updates(&updateBeerData).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error message": err.Error()})
		return
	}

	b.setBeerImage(c, beer)

	serializedBeer := beerResponse{}
	copier.Copy(&serializedBeer, beer)

	c.JSON(http.StatusOK, gin.H{
		"beer information": serializedBeer,
		"message":          "Update beer information successfully",
	})
}

func (b *BeerController) Delete(c *gin.Context) {
	beerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "Beer id is not a number",
		})
		return
	}

	b.DB.Unscoped().Delete(&models.Beer{}, beerId)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}

func (b *BeerController) findBeerById(c *gin.Context) (*models.Beer, error) {
	var beer models.Beer
	beerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "Beer id is not a number",
		})
		return nil, err
	}

	if err := b.DB.First(&beer, beerId).Error; err != nil {
		return nil, err
	}

	return &beer, nil
}

func (b *BeerController) setBeerImage(c *gin.Context, beer *models.Beer) error {
	file, err := c.FormFile("image")
	if err != nil || file == nil {
		return err
	}

	if beer.Image != "" {
		beer.Image = strings.Replace(beer.Image, os.Getenv("HOST"), "", 1)
		pwd, _ := os.Getwd()
		os.Remove(pwd + beer.Image)
	}

	path := "uploads/beers/" + strconv.Itoa(int(beer.ID))
	os.MkdirAll(path, 0755)
	fileName := path + "/" + file.Filename
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		return err
	}

	beer.Image = os.Getenv("HOST") + "/" + fileName
	b.DB.Save(beer)

	return nil
}
