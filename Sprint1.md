# Sprint 1

## Relevant User Stories
- As a reader, I want to be able to use a navigation bar, so that I can navigate between different sections on the blog.
- As a reader, I want the ability to view the homepage, so that I can go back to the homepage whenever I please.
- As a reader, I want the ability to open articles, so that I can read them.
- As a reader, I want the ability to search for articles, so that I can find and read them.

## Issues to Address
On the frontend side, we planned to create an initial React App ([See issue][i3]) and fully develop a functional navigation bar that routes the user between the main pages of the site ([See issue][i1]). In addition, the website logo should serve as another method of returning to the home page, so we planned to complete this routing issue as well ([See issue][i2]). Furthermore, we planned to create a home page with a view of the featured articles ([See issue][i11]) and to create a base article view of each page ([See issue][i12]). Finally, we also planned to create a backend simulator, to demonstrate the project([See issue][i10]). 

On the backend side, we planned to create a mock Go server with a default GET route ([See issue][i4]). We additionally planned to create an SQLite database that will be used for storing data such as the articles and the comments ([See issue][i5]) and within it populate an Articles table with dummy data to be used for testing search & retrieval behavior ([See issue][i6]). In order to enable the Go server to interact with the database, we planned to leverage GORM, an ORM library for Go ([See issue][i7]). Finally, we planned to implement the functionality for two routes on the Go server: one for retrieving articles by article ID ([See issue][i8]) and one for searching through articles by their titles ([See issue][i9]). Both of these functions require using GORM to interact with the SQLite database.

## Issues Completed
On the frontend side, we were successful in creating the intial React App ([See issue][i3]). We also created a functional navigation bar that routes the reader between the 3 main pages of the website ([See issue][i1]). However, while functional, the overall design of the navigation bar may be improved at a later date. We also created the home page, with a view of featured articles ([See issue][i11]), and also created a base article view of each page ([See issue][i12]). Lastly, we succesfully created a functional backend simulator ([See issue][i10]).

On the backend side, we successfully created a Go HTTP server on which we created multiple routes that accept GET requests ([See issue][i4]). We also created the SQLite database and an Articles table within it whose schema was cooperative with GORM ([See issue][i5]). We successfully populated the Articles table with 51 data points, such that their contents were composed of text and images arranged in various permutations, in order to test search & retrieval functionality ([See issue][i6]). We also set up GORM to interface with the SQLite database and to be usable by the server ([See issue][i7]). We completed the article retrieval behavior so that GET requests with article ID as a query parameter are met with a response of the data for the intended article, including the order of text and images and IDs of any images ([See issue][i8]).

## Issues Not Completed and Why
On the frontend side, we were not successful in ensuring that the website logo served as another method of returning to the home page ([See issue][i2]). While the logo appears on the home page, clicking the logo did did not return the reader to the home page. We had some issues with routing, which will be fixed at a later date. 


[i1]: https://github.com/apangasa/cen3031-skjsports/issues/1
[i2]: https://github.com/apangasa/cen3031-skjsports/issues/2
[i3]: https://github.com/apangasa/cen3031-skjsports/issues/3
[i4]: https://github.com/apangasa/cen3031-skjsports/issues/4
[i5]: https://github.com/apangasa/cen3031-skjsports/issues/5
[i6]: https://github.com/apangasa/cen3031-skjsports/issues/6
[i7]: https://github.com/apangasa/cen3031-skjsports/issues/7
[i8]: https://github.com/apangasa/cen3031-skjsports/issues/8
[i9]: https://github.com/apangasa/cen3031-skjsports/issues/9
[i10]:https://github.com/apangasa/cen3031-skjsports/issues/10
[i11]:https://github.com/apangasa/cen3031-skjsports/issues/11
[i12]:https://github.com/apangasa/cen3031-skjsports/issues/12 
