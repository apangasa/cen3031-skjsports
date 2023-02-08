# Sprint 1

## Relevant User Stories
- As a reader, I want to be able to use a navigation bar, so that I can navigate between different sections on the blog.
- As a reader, I want the ability to view the homepage, so that I can go back to the homepage whenever I please.
- As a reader, I want the ability to open articles, so that I can read them.
- As a reader, I want the ability to search for articles, so that I can find and read them.

## Issues to Address
On the frontend side, we planned to create an initial React App ([See issue][i3]) and fully develop a functional navigation bar that routes the user between the 3 main pages of the site ([See issue][i1]). In addition, the website logo should serve as another method of returning to the home page, so we planned to complete this routing issue as well ([See issue][i2]). 

On the backend side, we planned to create a mock Go server with a default GET route ([See issue][i4]). We additionally planned to create an SQLite database that will be used for storing data such as the articles and the comments ([See issue][i5]) and within it populate an Articles table with dummy data to be used for testing search & retrieval behavior ([See issue][i6]). In order to enable the Go server to interact with the database, we planned to leverage GORM, an ORM library for Go ([See issue][i7]). Finally, we planned to implement the functionality for two routes on the Go server: one for retrieving articles by article ID ([See issue][i8]) and one for searching through articles by their titles ([See issue][i9]). Both of these functions require using GORM to interact with the SQLite database.

## Successful Issues
On the frontend side, we were successful in creating the intial React App ([See issue][i3]). We also created a functional navigation bar that routes the reader between the 3 main pages of the website ([See issue][i1]). However, while functional, the overall design may be improved at a later date. 

## Unsuccessful Issues 
On the frontend side, we were not successful in ensuring that the website logo served as another method of returning to the home page ([See issue][i2]). 


[i1]: https://github.com/apangasa/cen3031-skjsports/issues/1
[i2]: https://github.com/apangasa/cen3031-skjsports/issues/2
[i3]: https://github.com/apangasa/cen3031-skjsports/issues/3
[i4]: https://github.com/apangasa/cen3031-skjsports/issues/4
[i5]: https://github.com/apangasa/cen3031-skjsports/issues/5
[i6]: https://github.com/apangasa/cen3031-skjsports/issues/6
[i7]: https://github.com/apangasa/cen3031-skjsports/issues/7
[i8]: https://github.com/apangasa/cen3031-skjsports/issues/8
[i9]: https://github.com/apangasa/cen3031-skjsports/issues/9
