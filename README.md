# Mock SSO Flow – Authorization Code Exchange (Go)

## Overview

This project implements a mock Single Sign-On (SSO) flow using the Authorization Code Grant concept.  
It is a command-line based simulation that demonstrates how real-world OAuth / SSO systems work internally.

No actual server or external Identity Provider is used.  
Everything is simulated in-memory for learning and explanation purposes.

* * *

## Why This Project Was Built

The goal of this assignment is to understand:

*   How SSO login works internally  
      
    
*   How authorization codes are generated and exchanged  
      
    
*   How tokens are issued and verified  
      
    
*   How security concepts like state, expiry, and scopes work  
      
    

This project focuses on concept clarity, not production deployment.

* * *

## Concepts Covered

*   Authorization Code Flow  
      
    
*   Identity Provider (IdP) simulation  
      
    
*   Token generation and validation  
      
    
*   State parameter validation (CSRF protection)  
      
    
*   Token expiry handling  
      
    
*   One-time authorization code usage  
      
    
*   Scopes and claims  
      
    
*   Logout and token revocation  
      
    

* * *

## Features Implemented

### Core Assignment Features

*   Redirect simulation to Identity Provider  
      
    
*   Username and password input (mock validation)  
      
    
*   Authorization code generation  
      
    
*   Token exchange  
      
    
*   ID token and access token generation  
      
    
*   Token verification  
      
    
*   Claim validation  
      
    

### Additional Features

*   State parameter validation  
      
    
*   One-time use authorization code  
      
    
*   Token expiry handling  
      
    
*   Scope-based token data  
      
    
*   Logout and token revocation  
      
    
*   Menu-driven flow using switch-case  
      
    
*   Continuous execution loop  
      
    

* * *

## Project Structure

mock-sso/

│

├── main.go

├── go.mod

└── README.doc

  

* * *

## How the Flow Works (Step-by-Step)

### 1\. Start SSO Login

Program prints:  
*   Redirecting to Identity Provider...  
*   User enters username and password  
*   Login is assumed valid  
      
  Program generates:  
*   Authorization Code  
*   State parameter  
      
    
* * *

### 2\. Authorization Code Exchange

User enters:  
*   Auth code  
*   State  
      
  Program validates:  
*   Auth code correctness  
*   State value  
*   One-time usage  
      
  If valid:  
*   ID token is generated 
*   Access token is generated  
*   Token expiry time is set     
*   Claims and scopes are attached  
      
* * *

### 3\. Token Verification

Program verifies:

*   Token is not empty  
*   Token structure is valid (. separated)  
*   Token is not expired  
      
  Required claims exist:  
  
*   sub  
*   iss  
*   role  

If all checks pass:  
Token Verified is printed along with claims.
* * *

### 4\. Logout (Token Revocation)

*   Clears:  
*   Tokens  
*   Auth code  
*   State  
*   Forces fresh login for next access  
      
    

* * *

## Data Structures Used

### Token Structure

```
Token {

  IDToken
  AccessToken
  Claims
  Expiry
  Scopes
}
```
  

### Claims Example

```
{
  "sub": "user123",
  "iss": "mock-idp",
  "role": "user"
}
```
  

* * *

## Security Concepts Demonstrated

*   CSRF protection using state parameter  
*   Prevention of auth code reuse  
*   Token expiry enforcement  
*   Claim-based identity validation  
*   Scope-based authorization data  

* * *

## Limitations

*   No real encryption or JWT signing  
*   No database or persistence  
*   Passwords are not validated securely  
*   Intended only for learning purposes
