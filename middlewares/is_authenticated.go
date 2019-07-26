package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func IsAuthenticated(store *sessions.FilesystemStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := store.Get(ctx.Request, "auth-session")
		if err != nil {
			fmt.Println("2")
			ctx.Redirect(http.StatusInternalServerError, "/error")
			ctx.Abort()
		}

		if _, ok := session.Values["profile"]; !ok {
			ctx.Redirect(http.StatusUnauthorized, "/index")
			// return errors.New("no profiles session")
		} else {
			ctx.Next()
		}
	}
}
