# SnippetBox

## API Routes

|S.N|Method|Pattern|Handler|Action|
|:--|:----:|:------|:-----:|:-----|
|1|ANY|/|home|Display the home page|
|2|ANY|/snippet/view/:id|snippetView|Display a specific snippet|
|3|GET|/snippet/create|snippetCreate|Display a HTML form for creating a new snippet|
|4|POST|/snippet/create|snippetCreatePost|Create a new snippet|
|5|ANY |/static/*filepath| http.FileServer| Serve a specific static file|
