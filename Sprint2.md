# Sprint 2

## Relevant User Stories
- As a reader, I want the ability to open articles, so that I can read them.
- As a reader, I want the ability to search for articles, so that I can find and read them.

## Unit Tests
The following unit tests were necessary for frontend behavior:  

- Testing the home page
- Testing the functionality of individual articles
- Testing the search bar function
- Identifying a method to handle react-router in Cypress

The following unit tests were necessary for backend behavior:

- Retrieving an article with an empty id
- Retrieving an article with an id that is not in the database
- Retrieving an article that is in the database and contains only text
- Retrieving an article that is in the database and contains both text and images
- Searching for articles with empty search string
- Search for a non-existent article
- Performing a search that retrieves multiple articles
- Performing a search that retrieves one article
- Ensuring the search function is not case sensitive
- Ensuring a search string that is blank returns no articles
- Ensuring a search string with multiple words still works

## Issues to Address
On the frontend side, we planned to implement a "search bar" function, so that users will be able to search for articles by keyword ([See issue][i24]).
To accomplish our task, we had to work with the backend team to set up our integrated application, in order to pull articles from the backend, and load them into the frontend. Also, we set out to complete unit tests and Cypress tests. Specifically, we created Cypress tests for the home page ([See issue][i28]), for individual articles ([See issue][i29]), for the search bar function ([See issue] [i30]), and to identify a method to handle react-router in Cypress([See issue][i31]).

On the backend side, we planned to work with the frontend team to integrate the frontend and backend in the sense that the frontend would be able to request articles by ID and search for articles via search terms ([See issue][i21]). In order to do this, we also had to enable cross-origin resource sharing on the backend ([See issue][i25]). We also planned to create API documentation as a YAML file using the OpenAPI specification ([See issue][i20]). Additionally, we planned to implement and run unit tests using Go's testing package for the article retrieval ([See issue][i22]) and for the article search ([See issue][i23]). Finally, we planned to create a new route on the server that accepts a POST request containing a subscriber information for readers to subscribe to the blog ([See issue][i32]).

## Issues Completed
On the frontend side, we managed to create the search bar ([See issue][i24]). If a user enters in a keyword, a list of articles containing that keyword appear. We also completed 3 Cypress tests, successfully creating and running tests for the home page ([See issue][i28]), for individual articles ([See issue][i29]), for the search bar function ([See issue][i30]). There is room for improvement though, as the search bar doesn't return articles to a different page. Ideally, the search bar would contain a drop-down menu, and would open articles on another page. Also, many of our tests can certainly be improved and made more functional at a later date. Both of these issues will be tackled in the near future.

On the backend side, we successfully created OpenAPI documentation for the GET requests that one can make to the backend server ([See issue][i20]). We did so by writing a YAML file in the OpenAPI specification and publishing it as an API Collection via Postman to generate an API guide. We also enabled CORS on the backend server in order to allow the frontend to successfully make requests and receive responses ([See issue][i25]). We did some refactoring of the code and implemented HTTP status codes for various errors, then completed unit testing for article retrieval ([See issue][i22]) and article search ([See issue][i23]), and all tests passed successfully. We also implemented the subscribe route that reads subscriber data from the POST request body and sends it to a method that would insert relevant data into the database's Subscribers table ([See issue][i32]). 

## Issues Not Completed and Why
On the frontend side, we did not manage to successfully create a Cypress test to identify a method to handle react-router in Cypress([See issue][i31]). We were not able to properly handle react-router in Cypress, due to unforseen effects on the search system. The viability of said test is not clear at the moment, sand it's something to discuss and handle at a later date. 

## Backend API Documentation
https://universal-crater-481750.postman.co/workspace/6185c94f-4893-4149-a166-29e01a85f960/api/e6f0ab7c-013c-4e89-b88e-f9d77a013bf7

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
[i20]:https://github.com/apangasa/cen3031-skjsports/issues/20 
[i21]:https://github.com/apangasa/cen3031-skjsports/issues/21
[i22]:https://github.com/apangasa/cen3031-skjsports/issues/22 
[i23]:https://github.com/apangasa/cen3031-skjsports/issues/23
[i24]:https://github.com/apangasa/cen3031-skjsports/issues/24
[i25]:https://github.com/apangasa/cen3031-skjsports/issues/25
[i28]:https://github.com/apangasa/cen3031-skjsports/issues/28
[i29]:https://github.com/apangasa/cen3031-skjsports/issues/29
[i30]:https://github.com/apangasa/cen3031-skjsports/issues/30
[i31]:https://github.com/apangasa/cen3031-skjsports/issues/31
[i32]:https://github.com/apangasa/cen3031-skjsports/issues/32 

