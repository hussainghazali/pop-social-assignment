# POP SOCIAL ASSIGNMENT - HUSSAIN GHAZALI

# Project Setup & Running the App

**Step 1:** Clone the repository to your local machine:

```bash
git clone https://github.com/hussainghazali/pop-social-assignment.git
```

**Step 2:** Change your current directory to the project directory:

```bash
cd pop-social-assignment
```

**Step 3:** Set up your Docker environment. Run the following command to start the project's containers using Docker Compose:

```bash
docker-compose up
```

**Step 4:** After the containers are up and running, you can start the project with the following command:

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/da8ab618-0ca1-432e-927c-518c3b6b00eb)

```bash
air
```

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/0bb9f624-48d6-4d95-a7c5-f77cc871ee80)

Your project is now running on port 8000, and you can access all the APIs.

Make sure to have Docker and Air installed in your development environment to execute the steps successfully.

# Unit Test the Application

1. Make sure you have the Go testing framework installed in your development environment.

2. To run the unit tests, use the following command:

```bash
go test ./controllers
```

3. The testing process will create a test database (`test.db`) and execute test cases, including scenarios for both success and errors.

4. The test results will be displayed in the console, allowing you to verify that the API endpoints function as expected.

By following these steps, you can ensure that your application's controllers and endpoints are thoroughly tested for various scenarios.

# 1. Setup and Database

- I have successfully configured PostgreSQL with my Go application. All essential entities, including User, Post, Comment, and Like, have been created in PgAdmin, complete with their respective relationships and schema mappings.

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/089d5166-57b9-495d-8515-721850800472)

- To ensure smooth project execution and the ability to make real-time updates, I've utilized a `.air.toml` configuration for hot reloading.

- Additionally, I've dockerized the PostgreSQL connection to streamline the development environment.

# 2. User Management:

**Register User API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/afbd4eda-d6a1-4cbf-920e-23c2e009f824)

- The Register User API allows users to create accounts by providing their email, password, and confirming the password. To enhance security, passwords are securely hashed before storage.

- Ensuring uniqueness, duplicate emails are not allowed. Each user is a distinct entity with a unique email address.

**Login User API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/0a3317ac-6a9d-48bd-a91d-6c0e7e42a117)

- Upon successfully logging in, a token is issued to the user. This token acts as a middleware for performing actions like posting, adding comments, and liking posts. 

- User-specific restrictions are in place: users can only update or delete their own posts, comments, and likes.

**Refresh Token API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/c2bd8a17-0336-4358-90f8-74746e58ad77)

- The Refresh Token API is responsible for renewing a user's session when it expires. A new token is issued to the signed-in user, providing continuous access to the application.

**User Logout API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/51067746-a038-4dc3-9ecf-6ab0b354f8b6)

- The User Logout API allows users to terminate their session. Once logged out, users cannot perform any further operations related to posts, comments, or likes. To regain access, they must log in again.

**Me (User Detail) API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/f54ada36-743b-4885-b771-3a2068e996c1)

- The Me API provides detailed user profile information. Users can access their own profile details through this endpoint.

By implementing these features, I've ensured robust user management and data security within the application.

# 3. Post Management

**Create Post API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/92bf8ab8-180c-4966-af24-c17be1940e72)

- Users with a valid token can create a post. 
- Posts contain details of likes and comments, including the count.
- This functionality ensures that only verified users can contribute to the platform.

**Update Post API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/c887f23d-9d0b-4a79-943c-8f56ae574357)

- Users receive a token upon login, and only the user who created a post can update it.
- The update functionality permits modifications to the post's content, including title, image, video, and text.

**Delete Post API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/6b4dd0b9-99e0-439a-8014-0db8e6d3333c)

- Similarly, users can delete a post. However, like the update functionality, this action is restricted to the user who created the post.

**Get All Posts API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/b4ff9f0a-c3e0-4632-b719-af1db9caf135)

- This API allows users to retrieve a list of posts, complete with the details of likes and comments for each post.

**Get Posts By ID API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/556c3795-0ca7-4d9e-92f1-31d23287db72)

- Users can request a specific post using this API, along with its corresponding likes and comments.

**Get Posts By User API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/bc1bafbd-f645-4f28-b5db-2a5e146c998a)

- Users can access a list of all the posts they have created, making it easy to review their contributions to the platform.

This comprehensive Post Management system ensures user-generated content can be created, updated, and deleted with a fine level of control, while also making it easy for users to access and interact with posts in a meaningful way.

# 4. Comment Management

**Create Comment API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/04a36c73-9b12-46eb-9a8b-e829d25ef0fa)

- Users can add comments to any post of their choice.
- Multiple comments can be added to any post, fostering discussion and engagement within the platform.

**Update Comment API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/8998b62a-dcc1-45ad-92f5-7f0f86d718ea)

- Users have the ability to update their comments on any post.
- Similar to adding comments, multiple comments can be updated for any post.

**Delete Comment API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/1d2f3c14-72db-4858-9806-75f8688c020d)

- Users can also delete their comments from any post.

# 5. Like Management

**Add Like API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/ad49fd2e-ac30-4b32-b575-23d6d387b299)

- Users can express their appreciation for a post by adding a "like" to it.
- Each user can add a like to any post they desire, contributing to the post's popularity.

**Remove Like API:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/94d4ecc0-f425-4705-836b-932eb5a9eaa1)

- Users have the flexibility to remove their "like" from a post.
- This feature enables users to manage their preferences and interactions with posts effectively, ensuring their likes are always a genuine expression of interest.

This versatile functionality allows users to manage their contributions and interactions effectively, facilitating a rich and dynamic community environment.

# 6. API Documentation

**Swagger:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/46b14fd3-dae8-4855-8714-f1dbc6075d54)

- The project includes comprehensive Swagger documentation that becomes available when you run the application and visit http://localhost:8000/swagger/index.html#/comments.
- The Swagger documentation provides detailed information about schemas, models, responses, status codes, parameters, and descriptions for all available APIs.

**Postman:**

![image](https://github.com/hussainghazali/pop-social-assignment/assets/50146615/b77dc986-a450-47a0-b5d9-45da4107fecb)

- The code repository contains a "Postman" folder that includes a JSON collection of the entire application.
- To access the collection and its detailed request and response examples, simply import the provided JSON file into your Postman workspace. This collection allows you to explore and interact with all the APIs seamlessly.

Thank You.
