import React from 'react';
import UserForm from '../components/UserForm';
import { Link } from 'react-router-dom';

const UsersPage: React.FC = () => {
    return (
        <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32, width: '100%' }}>
                <h1 style={{ color: 'var(--color-primary)', letterSpacing: 1, fontWeight: 800, fontSize: 32 }}>Users</h1>
                <Link
                    to="/users/list"
                    className="scale-hover"
                    style={{
                        display: 'inline-block',
                        background: 'linear-gradient(90deg, var(--color-primary) 60%, var(--color-secondary) 100%)',
                        color: '#fff',
                        border: 'none',
                        borderRadius: 8,
                        padding: '10px 28px',
                        fontWeight: 700,
                        fontSize: 16,
                        boxShadow: '0 2px 8px var(--color-shadow)',
                        textDecoration: 'none',
                        transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
                        cursor: 'pointer',
                    }}
                >
                    View All Users
                </Link>
            </div>
            <div className="fade-in scale-hover" style={{ maxWidth: 420, width: '100%', margin: '0 auto', background: 'linear-gradient(135deg, var(--color-secondary) 30%, var(--color-accent) 90%)', borderRadius: 'var(--radius)', boxShadow: '0 4px 24px var(--color-shadow)', padding: 32, marginBottom: 32, display: 'flex', flexDirection: 'column', alignItems: 'center', transition: 'box-shadow var(--transition)' }}>
                {/* Centered and styled UserForm */}
                <UserForm />
            </div>
        </div>
    );
};

export default UsersPage;