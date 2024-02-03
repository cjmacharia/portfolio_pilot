# Portfolio Pilot Project

This is a Golang-based Finance Application designed to showcase the integration of finance and technology. 
The application processes financial data, provides real-time stock prices, manages investment portfolios, 
and assists with budget management. It utilizes programming languages like Golang, a PostgreSQL database for data storage, 
and provides a set of RESTful API endpoints for data retrieval and manipulation.

## Features

- Real-time stock prices (available in version 2)
- Investment portfolio management
- Budget management
- Database integration with PostgreSQL
- RESTful API endpoints for data processing

## Technologies Used

- Golang
- PostgreSQL

## Getting Started

### Prerequisites

- Golang installed
- PostgreSQL database

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/portfolio-pilot.git
```Stocks
Endpoint 1: /api/stocks (GET): Retrieve a list of available stocks.

Portfolios
Endpoint 5: /api/user/{id:[0-9]+}/portfolio" (POST): view a users investment portfolio.

Transactions
Endpoint 10: /api/transactions (GET): Retrieve a list of user transactions.

Endpoint 11: /api/transactions/{transactionID} (GET): Retrieve details of a specific transaction.

Endpoint 12: /api/stock/{id}/transactions (POST): Add a new transaction.
