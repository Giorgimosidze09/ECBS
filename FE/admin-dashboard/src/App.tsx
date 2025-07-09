import React, { useState, useEffect, useCallback } from 'react';
import { BrowserRouter as Router, Route, Switch, NavLink, useLocation } from 'react-router-dom';
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

const linkStyle: React.CSSProperties = {
  color: '#fff',
  textDecoration: 'none',
  padding: '0.75rem 1rem',
  borderRadius: '8px',
  fontWeight: 500,
  transition: 'background 0.2s, transform 0.2s',
  display: 'block',
};

const activeLinkStyle: React.CSSProperties = {
  background: '#4a4e69',
  transform: 'scale(1.04)',
};

function Layout() {
  const location = useLocation();
  const [sidebarOpen, setSidebarOpen] = useState(false);
  const [isMobile, setIsMobile] = useState(window.innerWidth < 900);
  const isDashboard = location.pathname === '/';

  useEffect(() => {
    const handleResize = () => setIsMobile(window.innerWidth < 900);
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);

  // Only close sidebar on mobile when navigating
  const handleNavClick = useCallback(() => {
    if (isMobile) setSidebarOpen(false);
  }, [isMobile]);

  // On dashboard, close sidebar by default on mobile
  useEffect(() => {
    if (isDashboard && isMobile) setSidebarOpen(false);
  }, [isDashboard, isMobile]);

  return (
    <div className="app-container">
      {/* Sidebar */}
      <div
        className={`sidebar${sidebarOpen ? ' open' : ''}${isMobile ? ' mobile' : ' desktop'}`}
        style={{
          position: isMobile ? 'fixed' : 'relative',
          left: isMobile ? (sidebarOpen ? 0 : '-220px') : 0,
          top: 0,
          width: 220,
          minHeight: '100vh',
          background: '#22223b',
          padding: '2rem 1rem',
          color: '#fff',
          gap: '1rem',
          display: 'flex',
          flexDirection: 'column',
          boxShadow: isMobile && sidebarOpen ? '0 0 0 100vw rgba(0,0,0,0.3)' : 'none',
          zIndex: 1000,
          transition: 'left 0.3s cubic-bezier(.4,2,.6,1)',
        }}
      >
        <button
          className="scale-hover"
          style={{
            display: isMobile ? 'block' : 'none',
            position: 'absolute',
            top: 16,
            right: 16,
            background: 'none',
            border: 'none',
            color: '#fff',
            fontSize: 28,
            cursor: 'pointer',
            zIndex: 1100,
          }}
          onClick={() => setSidebarOpen(false)}
          aria-label="Close sidebar"
        >
          ×
        </button>
        <h2 className="nav-label" style={{ color: '#c9ada7', marginBottom: '2rem' }}>Admin Dashboard</h2>
        <NavLink to="/" exact style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Dashboard
        </NavLink>
        <NavLink to="/users" style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Users
        </NavLink>
        <NavLink to="/cards" style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Cards
        </NavLink>
        <NavLink to="/balances" style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Balances
        </NavLink>
        <NavLink to="/validate" style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Card Validation
        </NavLink>
        <NavLink to="/ride" style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Ride Cost
        </NavLink>
        <NavLink to="/activate" style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Activate Card
        </NavLink>
        <NavLink to="/devices" style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
          Devices
        </NavLink>
      </div>
      {/* Hamburger menu for mobile */}
      <button
        className="scale-hover hamburger"
        style={{
          position: 'fixed',
          top: 18,
          left: 18,
          background: '#4a4e69',
          color: '#fff',
          border: 'none',
          borderRadius: 8,
          fontSize: 28,
          width: 44,
          height: 44,
          zIndex: 1200,
          display: isMobile && !sidebarOpen ? 'block' : 'none',
        }}
        onClick={() => setSidebarOpen(true)}
        aria-label="Open sidebar"
      >
        ☰
      </button>
      {/* Main content */}
      <div
        className="main-content fade-in"
        style={{
          flex: 1,
          minHeight: '100vh',
          background: 'var(--color-bg)',
          padding: isDashboard ? '0' : '2rem',
          width: '100%',
          overflowX: 'auto',
        }}
      >
        <PageRoutes />
      </div>
    </div>
  );
}

function PageRoutes() {
  return (
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
  );
}

const App: React.FC = () => (
  <Router>
    <Layout />
  </Router>
);

export default App;