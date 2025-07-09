import React from 'react';
import { BrowserRouter as Router, Route, Switch, NavLink } from 'react-router-dom';
import DashboardPage from './pages/DashboardPage';
import UsersPage from './pages/UsersPage';
import CardsPage from './pages/CardsPage';
import BalancesPage from './pages/BalancesPage';
import ValidationPage from './pages/ValidationPage';
import RideCostPage from './pages/RideCostPage';
import ActivateCardPage from './pages/ActivateCardPage';
import CardsListPage from './pages/CardsListPage';
import UsersListPage from './pages/UsersListPage';
import DevicesListPage from './pages/DevicesListPage';
import DevicesListVisualPage from './pages/DevicesListVisualPage';
import BalancesListPage from './pages/BalancesListPage';
import ChargesListPage from './pages/ChargesListPage';

const navStyle: React.CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  width: 220,
  minHeight: '100vh',
  background: '#22223b',
  padding: '2rem 1rem',
  color: '#fff',
  gap: '1rem',
  position: 'fixed',
  left: 0,
  top: 0,
};

const linkStyle: React.CSSProperties = {
  color: '#fff',
  textDecoration: 'none',
  padding: '0.75rem 1rem',
  borderRadius: '8px',
  fontWeight: 500,
};

const activeLinkStyle: React.CSSProperties = {
  background: '#4a4e69',
};

const contentStyle: React.CSSProperties = {
  marginLeft: 240,
  padding: '2rem',
  background: '#f2e9e4',
  minHeight: '100vh',
};

const App: React.FC = () => (
  <Router>
    <div style={navStyle}>
      <h2 style={{ color: '#c9ada7', marginBottom: '2rem' }}>Admin Dashboard</h2>
      <NavLink to="/" exact style={linkStyle} activeStyle={activeLinkStyle}>
        Dashboard
      </NavLink>
      <NavLink to="/users" style={linkStyle} activeStyle={activeLinkStyle}>
        Users
      </NavLink>
      <NavLink to="/cards" style={linkStyle} activeStyle={activeLinkStyle}>
        Cards
      </NavLink>
      <NavLink to="/balances" style={linkStyle} activeStyle={activeLinkStyle}>
        Balances
      </NavLink>
      <NavLink to="/validate" style={linkStyle} activeStyle={activeLinkStyle}>
        Card Validation
      </NavLink>
       <NavLink to="/ride" style={linkStyle} activeStyle={activeLinkStyle}>
        Ride Cost
      </NavLink>
      <NavLink to="/activate" style={linkStyle} activeStyle={activeLinkStyle}>
        Activate Card
      </NavLink>
      <NavLink to="/devices" style={linkStyle} activeStyle={activeLinkStyle}>
        Devices
      </NavLink>
    </div>
    <div style={contentStyle}>
      <Switch>
        <Route path="/" exact component={DashboardPage} />
        <Route path="/users/list" component={UsersListPage} />
        <Route path="/users" component={UsersPage} />
        <Route path="/cards/list" component={CardsListPage} />
        <Route path="/cards" component={CardsPage} />
        <Route path="/balances" exact component={BalancesPage} />
        <Route path="/validate" component={ValidationPage} />
        <Route path="/ride" component={RideCostPage} />
        <Route path="/activate" component={ActivateCardPage} />
        <Route path="/devices" exact component={DevicesListPage} />
        <Route path="/devices/list" component={DevicesListVisualPage} />
        <Route path="/balances/list" component={BalancesListPage} />
        <Route path="/charges/list" component={ChargesListPage} />
      </Switch>
    </div>
  </Router>
);

export default App;