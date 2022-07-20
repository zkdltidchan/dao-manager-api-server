package models

import "fmt"

const (
	DefaultLimit int = 25
	// DefaultOffset int = 0
)

type Core struct {
	Size       int `json:"size"`
	PageIndex  int `json:"page_index"`
	PageCounts int `json:"page_counts"`
	Total      int `json:"total"`
}

type Pagination struct {
	Core
	Data interface{} `json:"data"`
}

func GetPages(total int, size int) int {
	fmt.Printf("%v,,%v", total, size)
	if size == 0 {
		size = DefaultLimit
	}

	pages := total / size

	if total%size > 0 {
		pages++
	}
	return pages
}

// db querry offet, page index
func GetOffSet(pageIndex int, limit int, pages int) int {
	if pageIndex > pages {
		pageIndex = pages
	}
	offSet := (pageIndex - 1) * limit
	return offSet
}

// db querry limit, page size
func GetLimit(size int) int {
	if size == 0 {
		size = DefaultLimit
	}
	return size
}

func GetCurrentPage(offSet int) int {
	if offSet < 1 {
		return 1
	} else {
		return offSet
	}
}
