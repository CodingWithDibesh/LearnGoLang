# SnippetBox

## API Routes

|S.N|Method|Pattern|Handler|Action|
|:--|:----:|:------|:-----:|:-----|
|1|ANY|/|home|Display the home page|
|2|ANY|/snippet/view/:id|snippetView|Display a specific snippet|
|3|GET|/snippet/create|snippetCreate|Display a HTML form for creating a new snippet|
|4|POST|/snippet/create|snippetCreatePost|Create a new snippet|
|5|GET|/user/signup|userSignup|Display a HTML form for signing up a new user|
|6|POST|/user/signup|userSignupPost|Create a new user|
|7|GET|/user/login|userLogin|Display a html form for logging in a user|
|8|POST|/user/login|userLoginPost|Authenticate and login the user|
|9|POST|/user/logout|userLogoutPost|Logout the user|
|10|ANY |/static/*filepath| http.FileServer| Serve a specific static file|
