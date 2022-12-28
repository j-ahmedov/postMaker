package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	auth_controller "postMaker/internal/controller/http/v1/auth"
	comment_controller "postMaker/internal/controller/http/v1/comment"
	comment_like_controller "postMaker/internal/controller/http/v1/comment_like"
	post_controller "postMaker/internal/controller/http/v1/post"
	post_file_controller "postMaker/internal/controller/http/v1/post_file"
	post_like_controller "postMaker/internal/controller/http/v1/post_like"
	user_controller "postMaker/internal/controller/http/v1/user"
	"postMaker/internal/middleware"
	"postMaker/internal/pkg/repository/postgres"
	"postMaker/internal/repository/comment"
	"postMaker/internal/repository/comment_like"
	"postMaker/internal/repository/post"
	"postMaker/internal/repository/post_file"
	"postMaker/internal/repository/post_like"
	"postMaker/internal/repository/user"
	auth_service "postMaker/internal/service/auth"
	comment_service "postMaker/internal/service/comment"
	comment_like_service "postMaker/internal/service/comment_like"
	file_service "postMaker/internal/service/file"
	post_service "postMaker/internal/service/post"
	post_file_service "postMaker/internal/service/post_file"
	post_like_service "postMaker/internal/service/post_like"
	user_service "postMaker/internal/service/user"
	auth_usecase "postMaker/internal/usecase/auth"
	comment_usecase "postMaker/internal/usecase/comment"
	post_usecase "postMaker/internal/usecase/post"
	user_usecase "postMaker/internal/usecase/user"
)

func main() {

	// databases
	postgresDB := postgres.NewDB()
	ctx := context.Background()
	postgres.CreateAllTables(postgresDB, ctx)

	//	repositories
	commentRepo := comment.NewRepository(postgresDB)
	commentLikeRepo := comment_like.NewRepository(postgresDB)
	postRepo := post.NewRepository(postgresDB)
	postLikeRepo := post_like.NewRepository(postgresDB)
	postFileRepo := post_file.NewRepository(postgresDB)
	userRepo := user.NewRepository(postgresDB)

	//	services
	commentService := comment_service.NewService(commentRepo)
	commentLikeService := comment_like_service.NewService(commentLikeRepo)
	postService := post_service.NewService(postRepo)
	postLikeService := post_like_service.NewService(postLikeRepo)
	postFileService := post_file_service.NewService(postFileRepo)
	fileService := file_service.NewService()
	userService := user_service.NewService(userRepo)
	authService := auth_service.NewService(userRepo)

	//	use cases
	commentUseCase := comment_usecase.NewUseCase(commentService, commentLikeService, userService)
	postUseCase := post_usecase.NewUseCase(postService, postLikeService, postFileService, fileService, userService)
	userUseCase := user_usecase.NewUseCase(userService, fileService)
	authUseCase := auth_usecase.NewUseCase(authService)

	//	controllers
	commentController := comment_controller.NewController(commentUseCase)
	commentLikeController := comment_like_controller.NewController(commentUseCase)
	postController := post_controller.NewController(postUseCase)
	postLikeController := post_like_controller.NewController(postUseCase)
	postFileController := post_file_controller.NewController(postUseCase)
	userController := user_controller.NewController(userUseCase)
	authController := auth_controller.NewController(authUseCase)

	// gin
	r := gin.Default()

	// No auth requests
	r.POST("/login", authController.GenerateToken)
	r.POST("/user/create", userController.CreateUser)

	// Post File
	r.GET("/post-file/list", postFileController.GetPostFileList).Use(middleware.Auth())
	r.GET("/post-file/detail/:id", postFileController.GetPostFileById).Use(middleware.Auth())
	r.POST("/post-file", postFileController.CreatePostFile).Use(middleware.Auth())
	r.DELETE("/post-file/:id", postFileController.DeletePostFile).Use(middleware.Auth())
	r.PUT("/post-file/update", postFileController.UpdatePostFile).Use(middleware.Auth())

	// Comment
	r.GET("/comment/list", commentController.GetCommentList).Use(middleware.Auth())
	r.GET("/comment/detail/:id", commentController.GetCommentById).Use(middleware.Auth())
	r.GET("/post/comment/:post_id", commentController.GetCommentByPostId).Use(middleware.Auth())
	r.POST("/comment/create", commentController.CreateComment).Use(middleware.Auth())
	r.PUT("/comment/update", commentController.UpdateComment).Use(middleware.Auth())
	r.DELETE("/comment/:id", commentController.DeleteComment).Use(middleware.Auth())

	// Comment Like
	r.GET("/comment-like/list", commentLikeController.GetCommentLikeList).Use(middleware.Auth())
	r.GET("/comment-like/detail/:id", commentLikeController.GetCommentLikeById).Use(middleware.Auth())
	r.POST("/comment-like/create", commentLikeController.CreateCommentLike).Use(middleware.Auth())
	r.PUT("/comment-like/update", commentLikeController.UpdateCommentLike).Use(middleware.Auth())
	r.DELETE("/comment-like/:id", commentLikeController.DeleteCommentLike).Use(middleware.Auth())

	// Post
	r.GET("/post/list", postController.GetPostList).Use(middleware.Auth())
	r.GET("/post/detail/:id", postController.GetPostById).Use(middleware.Auth())
	r.POST("/post/create", postController.CreatePost).Use(middleware.Auth())
	r.PUT("/post/update", postController.UpdatePost).Use(middleware.Auth())
	r.DELETE("/post/:id", postController.DeletePost).Use(middleware.Auth())

	// Post Like
	r.GET("/post-like/list", postLikeController.GetPostLikeList).Use(middleware.Auth())
	r.GET("/post-like/detail/:id", postLikeController.GetPostLikeById).Use(middleware.Auth())
	r.POST("/post-like/create", postLikeController.CreatePostLike).Use(middleware.Auth())
	r.PUT("/post-like/update", postLikeController.UpdatePostLike).Use(middleware.Auth())
	r.DELETE("/post-like/:id", postLikeController.DeletePostLike).Use(middleware.Auth())

	// User
	r.GET("/user/list", userController.GetUserList).Use(middleware.Auth())
	r.GET("/user/detail/:id", userController.GetUserById).Use(middleware.Auth())
	r.PUT("/user/update", userController.UpdateUser).Use(middleware.Auth())
	r.DELETE("/user/:id", userController.DeleteUser).Use(middleware.Auth())

	log.Fatalln(r.Run(":8008"))
}
