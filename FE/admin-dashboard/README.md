# Admin Dashboard

This project is an admin dashboard UI built with React that interfaces with a Go backend API. The dashboard provides various functionalities for managing users, cards, and balances.

## Project Structure

```
admin-dashboard
├── public
│   └── index.html          # Main HTML file for the React application
├── src
│   ├── api
│   │   └── index.ts        # API interaction methods
│   ├── components
│   │   ├── UserForm.tsx    # Form for creating a new user
│   │   ├── CardAssignmentForm.tsx # Form for assigning a card to a user
│   │   ├── BalanceTopUpForm.tsx   # Form for topping up a user's balance
│   │   └── CardValidationForm.tsx  # Form for validating a card
│   ├── pages
│   │   ├── UsersPage.tsx   # Page for user management
│   │   ├── CardsPage.tsx    # Page for card management
│   │   ├── BalancesPage.tsx  # Page for balance management
│   │   └── ValidationPage.tsx # Page for card validation
│   ├── App.tsx              # Main application component
│   ├── index.tsx            # Entry point for the React application
│   └── types
│       └── index.ts         # TypeScript interfaces and types
├── package.json             # npm configuration file
├── tsconfig.json            # TypeScript configuration file
└── README.md                # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd admin-dashboard
   ```

2. **Install dependencies:**
   ```
   npm install
   ```

3. **Run the application:**
   ```
   npm start
   ```

4. **Access the dashboard:**
   Open your browser and navigate to `http://localhost:3000`.

## Usage Guidelines

- Use the **Users Page** to create new users.
- Use the **Cards Page** to assign cards to users.
- Use the **Balances Page** to top up user balances.
- Use the **Validation Page** to validate cards.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.