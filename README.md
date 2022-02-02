# SharedChemBoard-Web

The main purpose of this software is to help students and teachers keep learning while being on quarantine.

This software provides a WEB app.

## Current functionality:

- Teachers can create rooms where students can join
- Board for all class where teacher can draw and students can watch
- Personal board with each student where both - teacher and this student can draw
- Teacher has personal chat with each student
- Each board has history and users have an option to roll it back

## How it is made

- All data is stored in the PostgreSQL database
- Server is written in Go, frontend in vanilla JS
- All boards are a set of vector graphics

### General problems

- Frontend doesn't look very pretty..
- Code of accessing database has bad codestyle so it was stashed

### Current bugs

- Problems with connecting mutiple computers to server (but you can open several browsers to see multiuserness in action)
