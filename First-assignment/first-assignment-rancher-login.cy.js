describe('Rancher Login', () => {
  it('should log in to Rancher', () => {
    // defining variables for baseurl, username and password
    const baseUrl = "http://localhost:80"
    const username = "admin"
    const password = "SuSetechnicaltest"
    // Visit the Rancher login page
    cy.visit(baseUrl)
    
    // Enter the login credentials
    cy.get('#username').type(username)
    cy.get('#password').type(password)
    
    // Submit the login form
    cy.get('#submit').click()
    
    // Wait for the page to load after login and checking the title of the webpage
    cy.title().should('eq', 'Rancher', { timeout: 10000 });
  
})
})
