package service

import (
    "errors"
    "rest-server/internal/model"
)

type BookmarkService struct {
    bookmarks []model.Bookmark
}

func NewBookmarkService() *BookmarkService {
    return &BookmarkService{
        bookmarks: []model.Bookmark{},
    }
}

func (s *BookmarkService) CreateBookmark(bookmark model.Bookmark) error {
    if bookmark.Name == "" || bookmark.Link == "" {
        return errors.New("name and link are required")
    }
    s.bookmarks = append(s.bookmarks, bookmark)
    return nil
}

func (s *BookmarkService) GetBookmarks() []model.Bookmark {
    return s.bookmarks
}

func (s *BookmarkService) UpdateBookmark(id int, updatedBookmark model.Bookmark) error {
    for i, bookmark := range s.bookmarks {
        if bookmark.ID == id {
            s.bookmarks[i] = updatedBookmark
            return nil
        }
    }
    return errors.New("bookmark not found")
}

func (s *BookmarkService) DeleteBookmark(id int) error {
    for i, bookmark := range s.bookmarks {
        if bookmark.ID == id {
            s.bookmarks = append(s.bookmarks[:i], s.bookmarks[i+1:]...)
            return nil
        }
    }
    return errors.New("bookmark not found")
}