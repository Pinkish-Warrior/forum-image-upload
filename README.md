## FORUM 👥🌐💬
[![Run Project](https://img.shields.io/badge/Run-Project-blue)](https://forum-image-upload.onrender.com)
## Imports:

The code begins with importing necessary packages, including standard libraries like database/sql, encoding/json, log, net/http, sync, and external packages like github.com/mattn/go-sqlite3 for SQLite database handling.

## Global Variables:

db: A global variable representing the SQLite database.
sessions: A map for storing session data, where the key is a session token and the value is a user ID.
mu: A sync.Mutex for handling concurrent access to shared resources.

## Main Function:

The main function is the entry point of the application.
It initializes the database, sets up routes, and starts the HTTP server.

##Database Initialization:
The forum.Init() function is called to initialize the database tables and schema.

## HTTP Route Handlers:

The code defines various HTTP route handlers for different pages and actions. Each handler corresponds to a specific endpoint.

##Handler Functions:
Functions like HandleMain, AuthenticatedMainPage, RegisterPage, etc., are responsible for rendering different pages based on user authentication and other conditions.

## User Authentication and Session Handling:

The code manages user authentication using cookies and session tokens. Functions like GetAuthenticatedUserData, GetAuthenticatedUserID, createSession, and clearSession handle session management.

## User Registration and Login:

RegisterPage renders the registration page, and Registration2 handles the registration process.
LoginHandler manages user login.

## Post and Comment Handling:

Functions like NewPostPage, AddPost, AddPostToDatabase, AddComment, and AddCommentHandler handle post and comment creation.

## Post and Comment Likes/Dislikes:

Functions like LikePostHandler, DislikePostHandler, LikeCommentHandler, and DislikeCommentHandler manage likes and dislikes for posts and comments.

### Profile Page:

ProfilePage renders the user's profile page.

### Database Queries:

Various functions perform database queries to fetch or update data. For example, GetPostsFromDatabase, GetCommentsForPost, etc.

### Category Filtering:

The code provides the ability to filter posts by category using the FilterPosts and GetPostsByCategory functions.

### Error Handling:

The code includes error handling to deal with potential issues during database operations, HTTP requests, and template rendering.

### Template Rendering:

The code uses Go's html/template package to render HTML templates.

### HTTP Server Initialization:

The server is started using http.ListenAndServe.

### Database Closure:

The defer db.Close() statement ensures that the database connection is closed when the program exits.

### Additional Features:

The code also includes functions for like and dislike counts, session handling, and various helper functions.

### Tests:

## Tests verbose:

```bash
go test -v
```

## Create Test Report

```bash
go test -coverprofile cover.out
```

## Open Report coverage in the browser

```bash
go tool cover -html=cover.out
```
