package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "bookmarks-app/rest-server/internal/service"
)

type BookmarkController struct {
    service *service.BookmarkService
}

func NewBookmarkController(s *service.BookmarkService) *BookmarkController {
    return &BookmarkController{service: s}
}

func (bc *BookmarkController) CreateBookmark(c *gin.Context) {
    var bookmark service.Bookmark
    if err := c.ShouldBindJSON(&bookmark); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := bc.service.CreateBookmark(&bookmark); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, bookmark)
}

func (bc *BookmarkController) GetBookmarks(c *gin.Context) {
    bookmarks, err := bc.service.GetBookmarks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, bookmarks)
}

func (bc *BookmarkController) UpdateBookmark(c *gin.Context) {
    id := c.Param("id")
    var bookmark service.Bookmark
    if err := c.ShouldBindJSON(&bookmark); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    bookmark.ID = id
    if err := bc.service.UpdateBookmark(&bookmark); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, bookmark)
}

func (bc *BookmarkController) DeleteBookmark(c *gin.Context) {
    id := c.Param("id")
    if err := bc.service.DeleteBookmark(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}