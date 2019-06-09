# Design

## Frontend

Taking advantage of create-react-app as it provides a lot of features out of the box and I can customize those thing if I really need to.

### Technical decisions

- create-react-app
  - debugging, testing, building tools included
  - can eject and change webpack/babel/etc configurations if needed.
  - kept up to date as it's maintained by developers on the react team
- For state management I just used react hooks rather than redux/mobx.
  - Less boiler plate
  - reduced complexity
  - if the project would become more complex or required coordination with more people, redux/mobx could become more necessary.
- axios to help with api requests
- partial JWT implementation for user login
  - has technical compromises mentioned below

### Envirment Variables

- `DEV_FRONTEND_PROXY`: proxy url for frontend. Only used in `dev` builds

## Backend

Wrote the backend in golang (requires v1.12+). Sorry if there's any issue with getting golang set up on your machine (Docker should make it easy to run the app). Since I'm using modules it should be a little bit easier, but it can be a little quirky to set up. I'm much more familiar with it for writing backend services from nothing. Ruby-On-Rails in api mode would possibly be much easier, but not as fun to talk about in an interview.

I'm making an assumption that there's going to initially always be an admin and a user. This will be set up with the `-resetPostgres -seedDatabase` flags (can see more in the [README.md](./README.md) on how to use).

### Technical Decisions

- Modules (requires Golang v1.12+)
  - gets dependencies easily and automatically
- Data interfaces for databases
  - allows you to work with a local in-memory database
  - implement different databases easily
  - currently implements a map-based in-memory database or postgresql
- JWT auth checks
  - compromises mentioned below

### Environment Variables

- `JWT_SECRET`: set the secret used in the JWTs. Defaults to `thisIsAnInsecureSecret`
- `ENV`: environment you're running in. Does not default! set to `DEV` for development. `PROD` would be production, but current there's no production configurations.
- `POSTGRES_CONNECTION`: sets the connection string to postgres. Defaults to `dbname=paypay sslmode=disable`

### Flags

- `-usePostgres`: use a postgres database (defaults to use in-memory database)
- `-resetDatabase`: reset the dev database (currently no-op if not using postgres)
- `-seedDatabase`: seed the dev database (currently no-op if not using postgres)

### Auth

Using JWTs for auth. However, for simplicity (although very insecure), I'm not expiring the JWT tokens quickly and not providing a refresh token. This is not a production safe app, but it should show the basic usages of the jwt while showing some auth-like features.

### Security

Making a few security compromises. First is mentioned in the Auth section. I'm not providing refresh tokens and the jwt token doesn't have a short expiration time.

Second, I'm not using https. This makes the app vulnerable to MITM attacks. 

Also, by default not using SSL for connecting to the database.

## Database

The design of the backend allows other implementations if necessary, but I'm a little more familiar with postgresql so I used it in this case.

### Tables

#### user

|name|description|
|-|-|
|id|reference id|
|name| users real/display name |
|username| username for login (unique) |
|password| bcrypted password hash |
|isAdmin|if the user has admin permissions|

#### reviews
|name|description|
|-|-|
|id|reference id|
|userId|user getting reviewed FK|
|isActive|if the performance review is still open for feedback|

#### review_feedback
|name|description|
|-|-|
|id|reference id|
|reviewId|reviews table FK|
|reviewerId|user giving feedback FK|
|message|feedback content|

### Database Design Compromises

As mentioned in the Auth section. I'm omitting some security things like refresh tokens. So there's no additional table for that. I'm also omitting timestamps (create, update, etc) from the tables for the sake of simplicity, but it'd be important to have in a real application.

Assigning a performance review can be added and deleted at any point. This can be a desired feature, but this current design doesn't allow for logical deletes. Logical deletes would allow you to delete, and undelete an item. This prevents you from losing the feedback message from the assigned user.

## Other Technologies Used

### Docker

I wanted to provide a way to easily run the application without needing to worry about configurations or having certain things installed.

### Postman/Newman

Postman is a tool I often use to check API endpoints. They have some testing tools built into it as well. I used this tool because it let me easily check existing functionality and ensure I maintain functionality as I make changes. I can make the tests more strict, but right now they're rather simple checks.

## Known issues / Compromises / Notes / Assumptions

- App is running in `DEV` mode in the containers and are not set up to build for production
- SSL/HTTPS not setup
- postman/newman api tests have hardcoded tokens that'll expire in like... 2 years. Can make this dynamic, I just need more time with the tool
- postman/newman api tests require clean database (`-resetPostgres -seedDatabase`)
- environment variable `JWT_TOKEN` is expected to not be set for the API tests
- JWT tokens expire in 2 years rather than in 15 minutes. There's no refresh token mechanism.
- Default (seeded) user/pass should be dev only and need a mechanism to have an initial admin for prod (where it forces you to change the password)
- Admins set the passwords for the users
- Users cannot change the passwords
- Not enough unit tests
- No pagination on the `/all` calls. This would be very bad for large datasets
- If someone submits feedback, and the admin removes them from the performance review, saves, and adds the user back. The users feedback is no longer available. (need soft deleting or do not allow feedback to be removed if already provided)
- when using in-memory database deleting user does not delete users in feedback/reviews (no cascading deletes)
- frontend doesn't give any good user feedback when it cannot connect to the backend.
- node_modules is too big and makes the current container size much bigger because we're still running in DEV (much smaller with the node:current-alpine image however) actually building and using something like nginx or kubernetes to route and serve files
- there's never enough tests
- Login won't work if local storage isn't supported by the browser or disabled. Also local storage is susceptable to XSS if you're not careful (if a library gets compromised, etc. This is why tokens are usually short lived rather than 2 years).
- Item ordering isn't sorted in any particular way from any of the `/all` calls

## Feature Ideas

Ideas of things that could be added if given more time (in addition to fixing above issues).

- allow more complex feedback options, maybe allow you to customize questions
- autoclose the performance review once all of the user feedback is in.
- user creation gives a invite code instead of setting the user password directly.
- typescript
- a11y linting
- precommit hooks for linters, formatters, tests
- regular users can change their own passwords
- proper pagination for get all users/performance reviews/feedback
- access codes for a specific reviewer/review
- redis cache?
- frontend docker image could use the build script, and nginx to proxy calls to the backend server
- proptypes
- more tests
- coordinate page colors better
- how to deal with 404s when cannot connect to backend?