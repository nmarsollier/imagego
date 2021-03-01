package image

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

var imageMock = Image{
	ID:    "f3db5ffd-77a8-4846-96ee-92a1a5f61d42",
	Image: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAIAAAACCAYAAABytg0kAAABP2lDQ1BJQ0MgUHJvZmlsZQAAKJFjYGASSCwoyGFhYGDIzSspCnJ3UoiIjFJgf8rAxCDGwMugwaCWmFxc4BgQ4ANUwgCjUcG3awyMIPqyLsisHp2sGzdMLjWpxgbGq5nycWKqRwFcKanFyUD6DxAnJhcUlTAwMCYA2crlJQUgdguQLVIEdBSQPQPEToew14DYSRD2AbCakCBnIPsKkC2QnJGYAmQ/AbJ1kpDE05HYUHtBgN3I2DfA1JSAS8kAJakVJSDaOb+gsigzPaNEwREYQqkKnnnJejoKRgZGhgwMoPCGqP58AxyOjGIcCLGCBgYGKw8GBuY8hFhsOwPDhvkMDPy1CDGN8wwMolIMDAecChKLEuEOYPzGUpxmbARhc29nYGCd9v//53CglzUZGP5e////9/b///8uA5p/C6j3GwCnf1v6rtZXfgAAAFZlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAA5KGAAcAAAASAAAARKACAAQAAAABAAAAAqADAAQAAAABAAAAAgAAAABBU0NJSQAAAFNjcmVlbnNob3TQWmjeAAAB0mlUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iWE1QIENvcmUgNi4wLjAiPgogICA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPgogICAgICA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIgogICAgICAgICAgICB4bWxuczpleGlmPSJodHRwOi8vbnMuYWRvYmUuY29tL2V4aWYvMS4wLyI+CiAgICAgICAgIDxleGlmOlBpeGVsWURpbWVuc2lvbj4yPC9leGlmOlBpeGVsWURpbWVuc2lvbj4KICAgICAgICAgPGV4aWY6UGl4ZWxYRGltZW5zaW9uPjI8L2V4aWY6UGl4ZWxYRGltZW5zaW9uPgogICAgICAgICA8ZXhpZjpVc2VyQ29tbWVudD5TY3JlZW5zaG90PC9leGlmOlVzZXJDb21tZW50PgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KFSdiyQAAABlJREFUCB1jLDn94z8jAwMDy9c/DAxAxAAAQ68GKu69WdEAAAAASUVORK5CYII=",
}

func TestFindOriginal(t *testing.T) {
	defer func(restore func(imageID string) (*Image, error)) {
		daoFind = restore
	}(daoFind)

	daoFind = func(imageID string) (*Image, error) {
		return &imageMock, nil
	}

	image, err := Find("123", 0)

	assert.IsEqual(err, nil)
	assert.IsEqual(image.ID, imageMock.ID)
	assert.IsEqual(image.Image, imageMock.Image)
}

func TestFind160(t *testing.T) {
	defer func(restore func(imageID string) (*Image, error)) {
		daoFind = restore
	}(daoFind)

	daoFind = func(imageID string) (*Image, error) {
		return &imageMock, nil
	}

	image, err := Find("123", 160)

	assert.IsEqual(err, nil)
	assert.IsEqual(image.ID, imageMock.ID)
	assert.IsEqual(image.Image, imageMock.Image)
}
