import React from 'react'
import Home from './../../src/Home'
import App from './../../src/App'

before(() => {
  // root-level hook
  // runs once before all tests
  "These are tests to check if the home has loaded"
})

describe('<Home />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react

    cy.mount(<Home />)
  })
  it('check home from app perspective', () => {
    cy.mount(<App />)
    cy.get('p').contains('home')

  })
})