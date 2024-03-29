# Sprint 4

## Relevant User Stories
- As a reader, I want the ability to view the homepage, so that I can go back to the homepage whenever I please. 
- As a reader, I want the ability to open articles, so that I can read them.
- As a reader, I want the ability to search for articles, so that I can find and read them.
- As a reader, I want to be able to hover over a player’s name and see statistics about that player, so that I can be better informed while reading an article. 
- As a reader, I want to be able to hover over a team’s name and see statistics about that team, so that I can be better informed while reading an article. 
- As a reader, I want the ability to comment on articles, so that I can share my thoughts about various articles.
- As a writer, I want to be able to create and save drafts, so that I can write articles and keep saved copies of my work.
- As a writer, I want to be able to publish drafts, so that readers can view my articles. 
- As a writer, I want to be able to edit published articles, so that I can go back and fix any mistakes after publication. 
- As an administrator, I want to ask potential writers for their email, so that they can sign up as a writer for the blog. 

## Unit Tests
The following unit tests were necessary for frontend behavior: 
- Testing the draft board
- Testing subscription button
- Testing article view
- Testing search results
- Testing standard home

The following unit tests were necessary for backend behavior:
- Testing the method for getting team stats ([See issue][i66])
- Testing login functionality
- Testing to see if multi-word nations were read correctly for player stats ([See issue][i72])


## Issues Completed
On the frontend side, to finish up the project, we had to work on implementing the draft board ([See issue][i95]), adding various writer features ([See issue][i96]), and connecting the frontend and backend ([See issue][i97]). We were successful in completing these issues. However, the login system was unable to work properly. Regardless, the final product was not affected. The draft board was fully implemented, and the other writer features we wanted to add (edit, save, publish drafts) were implemented as well. We also had to actually design the blog ([See issue][i73]). This was the final step of our process, as we had to make the website look aesthetically pleasing, and attract potential readers. We were successfully able to do so, and managed to create the product we had envisioned at the start of this project. 

On the backend side, we implemented login routes ([See issue][i74]), login functions ([See issue][i75]). We also added all draft-related behavior such as creating (([See issue][i79]), editing ([See issue][i81]), and publishing drafts ([See issue][i84]). Functionality was added to retrieve articles ([See issue][i85]) and drafts ([See issue][i83]) by author. We added a concept of authors to articles and drafts ([See issue][i82]). We also implemented image functionality ([See issue][i42]) with image upload and retrieval ([See issue][i41]). Additionally, we wrapped up web scraping with adding in team-level stats ([See issue][i76]).


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
[i54]:https://github.com/apangasa/cen3031-skjsports/issues/54
[i55]:https://github.com/apangasa/cen3031-skjsports/issues/55
[i56]:https://github.com/apangasa/cen3031-skjsports/issues/56
[i57]:https://github.com/apangasa/cen3031-skjsports/issues/57
[i58]:https://github.com/apangasa/cen3031-skjsports/issues/58
[i59]:https://github.com/apangasa/cen3031-skjsports/issues/59
[i60]:https://github.com/apangasa/cen3031-skjsports/issues/60
[i61]:https://github.com/apangasa/cen3031-skjsports/issues/61
[i63]:https://github.com/apangasa/cen3031-skjsports/issues/63
[i66]:https://github.com/apangasa/cen3031-skjsports/issues/66
[i72]:https://github.com/apangasa/cen3031-skjsports/issues/72
[i73]:https://github.com/apangasa/cen3031-skjsports/issues/73
[i74]:https://github.com/apangasa/cen3031-skjsports/issues/74
[i75]:https://github.com/apangasa/cen3031-skjsports/issues/75
[i76]:https://github.com/apangasa/cen3031-skjsports/issues/76
[i77]:https://github.com/apangasa/cen3031-skjsports/issues/77
[i79]:https://github.com/apangasa/cen3031-skjsports/issues/79
[i80]:https://github.com/apangasa/cen3031-skjsports/issues/80
[i81]:https://github.com/apangasa/cen3031-skjsports/issues/81
[i82]:https://github.com/apangasa/cen3031-skjsports/issues/82
[i83]:https://github.com/apangasa/cen3031-skjsports/issues/83
[i84]:https://github.com/apangasa/cen3031-skjsports/issues/84
[i85]:https://github.com/apangasa/cen3031-skjsports/issues/85
[i86]:https://github.com/apangasa/cen3031-skjsports/issues/86
[i87]:https://github.com/apangasa/cen3031-skjsports/issues/87
[i88]:https://github.com/apangasa/cen3031-skjsports/issues/88
[i89]:https://github.com/apangasa/cen3031-skjsports/issues/89
[i91]:https://github.com/apangasa/cen3031-skjsports/issues/91
[i92]:https://github.com/apangasa/cen3031-skjsports/issues/92
[i95]:https://github.com/apangasa/cen3031-skjsports/issues/95
[i96]:https://github.com/apangasa/cen3031-skjsports/issues/96
[i97]:https://github.com/apangasa/cen3031-skjsports/issues/97

