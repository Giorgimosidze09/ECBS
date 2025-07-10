import React, { useState, useEffect, useCallback } from 'react';
import { BrowserRouter as Router, Route, Switch, NavLink, useLocation, Redirect, useHistory } from 'react-router-dom';
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
import CustomerPage from './pages/CustomerPage';
import { AuthProvider, useAuth } from './auth';

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

// PrivateRoute: only for authenticated users
const PrivateRoute: React.FC<any> = ({ children, ...rest }) => {
  const { token } = useAuth();
  return <Route {...rest} render={() => (token ? children : <Redirect to="/login" />)} />;
};

// RoleRoute: only for users with a specific role
const RoleRoute: React.FC<any> = ({ role: requiredRole, children, ...rest }) => {
  const { token, role } = useAuth();
  return <Route {...rest} render={() => (token && role === requiredRole ? children : <Redirect to="/login" />)} />;
};

const LoginPage = React.lazy(() => import('./pages/LoginPage'));
const RegisterPage = React.lazy(() => import('./pages/RegisterPage'));

function Layout() {
  const location = useLocation();
  const { token, role, logout } = useAuth();
  const history = useHistory();
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

  // Hide sidebar and layout for unauthenticated or non-admin users
  if (!token || role !== 'admin') {
    return <PageRoutes />;
  }
  const handleLogout = () => {
    logout();
    history.push('/login');
  };
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
        <NavLink to="/admin" exact style={linkStyle} activeStyle={activeLinkStyle} onClick={handleNavClick} className="scale-hover">
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
        <button
          onClick={handleLogout}
          style={{
            marginTop: '2rem',
            background: 'linear-gradient(90deg, #c72c41 60%, #9a8c98 100%)',
            color: '#fff',
            border: 'none',
            borderRadius: 8,
            padding: '12px 0',
            fontWeight: 700,
            fontSize: 18,
            cursor: 'pointer',
            boxShadow: '0 2px 8px #c9ada7',
            width: '100%'
          }}
        >
          Logout
        </button>
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
    <React.Suspense fallback={<div>Loading...</div>}>
      <Switch>
        <Route path="/login" component={LoginPage} />
        <Route path="/register" component={RegisterPage} />
        <RoleRoute path="/admin" role="admin">
          <DashboardPage />
        </RoleRoute>
        <RoleRoute path="/users/list" role="admin">
          <UsersListPage />
        </RoleRoute>
        <RoleRoute path="/users" role="admin">
          <UsersPage />
        </RoleRoute>
        <RoleRoute path="/cards/list" role="admin">
          <CardsListPage />
        </RoleRoute>
        <RoleRoute path="/cards" role="admin">
          <CardsPage />
        </RoleRoute>
        <RoleRoute path="/balances/list" role="admin">
          <BalancesListPage />
        </RoleRoute>
        <RoleRoute path="/balances" role="admin">
          <BalancesPage />
        </RoleRoute>
        <RoleRoute path="/validate" role="admin">
          <ValidationPage />
        </RoleRoute>
        <RoleRoute path="/ride" role="admin">
          <RideCostPage />
        </RoleRoute>
        <RoleRoute path="/activate" role="admin">
          <ActivateCardPage />
        </RoleRoute>
        <RoleRoute path="/devices/list" role="admin">
          <DevicesListVisualPage />
        </RoleRoute>
        <RoleRoute path="/devices" role="admin">
          <DevicesListPage />
        </RoleRoute>
        <RoleRoute path="/charges/list" role="admin">
          <ChargesListPage />
        </RoleRoute>
        <RoleRoute path="/customer" role="customer">
          <CustomerPage />
        </RoleRoute>
        <Redirect to="/admin" />
      </Switch>
    </React.Suspense>
  );
}

const App: React.FC = () => (
  <AuthProvider>
    <Router>
      <Layout />
    </Router>
  </AuthProvider>
);

export default App;