import React from 'react'
import Home from './../../src/Home'
import App from './../../src/App'

before(() => {
    // root-level hook
    // runs once before all tests
    "These are tests to check if an article loads text"
})

describe('<Article />', () => {
    it('render article and check if text is there', () => {
        // see: https://on.cypress.io/mounting-react

        cy.mount(<App />)
        cy.get('p').contains('sohil').click()
        cy.get('p').contains('article 1 part 1')
    })

})