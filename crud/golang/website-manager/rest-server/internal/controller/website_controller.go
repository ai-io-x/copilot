package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "rest-server/internal/service"
    "rest-server/internal/models"
)

type WebsiteController struct {
    service service.WebsiteService
}

func NewWebsiteController(service service.WebsiteService) *WebsiteController {
    return &WebsiteController{service: service}
}

func (c *WebsiteController) CreateWebsite(ctx *gin.Context) {
    var website models.Website
    if err := ctx.ShouldBindJSON(&website); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdWebsite, err := c.service.CreateWebsite(website)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, createdWebsite)
}

func (c *WebsiteController) GetWebsites(ctx *gin.Context) {
    websites, err := c.service.GetWebsites()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, websites)
}

func (c *WebsiteController) UpdateWebsite(ctx *gin.Context) {
    var website models.Website
    if err := ctx.ShouldBindJSON(&website); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id := ctx.Param("id")
    updatedWebsite, err := c.service.UpdateWebsite(id, website)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, updatedWebsite)
}

func (c *WebsiteController) DeleteWebsite(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.DeleteWebsite(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.Status(http.StatusNoContent)
}