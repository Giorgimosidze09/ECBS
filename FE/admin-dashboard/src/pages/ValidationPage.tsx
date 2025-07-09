import React from 'react';
import CardValidationForm from '../components/CardValidationForm';

const ValidationPage: React.FC = () => {
    return (
        <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <h1
                style={{
                    color: 'var(--color-primary)',
                    textAlign: 'center',
                    marginBottom: 32,
                    letterSpacing: 1,
                    fontWeight: 800,
                    fontSize: 32,
                }}
            >
                Card Validation
            </h1>
            <div
                className="fade-in scale-hover"
                style={{
                    maxWidth: 420,
                    width: '100%',
                    margin: '0 auto',
                    background: 'linear-gradient(135deg, var(--color-accent) 30%, var(--color-bg) 90%)',
                    borderRadius: 'var(--radius)',
                    boxShadow: '0 4px 24px var(--color-shadow)',
                    padding: 32,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    transition: 'box-shadow var(--transition)',
                }}
            >
                <CardValidationForm />
            </div>
        </div>
    );
};

export default ValidationPage;