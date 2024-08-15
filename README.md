# record_shelf

##API endpoints

User profile endpoints:
- `GET /shelf/traffic` gives a simple html as response with the count of how many times the `/` endpoint was visited.
- `POST /login` takes in a form data, authenticates a user and logs them in.
- `POST /add/users` accepts user credentials in the request and creates a user in the database.
- `PUT /update/users/{userID}` updates the credentials of an existing user in the database.
- `DELETE /remove/users/{userID}` marks a user deleted in the database.

File endpoints:
- `GET /files/users/{userID}` adds a file to the database for user with userID, a file_name will be automatically assigned.
- `GET /files/users/{userID}` responds with `file_name`file(name extracted from queries) requested from the database.
- `PUT /files/users/{userID}` replaces the `file_name`file(name extracted from queries) with the old one.
- `DELETE /files/users/{userID}` marks the file with `file_name`file(name extracted from queries) deleted, hence making it not-accessible.
