# Sprint 3

## Relevant User Stories
- As a reader, I want the ability to open articles, so that I can read them.
- As a reader, I want the ability to search for articles, so that I can find and read them.
- As a reader, I want the ability to subscribe to the blog, so that I can get emails about new articles.
- As a reader, I want to be able to hover over a player’s name and see statistics about that player, so that I can be better informed while reading an article. 
- As a reader, I want to be able to hover over a team’s name and see statistics about that team, so that I can be better informed while reading an article. 
- As a reader, I want the ability to comment on articles, so that I can share my thoughts about various articles.
- As a writer, I want to be able to create and save drafts, so that I can write articles and keep saved copies of my work.

## Unit Tests
The following unit tests were necessary for frontend behavior:
- Testing the subscription button
- Testing the draft board

The following unit tests were necessary for backend behavior:
- Testing that new subscribers are added to and can be retrieved from the database
- Testing that individuals who resubscribe while already subscribed do not modify the database (original personal details are used)
- Testing that unsubscribers are not retrieved when querying subscribers
- Testing that resubscribers with the same details can be retrieved
- Testing that resubscribers with different details can be retrieved and have the same details (i.e. after unsubscribing, someone resubscribes under same email but new name)


## Issues to Address
On the frontend side, we decided to work on the subscription feature ([See issue][i40]) of the blog, so that readers can subscribe to the blog, and keep up with the newest articles. Also, we focused on implementing the draft board for the writer ([See issue][i59]), as well as the ability for the writer to actually open drafts on the draft board ([See issue][i60]). This way, the writer of the blog would be able to create, save, and edit drafts before publishing them to the site. In terms of tests, we wanted to create unit tests for both the subscription button ([See issue][i58]) and the draft board ([See issue][i61]).

On the backend side, we planned to tackle issues related to storing ([See issue][i41]) and retrieving ([See issue][i42]) images. Also, various behavior needed to be addressed with regards to subscribers. We planned to create a subscribers table ([See issue][i46]), implement functionality to add ([See issue][i45]) and remove ([See issue][i47]) users to/from it, create a route for the frontend to indicate an unsubscription event ([See issue][i44]), and handle the cases of resubscription when a user is already subscribed ([See issue][i43]) and when they have previously unsubscribed ([See issue][i48]). All this subscription behavior required a number of unit tests as mentioned previously ([See issue][i63]). We also wanted to tackle the player statistics part of the project - one of the first steps in doing so would be web scraping the relevant positions for each soccer player ([See issue][i53]) so that position-relevant statistics could be gathered. We also managed to write functions to read relevant statistics for players by 4 major positions: goalkeeper([See issue][i54]), defenders([See issue][i55]), midfielders ([See issue][i56]), and forwards ([See issue][i57]). Additionally, we wanted to add routes on the server for the frontend to request player ([See issue][i50]) and team ([See issue][i51]) stats.

## Issues Completed
On the frontend side, we managed to succesfully create the subscription button ([See issue][i40]). If a reader so wishes, they are able to enter their email address into the site, and they will have subscribed to the blog. Also, we managed to create the draft board ([See issue][i59]) for the writer, and also succesfully implemented the ability for the writer to actually open drafts on the draft board ([See issue][i60]). We also completed the unit test for the subscription button, using Cypress ([See issue][i58]).

On the backend side, we completely finished subscription functionality. That is, we created the subscribers table ([See issue][i46]), methods to add subscribers ([See issue][i45]) and remove subscribers ([See issue][i47]), and a server route for unsubscription events ([See issue][i44]). Additionally, we successfully dealt with the receipt of resubscription event while a user is already subscribed ([See issue][i43]) and after they have unsubscribed ([See issue][i48]). We fully wrote the 5 unit tests to test this behavior and they passed ([See issue][i63]). We started on player statistics by adding routes for getting player stats ([See issue][i50]) and getting team stats ([See issue][i51]). We also set up a web scraper to read fbref.com to find out the positions played by any particular player given their name ([See issue][i53]). This allowed us to procure player statistics by each of the four positions: goalkeeper ([See issue][i54]), defender ([See issue][i55]), midfielder ([See issue][i56]), and forward ([See issue][i57]).

## Issues Not Completed and Why
On the frontend side, we were not able to succesfully create a unit test for the draft board ([See issue][i61]). We were unable to configure the tests properly, so it is something we will come back to and fix at a later date.

On the backend side, we were also unable to create unit tests for getting player statistics ([See issue][i66]). This was due to loopholes remaining in our code whenever a search query returned multiple results and didn't reach the player page. Resolving this issue will be paramount moving forward.



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
[i40]:https://github.com/apangasa/cen3031-skjsports/issues/40
[i41]:https://github.com/apangasa/cen3031-skjsports/issues/41
[i42]:https://github.com/apangasa/cen3031-skjsports/issues/42 
[i43]:https://github.com/apangasa/cen3031-skjsports/issues/43 
[i44]:https://github.com/apangasa/cen3031-skjsports/issues/44 
[i45]:https://github.com/apangasa/cen3031-skjsports/issues/45 
[i46]:https://github.com/apangasa/cen3031-skjsports/issues/46 
[i47]:https://github.com/apangasa/cen3031-skjsports/issues/47 
[i48]:https://github.com/apangasa/cen3031-skjsports/issues/48 
[i49]:https://github.com/apangasa/cen3031-skjsports/issues/49
[i50]:https://github.com/apangasa/cen3031-skjsports/issues/50 
[i51]:https://github.com/apangasa/cen3031-skjsports/issues/51 
[i53]:https://github.com/apangasa/cen3031-skjsports/issues/53 
[i58]:https://github.com/apangasa/cen3031-skjsports/issues/58
[i59]:https://github.com/apangasa/cen3031-skjsports/issues/59
[i60]:https://github.com/apangasa/cen3031-skjsports/issues/60
[i61]:https://github.com/apangasa/cen3031-skjsports/issues/61

[i63]:https://github.com/apangasa/cen3031-skjsports/issues/63
