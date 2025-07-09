import React from 'react';
import UserForm from '../components/UserForm';
import { Link } from 'react-router-dom';

const UsersPage: React.FC = () => {
    return (
        <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
            <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
                <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>Users</h1>
                <Link
                    to="/users/list"
                    style={{
                        display: 'inline-block',
                        background: 'linear-gradient(90deg, #4a4e69 60%, #9a8c98 100%)',
                        color: '#fff',
                        border: 'none',
                        borderRadius: 8,
                        padding: '10px 28px',
                        fontWeight: 700,
                        fontSize: 16,
                        boxShadow: '0 2px 8px #9a8c9822',
                        textDecoration: 'none',
                        transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
                        cursor: 'pointer',
                    }}
                    onMouseOver={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1.06)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 4px 16px #9a8c9844'; }}
                    onMouseOut={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 2px 8px #9a8c9822'; }}
                >
                    View All Users 
                </Link>
            </div>
            <div style={{
                maxWidth: 420,
                margin: '0 auto',
                background: 'linear-gradient(135deg, #9a8c98 30%, #c9ada7 90%)',
                borderRadius: 16,
                boxShadow: '0 4px 24px #9a8c9822',
                padding: 32,
                marginBottom: 32,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center'
            }}>
                {/* Centered and styled UserForm */}
                <UserForm />
            </div>
        </div>
    );
};

export default UsersPage;