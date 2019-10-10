// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"morg/restapi/operations"
	"morg/restapi/operations/articles"
	"morg/restapi/operations/auth"
	"morg/restapi/operations/comments"
	"morg/restapi/operations/feeds"
	"morg/restapi/operations/me"
	"morg/restapi/operations/profile"
)

//go:generate swagger generate server --target ../../morg --name MinterOrgHub --spec ../swagger.yaml

func configureFlags(api *operations.MinterOrgHubAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MinterOrgHubAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	api.BearerAuthAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (bearerAuth) Authorization from header param [Authorization] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.CommentsDeleteCommentsCommentIDHandler == nil {
		api.CommentsDeleteCommentsCommentIDHandler = comments.DeleteCommentsCommentIDHandlerFunc(func(params comments.DeleteCommentsCommentIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation comments.DeleteCommentsCommentID has not yet been implemented")
		})
	}
	if api.FeedsDeleteFeedsFeedIDHandler == nil {
		api.FeedsDeleteFeedsFeedIDHandler = feeds.DeleteFeedsFeedIDHandlerFunc(func(params feeds.DeleteFeedsFeedIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.DeleteFeedsFeedID has not yet been implemented")
		})
	}
	if api.FeedsDeleteFeedsFeedIDFollowHandler == nil {
		api.FeedsDeleteFeedsFeedIDFollowHandler = feeds.DeleteFeedsFeedIDFollowHandlerFunc(func(params feeds.DeleteFeedsFeedIDFollowParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.DeleteFeedsFeedIDFollow has not yet been implemented")
		})
	}
	if api.FeedsDeleteFeedsFeedIDJoinHandler == nil {
		api.FeedsDeleteFeedsFeedIDJoinHandler = feeds.DeleteFeedsFeedIDJoinHandlerFunc(func(params feeds.DeleteFeedsFeedIDJoinParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.DeleteFeedsFeedIDJoin has not yet been implemented")
		})
	}
	if api.FeedsDeleteFeedsFeedIDJoinUserIDHandler == nil {
		api.FeedsDeleteFeedsFeedIDJoinUserIDHandler = feeds.DeleteFeedsFeedIDJoinUserIDHandlerFunc(func(params feeds.DeleteFeedsFeedIDJoinUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.DeleteFeedsFeedIDJoinUserID has not yet been implemented")
		})
	}
	if api.ProfileDeleteProfilesUsernameStarHandler == nil {
		api.ProfileDeleteProfilesUsernameStarHandler = profile.DeleteProfilesUsernameStarHandlerFunc(func(params profile.DeleteProfilesUsernameStarParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.DeleteProfilesUsernameStar has not yet been implemented")
		})
	}
	if api.ArticlesGetArticlesHandler == nil {
		api.ArticlesGetArticlesHandler = articles.GetArticlesHandlerFunc(func(params articles.GetArticlesParams) middleware.Responder {
			return middleware.NotImplemented("operation articles.GetArticles has not yet been implemented")
		})
	}
	if api.ArticlesGetArticlesArticleIDHandler == nil {
		api.ArticlesGetArticlesArticleIDHandler = articles.GetArticlesArticleIDHandlerFunc(func(params articles.GetArticlesArticleIDParams) middleware.Responder {
			return middleware.NotImplemented("operation articles.GetArticlesArticleID has not yet been implemented")
		})
	}
	if api.CommentsGetArticlesArticleIDCommentsHandler == nil {
		api.CommentsGetArticlesArticleIDCommentsHandler = comments.GetArticlesArticleIDCommentsHandlerFunc(func(params comments.GetArticlesArticleIDCommentsParams) middleware.Responder {
			return middleware.NotImplemented("operation comments.GetArticlesArticleIDComments has not yet been implemented")
		})
	}
	if api.CommentsGetCommentsCommentIDCommentsHandler == nil {
		api.CommentsGetCommentsCommentIDCommentsHandler = comments.GetCommentsCommentIDCommentsHandlerFunc(func(params comments.GetCommentsCommentIDCommentsParams) middleware.Responder {
			return middleware.NotImplemented("operation comments.GetCommentsCommentIDComments has not yet been implemented")
		})
	}
	if api.FeedsGetFeedsHandler == nil {
		api.FeedsGetFeedsHandler = feeds.GetFeedsHandlerFunc(func(params feeds.GetFeedsParams) middleware.Responder {
			return middleware.NotImplemented("operation feeds.GetFeeds has not yet been implemented")
		})
	}
	if api.MeGetMeHandler == nil {
		api.MeGetMeHandler = me.GetMeHandlerFunc(func(params me.GetMeParams) middleware.Responder {
			return middleware.NotImplemented("operation me.GetMe has not yet been implemented")
		})
	}
	if api.MeGetMeNotificationsHandler == nil {
		api.MeGetMeNotificationsHandler = me.GetMeNotificationsHandlerFunc(func(params me.GetMeNotificationsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation me.GetMeNotifications has not yet been implemented")
		})
	}
	if api.MeGetMeTagsHandler == nil {
		api.MeGetMeTagsHandler = me.GetMeTagsHandlerFunc(func(params me.GetMeTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation me.GetMeTags has not yet been implemented")
		})
	}
	if api.ProfileGetProfilesHandler == nil {
		api.ProfileGetProfilesHandler = profile.GetProfilesHandlerFunc(func(params profile.GetProfilesParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.GetProfiles has not yet been implemented")
		})
	}
	if api.ProfileGetProfilesUsernameHandler == nil {
		api.ProfileGetProfilesUsernameHandler = profile.GetProfilesUsernameHandlerFunc(func(params profile.GetProfilesUsernameParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.GetProfilesUsername has not yet been implemented")
		})
	}
	if api.ArticlesPostArticlesHandler == nil {
		api.ArticlesPostArticlesHandler = articles.PostArticlesHandlerFunc(func(params articles.PostArticlesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation articles.PostArticles has not yet been implemented")
		})
	}
	if api.ArticlesPostArticlesArticleIDVoteHandler == nil {
		api.ArticlesPostArticlesArticleIDVoteHandler = articles.PostArticlesArticleIDVoteHandlerFunc(func(params articles.PostArticlesArticleIDVoteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation articles.PostArticlesArticleIDVote has not yet been implemented")
		})
	}
	if api.CommentsPostCommentsHandler == nil {
		api.CommentsPostCommentsHandler = comments.PostCommentsHandlerFunc(func(params comments.PostCommentsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation comments.PostComments has not yet been implemented")
		})
	}
	if api.CommentsPostCommentsCommentIDVoteHandler == nil {
		api.CommentsPostCommentsCommentIDVoteHandler = comments.PostCommentsCommentIDVoteHandlerFunc(func(params comments.PostCommentsCommentIDVoteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation comments.PostCommentsCommentIDVote has not yet been implemented")
		})
	}
	if api.FeedsPostFeedsHandler == nil {
		api.FeedsPostFeedsHandler = feeds.PostFeedsHandlerFunc(func(params feeds.PostFeedsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.PostFeeds has not yet been implemented")
		})
	}
	if api.FeedsPostFeedsFeedIDFollowHandler == nil {
		api.FeedsPostFeedsFeedIDFollowHandler = feeds.PostFeedsFeedIDFollowHandlerFunc(func(params feeds.PostFeedsFeedIDFollowParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.PostFeedsFeedIDFollow has not yet been implemented")
		})
	}
	if api.FeedsPostFeedsFeedIDJoinHandler == nil {
		api.FeedsPostFeedsFeedIDJoinHandler = feeds.PostFeedsFeedIDJoinHandlerFunc(func(params feeds.PostFeedsFeedIDJoinParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.PostFeedsFeedIDJoin has not yet been implemented")
		})
	}
	if api.FeedsPostFeedsFeedIDJoinUserIDHandler == nil {
		api.FeedsPostFeedsFeedIDJoinUserIDHandler = feeds.PostFeedsFeedIDJoinUserIDHandlerFunc(func(params feeds.PostFeedsFeedIDJoinUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.PostFeedsFeedIDJoinUserID has not yet been implemented")
		})
	}
	if api.AuthPostLoginHandler == nil {
		api.AuthPostLoginHandler = auth.PostLoginHandlerFunc(func(params auth.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostLogin has not yet been implemented")
		})
	}
	if api.MePostMeNotificationsHandler == nil {
		api.MePostMeNotificationsHandler = me.PostMeNotificationsHandlerFunc(func(params me.PostMeNotificationsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation me.PostMeNotifications has not yet been implemented")
		})
	}
	if api.ProfilePostProfilesUsernameStarHandler == nil {
		api.ProfilePostProfilesUsernameStarHandler = profile.PostProfilesUsernameStarHandlerFunc(func(params profile.PostProfilesUsernameStarParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.PostProfilesUsernameStar has not yet been implemented")
		})
	}
	if api.AuthPostRegistrationHandler == nil {
		api.AuthPostRegistrationHandler = auth.PostRegistrationHandlerFunc(func(params auth.PostRegistrationParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostRegistration has not yet been implemented")
		})
	}
	if api.ArticlesPutArticlesArticleIDHandler == nil {
		api.ArticlesPutArticlesArticleIDHandler = articles.PutArticlesArticleIDHandlerFunc(func(params articles.PutArticlesArticleIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation articles.PutArticlesArticleID has not yet been implemented")
		})
	}
	if api.FeedsPutFeedsFeedIDHandler == nil {
		api.FeedsPutFeedsFeedIDHandler = feeds.PutFeedsFeedIDHandlerFunc(func(params feeds.PutFeedsFeedIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation feeds.PutFeedsFeedID has not yet been implemented")
		})
	}
	if api.MePutMeSettingsHandler == nil {
		api.MePutMeSettingsHandler = me.PutMeSettingsHandlerFunc(func(params me.PutMeSettingsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation me.PutMeSettings has not yet been implemented")
		})
	}
	if api.MePutMeSettingsNotificationsHandler == nil {
		api.MePutMeSettingsNotificationsHandler = me.PutMeSettingsNotificationsHandlerFunc(func(params me.PutMeSettingsNotificationsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation me.PutMeSettingsNotifications has not yet been implemented")
		})
	}
	if api.MePutMeSettingsPasswordHandler == nil {
		api.MePutMeSettingsPasswordHandler = me.PutMeSettingsPasswordHandlerFunc(func(params me.PutMeSettingsPasswordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation me.PutMeSettingsPassword has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
