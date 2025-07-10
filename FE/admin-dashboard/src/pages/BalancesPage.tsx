import React from 'react';
import BalanceTopUpForm from '../components/BalanceTopUpForm';
import { Link } from 'react-router-dom';

const BalancesPage: React.FC = () => {
    return (
        <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32, width: '100%' }}>
                <h1 style={{ color: 'var(--color-primary)', letterSpacing: 1, fontWeight: 800, fontSize: 32 }}>Balance Top-Up</h1>
                <div style={{ display: 'flex', gap: 16 }}>
                    <Link
                        to="/balances/list"
                        className="scale-hover"
                        style={{
                            display: 'inline-block',
                            background: 'linear-gradient(90deg, var(--color-primary) 60%, var(--color-accent) 100%)',
                            color: '#fff',
                            border: 'none',
                            borderRadius: 8,
                            padding: '10px 24px',
                            fontWeight: 700,
                            fontSize: 16,
                            boxShadow: '0 2px 8px var(--color-shadow)',
                            textDecoration: 'none',
                            transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
                            cursor: 'pointer',
                        }}
                    >
                        View All Balances
                    </Link>
                </div>
            </div>
            <div className="fade-in scale-hover" style={{ maxWidth: 420, width: '100%', margin: '0 auto', background: 'linear-gradient(135deg, var(--color-primary) 30%, var(--color-accent) 90%)', borderRadius: 'var(--radius)', boxShadow: '0 4px 24px var(--color-shadow)', padding: 32, marginBottom: 32, display: 'flex', flexDirection: 'column', alignItems: 'center', transition: 'box-shadow var(--transition)' }}>
                <BalanceTopUpForm />
            </div>
        </div>
    );
};

export default BalancesPage;
