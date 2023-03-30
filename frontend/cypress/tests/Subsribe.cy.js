import React from 'react'
import Home from './../../src/Home'
import App from './../../src/App'
import SubscribeForm from "../../src/Subscribe";

before(() => {
    // root-level hook
    // runs once before all tests
    "These are tests to check if subscribe works"
})
let i =0;
describe('<SubscribeForm />', () => {
    it('render home and subscribe', () => {
        // see: https://on.cypress.io/mounting-react
        if (i<1) {
            cy.mount(<App/>)
            i = i + 1

            if (cy.get('[id^=subscribeType]').should('not.be.disabled') && i<1) {
                cy.get('[id^=subscribeType]').type("kolliparas@ufl.edu")
                cy.get('[id^=subscribeButton]').click()
            }
        }
    })

})

//cy.get('input').type('Hello, World')