package v1

import (
	"apii_gateway/api/handlers/models"
	pbc "apii_gateway/genproto/comment_service"
	"apii_gateway/pkg/logger"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

//rpc CreateComment(Comment) returns (Comment);
//rpc GetAllCommentsByPostId(GetPostID) returns (AllComments);
//rpc GetAllCommentsByOwnerId(GetOwnerID) returns (AllComments);

// Create Comment
// @Summary create comment
// @Tags Comment
// @Description Create Comment
// @Accept json
// @Produce json
// @Param CommentInfo body models.Comment true "Create Comment"
// @Success 201 {object} models.Comment
// @Failure 400 string Error models.Error
// @Failure 500 string Error models.Error
// @Router /v1/comment/create [post]
func (h *handlerV1) CreateComment(c *gin.Context) {
	var (
		body        models.Comment
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	fmt.Println(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("cannot bind json", logger.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.CommentService().CreateComment(ctx, &pbc.Comment{
		Id:      body.Id,
		Content: body.Content,
		PostId:  body.PostId,
		OwnerId: body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("cannot create comment", logger.Error(err))
	}

	c.JSON(http.StatusOK, response)
}

// Get All Comments By Post Id
// @Summary get all comments by post id
// @Tags Comment
// @Description Get all comments by post id
// @Accept json
// @Produce json
// @Param id path string true "post_id"
// @Success 201 {object} models.Comments
// @Failure 400 string Error models.Error
// @Failure 500 string Error models.Error
// @Router /v1/comment/post/{id} [get]
func (h *handlerV1) GetAllCommentsByPostId(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	postId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetAllCommentsByPostId(ctx, &pbc.GetPostID{
		PostId: postId,
	})

	fmt.Println(response, "commentnttnntntn")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("cannot get comments by post id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get All Comments By Owner Id
// @Summary get all comments by owner id
// @Tags Comment
// @Description Get all comments by owner id
// @Accept json
// @Produce json
// @Param id path string true "owne_id"
// @Success 201 {object} models.Comments
// @Failure 400 string Error models.Error
// @Failure 500 string Error models.Error
// @Router /v1/comment/owner/{id} [get]
func (h *handlerV1) GetAllCommentsByOwnerId(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ownerId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeOut))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetAllCommentsByOwnerId(ctx, &pbc.GetOwnerID{
		OwnerId: ownerId,
	})

	fmt.Println(response)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("cannot get comments by owner id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
