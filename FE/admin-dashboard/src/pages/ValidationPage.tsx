import React from 'react';
import CardValidationForm from '../components/CardValidationForm';

const ValidationPage: React.FC = () => {
    return (
        <div
            style={{
                minHeight: '100vh',
                background: '#f2e9e4',
                padding: '40px 0',
            }}
        >
            <h1
                style={{
                    color: '#4a4e69',
                    textAlign: 'center',
                    marginBottom: 32,
                    letterSpacing: 1,
                }}
            >
                Card Validation
            </h1>
            <div
                style={{
                    maxWidth: 420,
                    margin: '0 auto',
                    background: 'linear-gradient(135deg, #c9ada7 30%, #f2e9e4 90%)',
                    borderRadius: 16,
                    boxShadow: '0 4px 24px #c9ada722',
                    padding: 32,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center'
                }}
            >
                <CardValidationForm />
            </div>
        </div>
    );
};

export default ValidationPage;