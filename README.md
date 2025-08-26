# Event Management API

This is an API created by following [Maximilian Schwarzmuller's Go course](https://www.udemy.com/course/go-the-complete-guide). It is for a website that manages events and has functionalities for event creation, handling, user sign up and log in, and event registration.

## Table of contents

- [Overview](#overview)
  - [The challenge](#the-challenge)
  - [Screenshot](#screenshot)
  - [Links](#links)
- [My process](#my-process)
  - [Built with](#built-with)
- [Acknowledgments](#acknowledgments)

## Overview
This project uses many different Go concepts and libraries to ensure the smoothness of the processes along with the necessary security measures to stop malicious attacks. It has password hashing and authentication to keep users' data safe, along with gating certain functionalities behind user verification to prevent unintended data changes. 
### The challenge

Users should be able to:

- Create an account and log in.
- Create events and be the sole manager of their events.
- Update their events and delete them.
- Register for available events and unregister at will.

### Links

- [Maximilian's Go course](https://www.udemy.com/course/go-the-complete-guide)

## My process
The project starts by creating the structure of the events and creating the table for them in the database. Then we move on to dealing with the event handling; creating, updating, and deleting. The user account structure is then set and the necessary table created and account creation and logging in is handled. Password hashing and user authentication for security purposes is then handled. And finally event registration is added to allow users to register for events. 
### Built with

- Golang
- [Gin Web Framework](https://gin-gonic.com/)
- [bcrypt library](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)

## Acknowledgments

Maximillian's course was really helpful at teaching me Go and this project is a testement to his teaching ability. All code credit goes to him