# Learnable BE  
---
## Description  
Have you ever had so many browser tabs open because you didn't want to lose the content that was on them? Well, this application was built for you! Our application allows you to have one web page open that stores all of that content you want to read, watch, and do. Here you can create playlists for storing articles, videos, or other content. This application can also help with time management. You can set a 'due date' for when you want to complete a playlist.  Once an item within a playlist is complete, you simply mark it done and watch your progress bar fill up.  
##### You can visit our website Learnable here: https://learnablebe.herokuapp.com/api/v0  

![GitHub Logo](/images/playlist.png)  

## Our code  
This is the back-end repository of our full-stack application. The coding language we chose to use was Golang, with a PostgreSQL database. In our database, we have users, playlists, and playlist items.  
We have created twelve API endpoints. The API endpoints are for our front-end team. It allows them to get the appropriate data back when they send us an HTTP request. Some of those requests include creating a playlist, getting all playlist for a certain user, deleting a playlist, creating a playlist item, updating a playlist item, and more.  
##### You can access our front-end repository here: https://github.com/learn-able/learnable-fe
---
## Installation Locally  
- Clone this repo to your local machine with: `git clone git@github.com:learn-able/learnable-be.git`
- cd into the directory
- Install Golang with `brew install go`
- Start Postgres (either the app or command line)
- Add 'application.yml' file to config folder
  1. DB_USER: "your_postgres_user_name"
  1. DB_PASSWORD: "postgres"
  1. DB_ADDRESS: "localhost:5432"
  1. DB_NAME: "postgres"
- In learnabledb.go and main.go:
  1. uncomment code between LOCAL tags
  1. comment out code between HEROKU tags
- In the command line, `go run main.go`


To get your postgres username from the command line:
  1. `psql postgres`
  1. `\du`
  1. `exit`


Golang documentation: https://golang.org/doc/
PostgreSQL documentation: https://www.postgresql.org
---
## Requests and Responses
![GitHub Logo](/images/req_res1.png)  
![GitHub Logo](/images/req_res2.png)  
![GitHub Logo](/images/req_res3.png)  
![GitHub Logo](/images/req_res4.png)  
![GitHub Logo](/images/req_res5.png)  
---
## Project Collaborators  
- Steven Anderson - https://github.com/alerrian  
- Meghan Stovall - https://github.com/meghanstovall  
- Elom Amouzou - https://github.com/eamouzou  
- Chris Postma - https://github.com/cjrpostma  
- Ryan Bahan - https://github.com/ryanbahan
