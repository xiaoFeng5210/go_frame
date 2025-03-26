// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"my_goframe/internal/dao/internal"
)

// internalNewsDao is an internal type for wrapping the internal DAO implementation.
type internalNewsDao = *internal.NewsDao

// newsDao is the data access object for the table news.
// You can define custom methods on it to extend its functionality as needed.
type newsDao struct {
	internalNewsDao
}

var (
	// News is a globally accessible object for table news operations.
	News = newsDao{
		internal.NewNewsDao(),
	}
)

// Add your custom methods and functionality below.
