/// <reference types="cypress" />

const API_TIMEOUT_MS = 90000;

describe("Main page", () => {
  it("allows user to shorten the same URL multiple times", () => {
    cy.visit("/");

    cy.get("input")
      .type("https://www.google.com/webmasters")
      .should("have.value", "https://www.google.com/webmasters");

    cy.get("button")
      .as("submitButton")
      .click();

    cy.contains("#shortURL", /^https:\/\/casca\.dev\/goto\/\w{4}$/, {
      timeout: API_TIMEOUT_MS
    }).then($shortURL => {
      cy.get("@submitButton").click();
      cy.get("#shortURL").should("not.be.visible");
      cy.contains("#shortURL", $shortURL.text(), {
        timeout: API_TIMEOUT_MS
      });
    });
  });

  xit("is visually unchanged", () => {
    cy.visit("/");
    cy.get("#logo")
      .should("be.visible")
      .then(() => {
        // cy.get("input").type("abc");
        // cy.wait(3000);
        cy.percySnapshot("Homepage");
      });
  });
});
