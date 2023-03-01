import React from 'react'
import App from '../../src/App'

before(() => {
  // root-level hook
  // runs once before all tests
  "These tests check if app is able to load up"
})

describe('<App />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    //test to see if search bar loaded up
    cy.mount(<App />)
  })
})