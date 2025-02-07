package apis

import (
	"blogo/config"
	"blogo/ent"
	"context"
	"fmt"

	"github.com/kataras/iris/v12"
)

type DeletePostRequest struct {
	ID    uint32 `json:"id"`
	Token string `json:"token"`
}

func DeletePost(config config.Config, db *ent.Client) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		var request DeletePostRequest

		err := ctx.ReadJSON(&request)
		if err != nil {
			fmt.Println("DeletePost:unmarshal json error")
			ctx.Err()
			return
		}

		if config.Password != request.Token {
			fmt.Println("DeletePost:token is error")
			ctx.Err()
			return
		}
		err = db.Post.DeleteOneID(request.ID).Exec(context.Background())

		if err != nil {
			fmt.Println("CreatePost:create post error")
		}

		ctx.JSON(Result{Ok: err == nil})
	}
}
