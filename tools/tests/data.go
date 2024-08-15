package tests

import (
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock Data
func TestUser() *security.User {
	return &security.User{
		ID:          primitive.NewObjectID().String(),
		Login:       "Login",
		Name:        "Name",
		Permissions: []string{"user"},
	}
}

func TestImage() *image.Image {
	return image.New("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=")
}

func TestInvalidImage() *image.Image {
	return image.New("___")
}

func TestResizeImage() *image.Image {
	return image.New("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAKUAAAClCAIAAACySaqNAAABhGlDQ1BJQ0MgcHJvZmlsZQAAKJF9kT1Iw0AYht+mSou0dLCDiEOE6mQXFXGsVShChVArtOpgcukfNGlJUlwcBdeCgz+LVQcXZ10dXAVB8AfE2cFJ0UVK/C4ptIjx4O4e3vvel7vvAKFVZZrZlwA03TIyqaSYy6+KgVcEEUGY1lGZmfU5SUrDc3zdw8f3uzjP8q77c4TVgskAn0icYHXDIt4gntm06pz3iaOsLKvE58QTBl2Q+JHristvnEsOCzwzamQz88RRYrHUw0oPs7KhEU8Tx1RNp3wh57LKeYuzVm2wzj35C0MFfWWZ6zRHkMIiliBBhIIGKqjCQpx2nRQTGTpPeviHHb9ELoVcFTByLKAGDbLjB/+D3701i1OTblIoCfS/2PbHGBDYBdpN2/4+tu32CeB/Bq70rr/WAmY/SW92tdgRENkGLq67mrIHXO4AQ0912ZAdyU9TKBaB9zP6pjwweAsMrLl965zj9AHIUq/SN8DBITBeoux1j3cHe/v2b02nfz9eInKemubWvQAAAAlwSFlzAAAuIwAALiMBeKU/dgAAAAd0SU1FB+gIDw0IOcHU+YsAAAAZdEVYdENvbW1lbnQAQ3JlYXRlZCB3aXRoIEdJTVBXgQ4XAAAAZklEQVR42u3BMQEAAADCoPVPbQdvoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB4DD+/AAGTr/KYAAAAAElFTkSuQmCC")
}
