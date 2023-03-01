import React from 'react'
import Article from './Article'

describe('<Article />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Article />)
  })
})