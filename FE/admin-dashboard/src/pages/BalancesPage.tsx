import React from 'react';
import BalanceTopUpForm from '../components/BalanceTopUpForm';
import { Link } from 'react-router-dom';

const BalancesPage: React.FC = () => {
    return (
        <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
            <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
                <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>Balance Top-Up</h1>
                <div style={{ display: 'flex', gap: 16 }}>
                    <Link
                        to="/balances/list"
                        style={{
                            display: 'inline-block',
                            background: 'linear-gradient(90deg, #4a4e69 60%, #b5ead7 100%)',
                            color: '#fff',
                            border: 'none',
                            borderRadius: 8,
                            padding: '10px 24px',
                            fontWeight: 700,
                            fontSize: 16,
                            boxShadow: '0 2px 8px #b5ead722',
                            textDecoration: 'none',
                            transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
                            cursor: 'pointer',
                        }}
                        onMouseOver={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1.06)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 4px 16px #b5ead744'; }}
                        onMouseOut={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 2px 8px #b5ead722'; }}
                    >
                        View All Balances 
                    </Link>
                    <Link
                        to="/charges/list"
                        style={{
                            display: 'inline-block',
                            background: 'linear-gradient(90deg, #4a4e69 60%, #c9ada7 100%)',
                            color: '#fff',
                            border: 'none',
                            borderRadius: 8,
                            padding: '10px 24px',
                            fontWeight: 700,
                            fontSize: 16,
                            boxShadow: '0 2px 8px #c9ada722',
                            textDecoration: 'none',
                            transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
                            cursor: 'pointer',
                        }}
                        onMouseOver={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1.06)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 4px 16px #c9ada744'; }}
                        onMouseOut={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 2px 8px #c9ada722'; }}
                    >
                        View All Charges 
                    </Link>
                </div>
            </div>
            <div style={{ maxWidth: 420, margin: '0 auto', background: 'linear-gradient(135deg, #4a4e69 30%, #b5ead7 90%)', borderRadius: 16, boxShadow: '0 4px 24px #4a4e6922', padding: 32, marginBottom: 32, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                <BalanceTopUpForm />
            </div>
        </div>
    );
};

export default BalancesPage;
