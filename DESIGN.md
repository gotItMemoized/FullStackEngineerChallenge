# Design

## Frontend

### Setting Up

Taking advantage of create-react-app as it provides a lot of features out of the box and I can customize those thing if I really need to.

## Backend

### Setting up

Writing the backend in golang. Sorry if there's any issue with getting golang set up on your machine. Since I'm using modules it should be a little bit easier, but it can be a little quirky to set up. I'm much more familiar with it for writing backend services from nothing. Ruby-On-Rails in api mode would possibly be much easier, but not as fun to talk about in an interview.

I'm making an assumption that there's going to initially always be an admin and a user. This will be set up with the `-setupDatabase -seedDatabase` flags (can see more in the README.md on how to use).

### Auth

Using JWTs for auth. However, for simplicity (although very insecure), I'm not expiring the JWT tokens quickly and not providing a refresh token. This is not a production safe app, but it should show the basic usages of the jwt while showing some auth-like features.

### Security

Making a few security compromises. First is mentioned in the Auth section. I'm not providing refresh tokens and the jwt token doesn't have a short expiration time.

Second, I'm not using https. This makes the app vulnerable to MITM attacks. 

Also, by default not using SSL for connecting to the database.

### Database

Using postgresql and taking advantage of sql and sqlx packages for golang.

#### Tables

##### user

|name|description|
|-|-|
|id|reference id|
|name| users real/display name |
|username| username for login |
|password| bcrypted password hash |
|isAdmin|if the user has admin permissions|

##### reviews
|name|description|
|-|-|
|id|reference id|
|userId|user getting reviewed|
|isActive|if the performance review is still open for feedback|

##### review_feedback
|name|description|
|-|-|
|id|reference id|
|userId|user getting reviewed|
|reviewerId|user giving feedback|
|feedback|feedback content|

#### Database Design Compromises

As mentioned in the Auth section. I'm omitting some security things like refresh tokens. So there's no additional table for that. I'm also omitting timestamps (create, update, etc) from the tables for the sake of simplicity, but it'd be important to have in a real application.